package handler

import (
	"net/http"
	"trilha-api/internal/account/dto"
	usecase "trilha-api/internal/account/use_case"
	db "trilha-api/internal/shared/database/sqlc"
	sharedDto "trilha-api/internal/shared/dto"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	usecase *usecase.AccountUseCase
}

func New(uc *usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{usecase: uc}
}

func (h *AccountHandler) Register(c *gin.Context) {
	req := dto.CreateAccountRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model := db.Account{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.usecase.Register(&model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.AccountResponse{
		Default: sharedDto.Default{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Time,
			UpdatedAt: model.UpdatedAt.Time,
			DeletedAt: model.DeletedAt.Time,
		},
		Name:  model.Name,
		Email: model.Email,
	}

	c.JSON(http.StatusCreated, sharedDto.APIResponse[dto.AccountResponse]{
		Status: "success",
		Data:   res,
	})
}
