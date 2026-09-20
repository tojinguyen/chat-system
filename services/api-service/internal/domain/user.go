package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Username     string         `gorm:"size:50;uniqueIndex;not null;column:username" json:"username"`
	PasswordHash string         `gorm:"not null;column:passwordHash" json:"-"`
	Name         string         `gorm:"not null;column:name" json:"name"`
	Phone        *string        `gorm:"size:20;column:phone" json:"phone,omitempty"`
	Address      *string        `gorm:"column:address" json:"address,omitempty"`
	CreatedAt    time.Time      `gorm:"column:createdAt;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updatedAt;not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deletedAt;index" json:"-"`
}

func (User) TableName() string {
	return "users"
}
