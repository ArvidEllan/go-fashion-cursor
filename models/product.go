package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product represents a clothing item in the system
type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Category    string         `gorm:"not null" json:"category"`
	Brand       string         `json:"brand"`
	ImageURL    string         `json:"image_url"`
	Sizes       []Size         `gorm:"many2many:product_sizes;" json:"sizes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Size represents available sizes for products
type Size struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (product *Product) BeforeCreate(tx *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}


// BeforeCreate will set a UUID rather than numeric ID
func (size *Size) BeforeCreate(tx *gorm.DB) error {
	size.ID = uuid.New()
	return nil
}
