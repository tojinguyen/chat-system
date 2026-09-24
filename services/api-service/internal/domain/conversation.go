package domain

import (
	"time"

	"github.com/google/uuid"
)

type ConversationType string

const (
	ConversationTypeDirect ConversationType = "direct"
	ConversationTypeGroup  ConversationType = "group"
	ConversationTypeClan   ConversationType = "clan"
	ConversationTypeRegion ConversationType = "region"
	ConversationTypeGlobal ConversationType = "global"
)

type MemberRole string

const (
	MemberRoleAdmin  MemberRole = "admin"
	MemberRoleMember MemberRole = "member"
)

type Conversation struct {
	ID        uuid.UUID        `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Name      *string          `gorm:"size:100;column:name" json:"name,omitempty"`
	Type      ConversationType `gorm:"type:conversation_type_enum;not null;default:'direct';column:type" json:"type"`
	IconURL   *string          `gorm:"column:iconUrl" json:"iconUrl,omitempty"`
	CreatedBy *uuid.UUID       `gorm:"type:uuid;column:createdBy" json:"createdBy,omitempty"`
	CreatedAt time.Time        `gorm:"column:createdAt;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time        `gorm:"column:updatedAt;not null;autoUpdateTime" json:"updatedAt"`

	Members []ConversationMember `gorm:"foreignKey:ConversationID" json:"members,omitempty"`
}

func (Conversation) TableName() string {
	return "conversations"
}

type ConversationMember struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	ConversationID uuid.UUID  `gorm:"type:uuid;not null;column:conversationId;index" json:"conversationId"`
	UserID         uuid.UUID  `gorm:"type:uuid;not null;column:userId;index" json:"userId"`
	Role           MemberRole `gorm:"type:member_role_enum;not null;default:'member';column:role" json:"role"`
	JoinedAt       time.Time  `gorm:"column:joinedAt;not null;autoCreateTime" json:"joinedAt"`

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (ConversationMember) TableName() string {
	return "conversation_members"
}

type ConversationReadReceipt struct {
	ConversationID  uuid.UUID `gorm:"type:uuid;primaryKey;column:conversationId" json:"conversationId"`
	UserID          uuid.UUID `gorm:"type:uuid;primaryKey;column:userId" json:"userId"`
	LastMessageSeen string    `gorm:"size:64;not null;column:lastMessageSeen" json:"lastMessageSeen"`
	LastReadAt      time.Time `gorm:"column:lastReadAt;not null;autoCreateTime" json:"lastReadAt"`
}

func (ConversationReadReceipt) TableName() string {
	return "conversation_read_receipts"
}
