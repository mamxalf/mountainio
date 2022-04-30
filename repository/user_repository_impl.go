package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mountainio/app/config"
	"mountainio/domain/entity"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (repository *userRepositoryImpl) Insert(user entity.User) (entity.User, error) {
	ctx, cancel := config.DBContext(10)
	defer cancel()

	tx := repository.db.WithContext(ctx)
	err := tx.Create(&user).Error
	return user, err
}

func (repository *userRepositoryImpl) FindByID(id uuid.UUID) (entity.User, error) {
	var user entity.User

	//ctx, cancel := config.DBContext(10)
	//defer cancel()
	//
	//tx := repository.db.WithContext(ctx)
	//err := tx.Where("id = ?", id).Find(&user).Error
	err := repository.db.Where("id = ?", id).Find(&user).Error
	return user, err
}

func (repository *userRepositoryImpl) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	ctx, cancel := config.DBContext(10)
	defer cancel()

	tx := repository.db.WithContext(ctx)
	err := tx.Where("email = ?", email).Find(&user).Error
	//err := repository.db.Where("id = ?", id).Find(&user).Error

	return user, err
}
