package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TryOn represents a virtual try-on session
type TryOn struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	ProductID     uuid.UUID      `gorm:"type:uuid;not null" json:"product_id"`
	OriginalImage string         `gorm:"not null" json:"original_image"`
	ResultImage   string         `json:"result_image"`
	Status        string         `gorm:"not null" json:"status"` // pending, processing, completed, failed
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (tryon *TryOn) BeforeCreate(tx *gorm.DB) error {
	tryon.ID = uuid.New()
	return nil
}

// TryOnHistory represents the history of try-on sessions for a user
type TryOnHistory struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	TryOnID   uuid.UUID      `gorm:"type:uuid;not null" json:"try_on_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (history *TryOnHistory) BeforeCreate(tx *gorm.DB) error {
	history.ID = uuid.New()
	return nil
}
