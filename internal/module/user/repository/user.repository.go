package repository

import (
	"github.com/HiBang15/sample-gorm.git/internal/database"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.Connection}
}
func (userRepo *UserRepository) Create(user *entities.User) error {
	result := userRepo.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
