package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
