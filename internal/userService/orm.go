package userservice

import (
	"golang/pet_project/internal/tasksService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint                `json:"user_id"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Tasks    []tasksService.Task `json:"tasks"`
}
