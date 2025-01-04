package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	DataNasc string
	Password string
	Task     []Task
}

type UserResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	DataNasc  string    `json:"dataNasc"`
	Password  string    `json:"password"`
	Task      []Task    `json:"task" gorm:"foreignKey:UserId"`
}
