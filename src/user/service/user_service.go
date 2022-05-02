package service

import (
	"mountainio/domain/entity"
	"mountainio/domain/model"
)

type UserService interface {
	RegisterUser(params model.RegisterUser) (model.UserResponse, error)
	FindUserByID(id string) (model.UserResponse, error)
	FindUserByEmail(email string) (entity.User, error)
}
