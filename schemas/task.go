package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string
	Description string
	UserId      uint
}

type TaskResponse struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserId      uint      `json:"user_id"`
}
