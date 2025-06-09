package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}

// UserProfile represents additional user information
type UserProfile struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Height    float64        `json:"height"`
	Weight    float64        `json:"weight"`
	BodyType  string         `json:"body_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (profile *UserProfile) BeforeCreate(tx *gorm.DB) error {
	profile.ID = uuid.New()
	return nil
}
