package handler

import (
	"net/http"
	"trilha-api/internal/account/dto"
	"trilha-api/internal/account/entity"
	usecase "trilha-api/internal/account/use_case"
	sharedDto "trilha-api/internal/shared/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusBadRequest, sharedDto.APIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	model := entity.AccountEntity{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
	}

	if err := h.usecase.Register(&model); err != nil {
		c.JSON(http.StatusInternalServerError, sharedDto.APIResponse[any]{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	res := dto.AccountResponse{
		Default: sharedDto.Default{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt,
		},
		Name:  model.Name,
		Email: model.Email,
	}

	c.JSON(http.StatusCreated, sharedDto.APIResponse[dto.AccountResponse]{
		Status: http.StatusCreated,
		Data:   res,
	})
}

func (h *AccountHandler) Find(c *gin.Context) {
	id := c.Param("id")

	accountId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, sharedDto.APIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: "Invalid account ID",
		})
		return
	}

	account := &entity.AccountEntity{
		ID: accountId,
	}

	err = h.usecase.Find(account)

	if err != nil {
		c.JSON(http.StatusNotFound, sharedDto.APIResponse[any]{
			Status:  http.StatusNotFound,
			Message: "Account not found",
		})
		return
	}

	c.JSON(http.StatusOK, sharedDto.APIResponse[dto.AccountResponse]{
		Status: http.StatusOK,
		Data: dto.AccountResponse{
			Default: sharedDto.Default{
				ID:        account.ID,
				CreatedAt: account.CreatedAt,
				UpdatedAt: account.UpdatedAt,
				DeletedAt: account.DeletedAt,
			},
			Name:   account.Name,
			Email:  account.Email,
			Avatar: account.Avatar,
		},
	})
}
