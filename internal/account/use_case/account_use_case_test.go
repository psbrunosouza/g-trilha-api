package usecase_test

import (
	"errors"
	"testing"
	"time"
	"trilha-api/internal/account/entity"
	"trilha-api/internal/account/mocks"
	usecase "trilha-api/internal/account/use_case"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAccountUseCase_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountRepositoryInterface(ctrl)
	uc := usecase.New(mockRepo)

	account := &entity.AccountEntity{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Avatar:   "test image",
	}

	mockRepo.EXPECT().Register(gomock.Any()).DoAndReturn(func(acc *entity.AccountEntity) error {
		assert.NotEqual(t, "password123", acc.Password)
		err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte("password123"))
		assert.NoError(t, err)
		return nil
	})

	err := uc.Register(account)

	assert.NoError(t, err)
}

func TestAccountUseCase_Find(t *testing.T) {
	t.Run("Should find an account by id with success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockAccountRepositoryInterface(ctrl)
		uc := usecase.New(mockRepo)

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

		mockRepo.EXPECT().Find(account).DoAndReturn(func(acc *entity.AccountEntity) error {
			*acc = *expectedAccount
			return nil
		})

		err := uc.Find(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedAccount, account)
	})

	t.Run("should return an error when account does not exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockAccountRepositoryInterface(ctrl)
		uc := usecase.New(mockRepo)

		accountId := uuid.New()

		account := &entity.AccountEntity{
			ID: accountId,
		}

		mockRepo.EXPECT().Find(account).Return(errors.New("account was not found"))

		err := uc.Find(account)

		assert.Error(t, err)
	})
}
