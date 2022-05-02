package service

import (
	"mountainio/app/constant"
	"mountainio/app/middleware"
	"mountainio/domain/entity"
	"mountainio/domain/model"
	"time"
)

type authServiceImpl struct {
}

func NewAuthService() AuthService {
	return &authServiceImpl{}
}

func (service *authServiceImpl) GenerateTokenAuth(user entity.User) (model.LoginSuccess, error) {
	auth := model.AuthClaim{
		UserID:  user.ID,
		Role:    user.Role,
		Expired: time.Now().Add(time.Hour * constant.AuthExpired).Unix(),
	}

	token, err := middleware.GenerateToken(auth)
	res := model.LoginSuccess{
		UserID: user.ID,
		Email:  user.Email,
		Token:  token,
	}
	return res, err
}
