package todo

import (
	"time"
)

// Todo represents the core domain model for a todo item.
type Todo struct {
	ID          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	Name        string
	Description string
}
