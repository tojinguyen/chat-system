package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"api-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MemberInput struct {
	UserID uuid.UUID
	Role   domain.MemberRole
}

type ConversationRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Conversation, error)
	FindDirectConversation(ctx context.Context, userA, userB uuid.UUID) (*domain.Conversation, error)
	FindUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error)
	Create(ctx context.Context, convo *domain.Conversation, members []MemberInput) (*domain.Conversation, error)
	AddMember(ctx context.Context, conversationID, userID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error)
	RemoveMember(ctx context.Context, conversationID, userID uuid.UUID) (bool, error)
	FindMember(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationMember, error)
	UpsertReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error)
	GetReadReceipt(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationReadReceipt, error)
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Conversation, error) {
	var convo domain.Conversation
	err := r.db.WithContext(ctx).
		Preload("Members.User").
		Where("id = ?", id).
		First(&convo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find conversation by id: %w", err)
	}
	return &convo, nil
}

func (r *conversationRepository) FindDirectConversation(ctx context.Context, userA, userB uuid.UUID) (*domain.Conversation, error) {
	var convoID uuid.UUID
	// Tìm conversation ID loại 'direct' mà cả userA và userB đều là members
	query := `
		SELECT c.id FROM conversations c
		JOIN conversation_members m1 ON m1."conversationId" = c.id AND m1."userId" = ?
		JOIN conversation_members m2 ON m2."conversationId" = c.id AND m2."userId" = ?
		WHERE c.type = 'direct'
		LIMIT 1
	`
	err := r.db.WithContext(ctx).Raw(query, userA, userB).Scan(&convoID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to check direct conversation: %w", err)
	}
	if convoID == uuid.Nil {
		return nil, nil
	}
	return r.FindByID(ctx, convoID)
}

func (r *conversationRepository) FindUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error) {
	var convos []domain.Conversation
	subQuery := r.db.Model(&domain.ConversationMember{}).
		Select("\"conversationId\"").
		Where("\"userId\" = ?", userID)

	err := r.db.WithContext(ctx).
		Preload("Members.User").
		Where("id IN (?)", subQuery).
		Order("\"updatedAt\" DESC").
		Find(&convos).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user conversations: %w", err)
	}
	return convos, nil
}

func (r *conversationRepository) Create(ctx context.Context, convo *domain.Conversation, members []MemberInput) (*domain.Conversation, error) {
	if convo.ID == uuid.Nil {
		convo.ID = uuid.New()
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(convo).Error; err != nil {
			return err
		}

		for _, m := range members {
			memEntity := domain.ConversationMember{
				ID:             uuid.New(),
				ConversationID: convo.ID,
				UserID:         m.UserID,
				Role:           m.Role,
			}
			if err := tx.Create(&memEntity).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create conversation with members: %w", err)
	}

	return r.FindByID(ctx, convo.ID)
}

func (r *conversationRepository) AddMember(ctx context.Context, conversationID, userID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error) {
	var member domain.ConversationMember
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		First(&member).Error
	if err == nil {
		return &member, nil
	}

	newMember := domain.ConversationMember{
		ID:             uuid.New(),
		ConversationID: conversationID,
		UserID:         userID,
		Role:           role,
	}
	if err := r.db.WithContext(ctx).Create(&newMember).Error; err != nil {
		return nil, fmt.Errorf("failed to add member: %w", err)
	}

	_ = r.db.WithContext(ctx).Preload("User").Where("id = ?", newMember.ID).First(&newMember)
	return &newMember, nil
}

func (r *conversationRepository) RemoveMember(ctx context.Context, conversationID, userID uuid.UUID) (bool, error) {
	res := r.db.WithContext(ctx).
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		Delete(&domain.ConversationMember{})
	return res.RowsAffected > 0, res.Error
}

func (r *conversationRepository) FindMember(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationMember, error) {
	var member domain.ConversationMember
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		First(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find conversation member: %w", err)
	}
	return &member, nil
}

func (r *conversationRepository) UpsertReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error) {
	receipt := domain.ConversationReadReceipt{
		ConversationID:  conversationID,
		UserID:          userID,
		LastMessageSeen: lastMessageSeen,
		LastReadAt:      time.Now().UTC(),
	}

	err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "conversationId"}, {Name: "userId"}},
		DoUpdates: clause.AssignmentColumns([]string{"lastMessageSeen", "lastReadAt"}),
	}).Create(&receipt).Error

	if err != nil {
		return nil, fmt.Errorf("failed to upsert read receipt: %w", err)
	}
	return &receipt, nil
}

func (r *conversationRepository) GetReadReceipt(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationReadReceipt, error) {
	var receipt domain.ConversationReadReceipt
	err := r.db.WithContext(ctx).
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		First(&receipt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get read receipt: %w", err)
	}
	return &receipt, nil
}
