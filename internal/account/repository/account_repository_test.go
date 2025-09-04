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

func setup(t *testing.T) (*mocks.MockQuerier, *AccountRepository) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := mocks.NewMockQuerier(ctrl)
	repo := New(dbMock)

	return dbMock, repo
}

func TestAccountRepository_Register(t *testing.T) {
	dbMock, repo := setup(t)

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
	dbMock, repo := setup(t)

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

func TestAccountRepository_FindByEmail(t *testing.T) {
	dbMock, repo := setup(t)

	t.Run("should find an account by email and return account", func(t *testing.T) {
		now := time.Now()

		account := &entity.AccountEntity{
			Email: "gandalf@lor.com.br",
		}

		expectedAccount := db.FindAccountByEmailRow{
			Email:     account.Email,
			ID:        uuid.New(),
			Name:      "gandalf",
			Password:  "gandalf123",
			Avatar:    utils.ToPgText("url"),
			CreatedAt: utils.TimeToPgTimestamp(&now),
			UpdatedAt: utils.TimeToPgTimestamp(&now),
		}

		dbMock.EXPECT().FindAccountByEmail(context.Background(), account.Email).Return(expectedAccount, nil)

		err := repo.FindByEmail(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedAccount.Email, account.Email)
		assert.Equal(t, expectedAccount.Avatar.String, account.Avatar)
		assert.Equal(t, expectedAccount.ID, account.ID)
		assert.Equal(t, expectedAccount.Name, account.Name)
		assert.Equal(t, expectedAccount.Password, account.Password)
		assert.Equal(t, expectedAccount.CreatedAt.Time, account.CreatedAt)
		assert.Equal(t, expectedAccount.UpdatedAt.Time, account.UpdatedAt)
	})

	t.Run("should return an error when account not found", func(t *testing.T) {
		account := &entity.AccountEntity{
			Email: "gandalf@lor.com.br",
		}

		dbMock.EXPECT().FindAccountByEmail(context.Background(), account.Email).Return(db.FindAccountByEmailRow{}, errors.New("database error"))

		err := repo.FindByEmail(account)

		assert.Error(t, err)
	})
}
