package userservice

import "gorm.io/gorm"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	gorm.Model
}
