package usecase

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrCannotSelfFriend   = errors.New("Không thể kết bạn với chính mình")
	ErrTargetNotFound     = errors.New("Người dùng không tồn tại")
	ErrRequestAlreadySent = errors.New("Lời mời kết bạn đã tồn tại hoặc đã là bạn bè")
	ErrRelationshipNotFound = errors.New("Không tìm thấy mối quan hệ hoặc lời mời")
	ErrNotAddressee       = errors.New("Bạn không có quyền thao tác trên lời mời này")
)

type FriendUsecase interface {
	GetFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	GetPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	SendFriendRequest(ctx context.Context, requesterID, addresseeID uuid.UUID) (*domain.Friend, error)
	AcceptFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error)
	DeclineFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error)
	BlockUser(ctx context.Context, userID, targetUserID uuid.UUID) (*domain.Friend, error)
	UnfriendOrCancel(ctx context.Context, userID, relationshipID uuid.UUID) (bool, error)
}

type friendUsecase struct {
	friendRepo repository.FriendRepository
	userRepo   repository.UserRepository
}

func NewFriendUsecase(friendRepo repository.FriendRepository, userRepo repository.UserRepository) FriendUsecase {
	return &friendUsecase{
		friendRepo: friendRepo,
		userRepo:   userRepo,
	}
}

func (u *friendUsecase) GetFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	return u.friendRepo.FindFriendsList(ctx, userID)
}

func (u *friendUsecase) GetPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	return u.friendRepo.FindPendingRequests(ctx, userID)
}

func (u *friendUsecase) SendFriendRequest(ctx context.Context, requesterID, addresseeID uuid.UUID) (*domain.Friend, error) {
	if requesterID == addresseeID {
		return nil, ErrCannotSelfFriend
	}

	targetUser, err := u.userRepo.FindByID(ctx, addresseeID)
	if err != nil {
		return nil, err
	}
	if targetUser == nil {
		return nil, ErrTargetNotFound
	}

	existing, err := u.friendRepo.FindByUsers(ctx, requesterID, addresseeID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		if existing.Status == domain.FriendStatusPending || existing.Status == domain.FriendStatusAccepted {
			return nil, ErrRequestAlreadySent
		}
		if existing.Status == domain.FriendStatusBlocked {
			return nil, errors.New("Không thể gửi lời mời do người dùng đã bị chặn")
		}
		// Nếu trước đó declined, cho phép gửi lại bằng cách cập nhật sang pending
		return u.friendRepo.UpdateStatus(ctx, existing.ID, domain.FriendStatusPending)
	}

	newFriend := &domain.Friend{
		ID:          uuid.New(),
		RequesterID: requesterID,
		AddresseeID: addresseeID,
		Status:      domain.FriendStatusPending,
	}

	if err := u.friendRepo.Create(ctx, newFriend); err != nil {
		return nil, fmt.Errorf("failed to send friend request: %w", err)
	}

	return u.friendRepo.FindByID(ctx, newFriend.ID)
}

func (u *friendUsecase) AcceptFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error) {
	friend, err := u.friendRepo.FindByID(ctx, requestID)
	if err != nil {
		return nil, err
	}
	if friend == nil {
		return nil, ErrRelationshipNotFound
	}
	if friend.AddresseeID != userID {
		return nil, ErrNotAddressee
	}
	if friend.Status != domain.FriendStatusPending {
		return nil, errors.New("Lời mời không ở trạng thái chờ duyệt")
	}

	return u.friendRepo.UpdateStatus(ctx, requestID, domain.FriendStatusAccepted)
}

func (u *friendUsecase) DeclineFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error) {
	friend, err := u.friendRepo.FindByID(ctx, requestID)
	if err != nil {
		return nil, err
	}
	if friend == nil {
		return nil, ErrRelationshipNotFound
	}
	if friend.AddresseeID != userID {
		return nil, ErrNotAddressee
	}

	return u.friendRepo.UpdateStatus(ctx, requestID, domain.FriendStatusDeclined)
}

func (u *friendUsecase) BlockUser(ctx context.Context, userID, targetUserID uuid.UUID) (*domain.Friend, error) {
	if userID == targetUserID {
		return nil, errors.New("Không thể chặn chính mình")
	}

	existing, err := u.friendRepo.FindByUsers(ctx, userID, targetUserID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return u.friendRepo.UpdateStatus(ctx, existing.ID, domain.FriendStatusBlocked)
	}

	newBlock := &domain.Friend{
		ID:          uuid.New(),
		RequesterID: userID,
		AddresseeID: targetUserID,
		Status:      domain.FriendStatusBlocked,
	}
	if err := u.friendRepo.Create(ctx, newBlock); err != nil {
		return nil, err
	}

	return u.friendRepo.FindByID(ctx, newBlock.ID)
}

func (u *friendUsecase) UnfriendOrCancel(ctx context.Context, userID, relationshipID uuid.UUID) (bool, error) {
	friend, err := u.friendRepo.FindByID(ctx, relationshipID)
	if err != nil {
		return false, err
	}
	if friend == nil {
		return false, ErrRelationshipNotFound
	}
	if friend.RequesterID != userID && friend.AddresseeID != userID {
		return false, errors.New("Bạn không thuộc mối quan hệ này")
	}

	return u.friendRepo.Delete(ctx, relationshipID)
}
