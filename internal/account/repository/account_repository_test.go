package repository

import (
	"context"
	"errors"
	"testing"
	"time"
	"trilha-api/internal/account/entity"
	"trilha-api/internal/shared/database/mocks"
	db "trilha-api/internal/shared/database/sqlc"
	"trilha-api/internal/shared/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountRepository_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := mocks.NewMockQuerier(ctrl)
	repo := New(dbMock)

	t.Run("should register a new account with success", func(t *testing.T) {
		account := &entity.AccountEntity{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			Avatar:   "test image",
		}

		createAccountParams := db.CreateAccountParams{
			Name:     account.Name,
			Email:    account.Email,
			Password: account.Password,
			Avatar:   utils.ToPgText(account.Avatar),
		}

		expectedAccount := db.Account{
			ID:       uuid.New(),
			Name:     account.Name,
			Email:    account.Email,
			Password: account.Password,
		}

		dbMock.EXPECT().CreateAccount(context.Background(), createAccountParams).Return(expectedAccount, nil)

		err := repo.Register(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedAccount.ID, account.ID)
	})

	t.Run("should return an error when fails to register a new account", func(t *testing.T) {
		account := &entity.AccountEntity{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			Avatar:   "test image",
		}

		createAccountParams := db.CreateAccountParams{
			Name:     account.Name,
			Email:    account.Email,
			Password: account.Password,
			Avatar:   utils.ToPgText(account.Avatar),
		}

		dbMock.EXPECT().CreateAccount(context.Background(), createAccountParams).Return(db.Account{}, errors.New("database error"))

		err := repo.Register(account)

		assert.Error(t, err)
	})
}

func TestAccountRepository_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := mocks.NewMockQuerier(ctrl)
	repo := New(dbMock)

	t.Run("should return an account by id", func(t *testing.T) {
		account := &entity.AccountEntity{
			ID:        uuid.New(),
			Name:      "Gandalf O Branco",
			Email:     "gandalf@gmail.com",
			Password:  "123mudar",
			Avatar:    "url",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: func() *time.Time { t := time.Now(); return &t }(),
		}

		expectedAccount := db.FindAccountRow{
			ID:        account.ID,
			Name:      account.Name,
			Email:     account.Email,
			Password:  account.Password,
			Avatar:    utils.ToPgText(account.Avatar),
			CreatedAt: utils.TimeToPgTimestamp(&account.CreatedAt),
			UpdatedAt: utils.TimeToPgTimestamp(&account.UpdatedAt),
			DeletedAt: utils.TimeToPgTimestamp(account.DeletedAt),
		}

		dbMock.EXPECT().FindAccount(context.Background(), account.ID).Return(expectedAccount, nil)

		err := repo.Find(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedAccount.ID, account.ID)
	})

	t.Run("should return an error when account does not exist", func(t *testing.T) {
		account := &entity.AccountEntity{
			ID:        uuid.New(),
			Name:      "Gandalf O Branco",
			Email:     "gandalf@gmail.com",
			Password:  "123mudar",
			Avatar:    "url",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: func() *time.Time { t := time.Now(); return &t }(),
		}

		dbMock.EXPECT().FindAccount(context.Background(), account.ID).Return(db.FindAccountRow{}, errors.New("database error"))

		err := repo.Find(account)

		assert.Error(t, err)
	})
}
