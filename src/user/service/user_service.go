package service

import (
	"mountainio/domain/entity"
	"mountainio/domain/model"
)

type UserService interface {
	RegisterUser(params model.RegisterUser) (entity.User, error)
	FindUserByID(id string) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
}
