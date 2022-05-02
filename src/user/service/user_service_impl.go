package service

import (
	"errors"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"mountainio/app/exception"
	"mountainio/app/helper"
	"mountainio/domain/entity"
	"mountainio/domain/model"
	"mountainio/src/user/repository"
	"mountainio/validation"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) RegisterUser(params model.RegisterUser) (model.RegisterUserResponse, error) {
	validation.ValidateRegisterUser(params)

	user := entity.User{
		ID:    uuid.New(),
		Name:  params.Name,
		Email: params.Email,
		Role:  "user",
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.MinCost)
	exception.PanicIfNeeded(err)

	// Insert password to Struct
	resetPasswordToken, err := gonanoid.New()
	exception.PanicIfNeeded(err)

	user.PasswordHash = string(passwordHash)
	user.ResetPasswordToken = resetPasswordToken

	result, err := service.UserRepository.Insert(user)

	response := model.RegisterUserResponse{
		Name:      result.Name,
		Email:     result.Email,
		Role:      result.Role,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return response, err
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
