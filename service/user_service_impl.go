package service

import (
	"mountainio/domain/entity"
	"mountainio/domain/model"
	"mountainio/repository"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *repository,
	}
}

func (service *userServiceImpl) RegisterUser(params model.RegisterUser, password string) (entity.User, error) {
	user := entity.User{
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: password,
	}
	result, err := service.UserRepository.Insert(user)
	return result, err
}
