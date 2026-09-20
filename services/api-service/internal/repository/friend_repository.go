package repository

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FriendRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Friend, error)
	FindByUsers(ctx context.Context, userA, userB uuid.UUID) (*domain.Friend, error)
	FindPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	FindFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	Create(ctx context.Context, friend *domain.Friend) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status domain.FriendStatus) (*domain.Friend, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
}

type friendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return &friendRepository{db: db}
}

func (r *friendRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Friend, error) {
	var friend domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Preload("Addressee").
		Where("id = ?", id).
		First(&friend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find friend by id: %w", err)
	}
	return &friend, nil
}

func (r *friendRepository) FindByUsers(ctx context.Context, userA, userB uuid.UUID) (*domain.Friend, error) {
	var friend domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Preload("Addressee").
		Where("((\"requesterId\" = ? AND \"addresseeId\" = ?) OR (\"requesterId\" = ? AND \"addresseeId\" = ?))", userA, userB, userB, userA).
		First(&friend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find friend relationship: %w", err)
	}
	return &friend, nil
}

func (r *friendRepository) FindPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	var list []domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Where("\"addresseeId\" = ? AND status = ?", userID, domain.FriendStatusPending).
		Order("\"createdAt\" DESC").
		Find(&list).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find pending friend requests: %w", err)
	}
	return list, nil
}

func (r *friendRepository) FindFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	var list []domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Preload("Addressee").
		Where("(\"requesterId\" = ? OR \"addresseeId\" = ?) AND status = ?", userID, userID, domain.FriendStatusAccepted).
		Order("\"updatedAt\" DESC").
		Find(&list).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find friends list: %w", err)
	}
	return list, nil
}

func (r *friendRepository) Create(ctx context.Context, friend *domain.Friend) error {
	if friend.ID == uuid.Nil {
		friend.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(friend).Error
}

func (r *friendRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status domain.FriendStatus) (*domain.Friend, error) {
	if err := r.db.WithContext(ctx).Model(&domain.Friend{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return nil, fmt.Errorf("failed to update friend status: %w", err)
	}
	return r.FindByID(ctx, id)
}

func (r *friendRepository) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Friend{})
	return res.RowsAffected > 0, res.Error
}
