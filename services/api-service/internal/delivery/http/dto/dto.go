package dto

import (
	"time"

	"api-service/internal/domain"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username string  `json:"userName" binding:"required"`
	Password string  `json:"password" binding:"required,min=6"`
	Name     string  `json:"name" binding:"required"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

type LoginRequest struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Phone     *string   `json:"phone,omitempty"`
	Address   *string   `json:"address,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponse(u *domain.User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:        u.ID.String(),
		Username:  u.Username,
		Name:      u.Name,
		Phone:     u.Phone,
		Address:   u.Address,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type FriendResponse struct {
	ID          string        `json:"id"`
	RequesterID string        `json:"requesterId"`
	AddresseeID string        `json:"addresseeId"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Requester   *UserResponse `json:"requester,omitempty"`
	Addressee   *UserResponse `json:"addressee,omitempty"`
}

func ToFriendResponse(f *domain.Friend) *FriendResponse {
	if f == nil {
		return nil
	}
	return &FriendResponse{
		ID:          f.ID.String(),
		RequesterID: f.RequesterID.String(),
		AddresseeID: f.AddresseeID.String(),
		Status:      string(f.Status),
		CreatedAt:   f.CreatedAt,
		UpdatedAt:   f.UpdatedAt,
		Requester:   ToUserResponse(f.Requester),
		Addressee:   ToUserResponse(f.Addressee),
	}
}

type CreateDirectConversationRequest struct {
	PartnerID string `json:"partnerId" binding:"required"`
}

type CreateGroupConversationRequest struct {
	Name      string   `json:"name" binding:"required"`
	IconURL   *string  `json:"iconUrl"`
	MemberIDs []string `json:"memberIds"`
}

type AddMemberRequest struct {
	UserID string            `json:"userId" binding:"required"`
	Role   domain.MemberRole `json:"role"`
}

type UpdateReadReceiptRequest struct {
	LastMessageSeen string `json:"lastMessageSeen" binding:"required"`
}

type ConversationMemberResponse struct {
	ID             string        `json:"id"`
	ConversationID string        `json:"conversationId"`
	UserID         string        `json:"userId"`
	Role           string        `json:"role"`
	JoinedAt       time.Time     `json:"joinedAt"`
	User           *UserResponse `json:"user,omitempty"`
}

type ConversationResponse struct {
	ID        string                       `json:"id"`
	Name      *string                      `json:"name,omitempty"`
	Type      string                       `json:"type"`
	IconURL   *string                      `json:"iconUrl,omitempty"`
	CreatedBy *string                      `json:"createdBy,omitempty"`
	CreatedAt time.Time                    `json:"createdAt"`
	UpdatedAt time.Time                    `json:"updatedAt"`
	Members   []ConversationMemberResponse `json:"members"`
}

func ToConversationResponse(c *domain.Conversation) *ConversationResponse {
	if c == nil {
		return nil
	}
	var createdByStr *string
	if c.CreatedBy != nil && *c.CreatedBy != uuid.Nil {
		s := c.CreatedBy.String()
		createdByStr = &s
	}

	members := make([]ConversationMemberResponse, 0, len(c.Members))
	for _, m := range c.Members {
		members = append(members, ConversationMemberResponse{
			ID:             m.ID.String(),
			ConversationID: m.ConversationID.String(),
			UserID:         m.UserID.String(),
			Role:           string(m.Role),
			JoinedAt:       m.JoinedAt,
			User:           ToUserResponse(m.User),
		})
	}

	return &ConversationResponse{
		ID:        c.ID.String(),
		Name:      c.Name,
		Type:      string(c.Type),
		IconURL:   c.IconURL,
		CreatedBy: createdByStr,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Members:   members,
	}
}
