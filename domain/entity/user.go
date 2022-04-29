package entity

import (
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
	"time"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name               string
	Email              string `gorm:"uniqueIndex"`
	AvatarFileName     string
	PasswordHash       string
	ResetPasswordToken string
	Role               string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
	IsDeleted          soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
