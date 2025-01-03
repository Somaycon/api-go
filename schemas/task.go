package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
}

type TaskResponse struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
