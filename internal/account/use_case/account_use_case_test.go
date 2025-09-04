package usecase_test

import (
	"errors"
	"testing"
	"time"
	"trilha-api/internal/account/entity"
	"trilha-api/internal/account/mocks"
	usecase "trilha-api/internal/account/use_case"
	db "trilha-api/internal/shared/database/sqlc"
	"trilha-api/internal/shared/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func setup(t *testing.T) (*mocks.MockAccountRepositoryInterface, *usecase.AccountUseCase) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mocks.NewMockAccountRepositoryInterface(ctrl)
	uc := usecase.New(mock)

	return mock, uc
}

func TestAccountUseCase_Register(t *testing.T) {
	mock, uc := setup(t)

	account := &entity.AccountEntity{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Avatar:   "test image",
	}

	mock.EXPECT().Register(gomock.Any()).DoAndReturn(func(acc *entity.AccountEntity) error {
		assert.NotEqual(t, "password123", acc.Password)
		err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte("password123"))
		assert.NoError(t, err)
		return nil
	})

	err := uc.Register(account)

	assert.NoError(t, err)
}

func TestAccountUseCase_Find(t *testing.T) {
	mock, uc := setup(t)

	t.Run("Should find an account by id with success", func(t *testing.T) {

		accountId := uuid.New()
		now := time.Now()

		account := &entity.AccountEntity{
			ID: accountId,
		}

		expectedAccount := &entity.AccountEntity{
			ID:        accountId,
			Name:      "Gandalf",
			Email:     "gandalf@lor.com.br",
			Password:  "123mago",
			Avatar:    "url",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mock.EXPECT().Find(account).DoAndReturn(func(acc *entity.AccountEntity) error {
			*acc = *expectedAccount
			return nil
		})

		err := uc.Find(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedAccount, account)
	})

	t.Run("should return an error when account not found", func(t *testing.T) {

		accountId := uuid.New()

		account := &entity.AccountEntity{
			ID: accountId,
		}

		mock.EXPECT().Find(account).Return(errors.New("account not found"))

		err := uc.Find(account)

		assert.Error(t, err)
	})
}

func TestAccountUseCase_FindByEmail(t *testing.T) {
	mock, uc := setup(t)

	accountId := uuid.New()
	now := time.Now()

	account := &entity.AccountEntity{
		Email: "gandalf@lor.com.br",
	}

	expectedAccount := db.FindAccountByEmailRow{
		Email:     account.Email,
		ID:        accountId,
		Name:      "Gandalf",
		Avatar:    utils.ToPgText("url"),
		Password:  "gandalf123",
		CreatedAt: utils.TimeToPgTimestamp(&now),
		UpdatedAt: utils.TimeToPgTimestamp(&now),
	}

	t.Run("should return an account by email", func(t *testing.T) {
		mock.EXPECT().FindByEmail(account).DoAndReturn(func(acc *entity.AccountEntity) error {
			*acc = entity.AccountEntity{
				Email:     expectedAccount.Email,
				ID:        expectedAccount.ID,
				Name:      expectedAccount.Name,
				Avatar:    expectedAccount.Avatar.String,
				Password:  expectedAccount.Password,
				CreatedAt: expectedAccount.CreatedAt.Time,
				UpdatedAt: expectedAccount.UpdatedAt.Time,
			}
			return nil
		})

		err := uc.FindByEmail(account)

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
		mock.EXPECT().FindByEmail(account).Return(errors.New("account not found"))

		err := uc.FindByEmail(account)

		assert.Error(t, err)
	})
}
