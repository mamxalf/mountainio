package repository

import (
	"github.com/google/uuid"
	"mountainio/domain/entity"
)

type UserRepository interface {
	Insert(user entity.User) (entity.User, error)
	FindByID(id uuid.UUID) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
