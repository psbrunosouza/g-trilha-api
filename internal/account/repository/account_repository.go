package repository

import (
	"context"
	"fmt"
	"trilha-api/internal/account/entity"
	db "trilha-api/internal/shared/database/sqlc"
	"trilha-api/internal/shared/utils"
)

type AccountRepository struct {
	db db.Querier
}

type AccountRepositoryInterface interface {
	Register(account *entity.AccountEntity) error
}

func New(db db.Querier) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Register(account *entity.AccountEntity) error {
	fields := db.CreateAccountParams{
		Name:     account.Name,
		Email:    account.Email,
		Password: account.Password,
		Avatar:   utils.ToPgText(account.Avatar),
	}

	acc, err := r.db.CreateAccount(context.Background(), fields)

	if err != nil {
		return fmt.Errorf("erro ao registrar conta: %w", err)
	}

	*account = entity.AccountEntity{
		ID:        acc.ID,
		Name:      acc.Name,
		Email:     acc.Email,
		Password:  acc.Password,
		CreatedAt: acc.CreatedAt.Time,
		UpdatedAt: acc.UpdatedAt.Time,
		DeletedAt: &acc.DeletedAt.Time,
		Avatar:    acc.Avatar.String,
	}

	return nil
}
