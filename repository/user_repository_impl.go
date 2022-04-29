package repository

import (
	"gorm.io/gorm"
	"mountainio/app/config"
	"mountainio/domain/entity"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (repository *userRepositoryImpl) Insert(user entity.User) (entity.User, error) {
	ctx, cancel := config.DBContext(10)
	defer cancel()

	tx := repository.db.WithContext(ctx)
	err := tx.Create(&user).Error
	return user, err
}
