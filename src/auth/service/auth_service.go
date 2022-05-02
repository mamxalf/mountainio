package service

import (
	"mountainio/domain/entity"
	"mountainio/domain/model"
)

type AuthService interface {
	GenerateTokenAuth(user entity.User) (model.LoginSuccess, error)
}
