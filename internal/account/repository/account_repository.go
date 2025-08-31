package repository

import (
	"context"
	"fmt"
	"time"
	"trilha-api/internal/account/entity"
	db "trilha-api/internal/shared/database/sqlc"
	"trilha-api/internal/shared/utils"
)

type AccountRepository struct {
	db db.Querier
}

//go:generate mockgen -source=account_repository.go -destination=../mocks/account_repository_mock.go -package=mocks

type AccountRepositoryInterface interface {
	Register(account *entity.AccountEntity) error
	Find(account *entity.AccountEntity) error
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

	var deletedAt *time.Time
	if acc.DeletedAt.Valid {
		deletedAt = &acc.DeletedAt.Time
	}

	*account = entity.AccountEntity{
		ID:        acc.ID,
		Name:      acc.Name,
		Email:     acc.Email,
		Password:  acc.Password,
		CreatedAt: acc.CreatedAt.Time,
		UpdatedAt: acc.UpdatedAt.Time,
		DeletedAt: deletedAt,
		Avatar:    acc.Avatar.String,
	}

	return nil
}

func (r *AccountRepository) Find(account *entity.AccountEntity) error {
	acc, err := r.db.FindAccount(context.Background(), account.ID)

	if err != nil {
		return fmt.Errorf("account does not exists: %w", err)
	}

	var deletedAt *time.Time
	if acc.DeletedAt.Valid {
		deletedAt = &acc.DeletedAt.Time
	}

	*account = entity.AccountEntity{
		ID:        acc.ID,
		Name:      acc.Name,
		Email:     acc.Email,
		Password:  acc.Password,
		CreatedAt: acc.CreatedAt.Time,
		UpdatedAt: acc.UpdatedAt.Time,
		DeletedAt: deletedAt,
		Avatar:    acc.Avatar.String,
	}

	return nil
}
