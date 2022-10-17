package services

import (
	"errors"
	"github.com/HiBang15/sample-gorm.git/constant"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/dto"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/entities"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/repository"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/transformers"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo        repository.UserRepository
	UserTransformer *transformers.UserTransformer
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo:        *repository.NewUserRepository(),
		UserTransformer: transformers.NewUserTransformer(),
	}
}

func (userService *UserService) CreateUser(
	request *dto.CreateUserRequest,
) (user *dto.User, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(constant.ErrHashCode)
	}
	request.Password = string(hashedPassword)

	// create verify code

	userEntities := &entities.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Password:    string(hashedPassword),
		PhoneNumber: request.PhoneNumber,
	}

	err = userService.UserRepo.Create(userEntities)
	if err != nil {
		if err.Error() == constant.ErrDuplicateEmailMessage {
			return nil, errors.New(constant.ErrorEmailExist)
		}
		return nil, errors.New(constant.ErrCreateUserFail)
	}

	return userService.UserTransformer.UserEntityToDto(userEntities), nil
}
