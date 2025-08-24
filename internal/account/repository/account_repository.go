package repository

import (
	"context"
	"fmt"
	db "trilha-api/internal/shared/database/sqlc"
)

type AccountRepository struct {
	db db.Querier
}

type AccountRepositoryInterface interface {
	Register(account *db.Account) error
}

func New(db db.Querier) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Register(account *db.Account) error {
	fields := db.CreateAccountParams{
		Name:     account.Name,
		Email:    account.Email,
		Password: account.Password,
	}

	acc, err := r.db.CreateAccount(context.Background(), fields)

	if err != nil {
		return fmt.Errorf("erro ao registrar conta: %w", err)
	}

	*account = db.Account(acc)

	return nil
}
