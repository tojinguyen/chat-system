package usecase

import (
	"context"
	"errors"

	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrConversationNotFound = errors.New("Cuộc hội thoại không tồn tại hoặc bạn không phải thành viên")
	ErrNotMember            = errors.New("Bạn không phải thành viên của cuộc hội thoại này")
	ErrCannotChatSelf       = errors.New("Không thể tạo cuộc hội thoại với chính mình")
	ErrNotAdmin             = errors.New("Chỉ quản trị viên mới có quyền thao tác")
)

type ConversationUsecase interface {
	GetUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error)
	GetOrCreateDirectConversation(ctx context.Context, userID, partnerID uuid.UUID) (*domain.Conversation, error)
	CreateGroupConversation(ctx context.Context, creatorID uuid.UUID, name string, iconURL *string, memberIDs []uuid.UUID) (*domain.Conversation, error)
	GetConversationByID(ctx context.Context, conversationID, userID uuid.UUID) (*domain.Conversation, error)
	AddMember(ctx context.Context, conversationID, requesterID, newMemberID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error)
	RemoveMember(ctx context.Context, conversationID, requesterID, targetUserID uuid.UUID) (bool, error)
	UpdateReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error)
}

type conversationUsecase struct {
	convoRepo repository.ConversationRepository
	userRepo  repository.UserRepository
}

func NewConversationUsecase(convoRepo repository.ConversationRepository, userRepo repository.UserRepository) ConversationUsecase {
	return &conversationUsecase{
		convoRepo: convoRepo,
		userRepo:  userRepo,
	}
}

func (u *conversationUsecase) GetUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error) {
	return u.convoRepo.FindUserConversations(ctx, userID)
}

func (u *conversationUsecase) GetOrCreateDirectConversation(ctx context.Context, userID, partnerID uuid.UUID) (*domain.Conversation, error) {
	if userID == partnerID {
		return nil, ErrCannotChatSelf
	}

	partner, err := u.userRepo.FindByID(ctx, partnerID)
	if err != nil {
		return nil, err
	}
	if partner == nil {
		return nil, ErrTargetNotFound
	}

	existing, err := u.convoRepo.FindDirectConversation(ctx, userID, partnerID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}

	newConvo := &domain.Conversation{
		ID:        uuid.New(),
		Type:      domain.ConversationTypeDirect,
		CreatedBy: &userID,
	}

	members := []repository.MemberInput{
		{UserID: userID, Role: domain.MemberRoleMember},
		{UserID: partnerID, Role: domain.MemberRoleMember},
	}

	return u.convoRepo.Create(ctx, newConvo, members)
}

func (u *conversationUsecase) CreateGroupConversation(ctx context.Context, creatorID uuid.UUID, name string, iconURL *string, memberIDs []uuid.UUID) (*domain.Conversation, error) {
	if name == "" {
		return nil, errors.New("Tên nhóm chat không được để trống")
	}

	newConvo := &domain.Conversation{
		ID:        uuid.New(),
		Name:      &name,
		Type:      domain.ConversationTypeGroup,
		IconURL:   iconURL,
		CreatedBy: &creatorID,
	}

	memberMap := make(map[uuid.UUID]bool)
	memberMap[creatorID] = true

	members := []repository.MemberInput{
		{UserID: creatorID, Role: domain.MemberRoleAdmin},
	}

	for _, mid := range memberIDs {
		if !memberMap[mid] {
			memberMap[mid] = true
			members = append(members, repository.MemberInput{
				UserID: mid,
				Role:   domain.MemberRoleMember,
			})
		}
	}

	return u.convoRepo.Create(ctx, newConvo, members)
}

func (u *conversationUsecase) GetConversationByID(ctx context.Context, conversationID, userID uuid.UUID) (*domain.Conversation, error) {
	member, err := u.convoRepo.FindMember(ctx, conversationID, userID)
	if err != nil {
		return nil, err
	}
	if member == nil {
		return nil, ErrConversationNotFound
	}

	return u.convoRepo.FindByID(ctx, conversationID)
}

func (u *conversationUsecase) AddMember(ctx context.Context, conversationID, requesterID, newMemberID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error) {
	convo, err := u.convoRepo.FindByID(ctx, conversationID)
	if err != nil {
		return nil, err
	}
	if convo == nil {
		return nil, ErrConversationNotFound
	}

	reqMember, err := u.convoRepo.FindMember(ctx, conversationID, requesterID)
	if err != nil {
		return nil, err
	}
	if reqMember == nil {
		return nil, ErrNotMember
	}

	if convo.Type == domain.ConversationTypeDirect {
		return nil, errors.New("Không thể thêm thành viên vào cuộc hội thoại 1-1")
	}

	if role == "" {
		role = domain.MemberRoleMember
	}

	return u.convoRepo.AddMember(ctx, conversationID, newMemberID, role)
}

func (u *conversationUsecase) RemoveMember(ctx context.Context, conversationID, requesterID, targetUserID uuid.UUID) (bool, error) {
	convo, err := u.convoRepo.FindByID(ctx, conversationID)
	if err != nil {
		return false, err
	}
	if convo == nil {
		return false, ErrConversationNotFound
	}

	reqMember, err := u.convoRepo.FindMember(ctx, conversationID, requesterID)
	if err != nil {
		return false, err
	}
	if reqMember == nil {
		return false, ErrNotMember
	}

	// Người dùng tự rời nhóm
	if requesterID == targetUserID {
		return u.convoRepo.RemoveMember(ctx, conversationID, targetUserID)
	}

	// Quản trị viên xóa thành viên khác
	if reqMember.Role != domain.MemberRoleAdmin {
		return false, ErrNotAdmin
	}

	return u.convoRepo.RemoveMember(ctx, conversationID, targetUserID)
}

func (u *conversationUsecase) UpdateReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error) {
	member, err := u.convoRepo.FindMember(ctx, conversationID, userID)
	if err != nil {
		return nil, err
	}
	if member == nil {
		return nil, ErrNotMember
	}

	return u.convoRepo.UpsertReadReceipt(ctx, conversationID, userID, lastMessageSeen)
}
