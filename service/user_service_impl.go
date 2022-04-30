package service

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mountainio/app/helper"
	"mountainio/domain/entity"
	"mountainio/domain/model"
	"mountainio/repository"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) RegisterUser(params model.RegisterUser) (entity.User, error) {
	user := entity.User{
		ID:    uuid.New(),
		Name:  params.Name,
		Email: params.Email,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	// Insert password to hash
	user.PasswordHash = string(passwordHash)

	result, err := service.UserRepository.Insert(user)
	return result, err
}

func (service *userServiceImpl) FindUserByID(id string) (entity.User, error) {
	userID := helper.ConvertUUID(id)
	user, err := service.UserRepository.FindByID(userID)

	if user.ID == helper.CheckNilDataFromUUID() {
		return user, errors.New("Data Not Found!")
	}

	return user, err
}

func (service *userServiceImpl) FindUserByEmail(email string) (entity.User, error) {
	return service.UserRepository.FindByEmail(email)
}
