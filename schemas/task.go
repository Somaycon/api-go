package schemas

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Id          int
	Name        string
	Description string
}
