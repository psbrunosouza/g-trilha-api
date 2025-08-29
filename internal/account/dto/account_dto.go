package dto

import (
	"trilha-api/internal/shared/dto"

	"github.com/google/uuid"
)

type AccountResponse struct {
	dto.Default
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type CreateAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Avatar   string `json:"avatar"`
}

type UpdadeAccountRequest struct {
	ID       uuid.UUID `json:id`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	Avatar   string    `json:"avatar"`
}
