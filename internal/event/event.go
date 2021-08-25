package event

import (
	"time"

	"gorm.io/gorm"
)

//Event model
type Event struct {
	ID          string         `gorm:"size:40;primaryKey" json:"id"`
	UserID      string         `json:"user_id"`
	Title       string         `gorm:"size:70" json:"title"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	Days        int64          `gorm:"-" json:"days"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
