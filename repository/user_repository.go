package repository

import (
	"gorm.io/gorm"
	"trivial-todo-be/model"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}
