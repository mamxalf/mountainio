package repository

import "mountainio/domain/entity"

type UserRepository interface {
	Insert(user entity.User) (entity.User, error)
}
