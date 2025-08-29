package handler

import (
	"net/http"
	"trilha-api/internal/account/dto"
	"trilha-api/internal/account/entity"
	usecase "trilha-api/internal/account/use_case"
	sharedDto "trilha-api/internal/shared/dto"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	usecase usecase.AccountUseCaseInterface
}

func New(uc usecase.AccountUseCaseInterface) *AccountHandler {
	return &AccountHandler{usecase: uc}
}

func (h *AccountHandler) Register(c *gin.Context) {
	req := dto.CreateAccountRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model := entity.AccountEntity{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
	}

	if err := h.usecase.Register(&model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.AccountResponse{
		Default: sharedDto.Default{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Name:  model.Name,
		Email: model.Email,
	}

	if model.DeletedAt != nil {
		res.DeletedAt = *model.DeletedAt
	}

	c.JSON(http.StatusCreated, sharedDto.APIResponse[dto.AccountResponse]{
		Status: "success",
		Data:   res,
	})
}
