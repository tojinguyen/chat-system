package domain

import (
	"time"

	"github.com/google/uuid"
)

type FriendStatus string

const (
	FriendStatusPending  FriendStatus = "pending"
	FriendStatusAccepted FriendStatus = "accepted"
	FriendStatusDeclined FriendStatus = "declined"
	FriendStatusBlocked  FriendStatus = "blocked"
)

type Friend struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	RequesterID uuid.UUID    `gorm:"type:uuid;not null;column:requesterId;index" json:"requesterId"`
	AddresseeID uuid.UUID    `gorm:"type:uuid;not null;column:addresseeId;index" json:"addresseeId"`
	Status      FriendStatus `gorm:"type:friend_status_enum;not null;default:'pending';column:status" json:"status"`
	CreatedAt   time.Time    `gorm:"column:createdAt;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time    `gorm:"column:updatedAt;not null;autoUpdateTime" json:"updatedAt"`

	Requester *User `gorm:"foreignKey:RequesterID" json:"requester,omitempty"`
	Addressee *User `gorm:"foreignKey:AddresseeID" json:"addressee,omitempty"`
}

func (Friend) TableName() string {
	return "friends"
}
