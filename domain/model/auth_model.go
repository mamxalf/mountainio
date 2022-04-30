package model

import "github.com/google/uuid"

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginSuccess struct {
	UserID uuid.UUID
	Email  string
	Token  string
}

type AuthClaim struct {
	UserID  uuid.UUID
	Role    string
	Expired int64
}
