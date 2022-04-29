package service

import (
	"mountainio/domain/entity"
	"mountainio/domain/model"
)

type UserService interface {
	RegisterUser(params model.RegisterUser, password string) (entity.User, error)
}
