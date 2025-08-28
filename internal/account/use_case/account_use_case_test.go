package usecase_test

import (
	"testing"
	"trilha-api/internal/account/use_case"
	db "trilha-api/internal/shared/database/sqlc"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockAccountRepository struct {
	RegisterFunc func(account *db.Account) error
}

func (m *mockAccountRepository) Register(account *db.Account) error {
	return m.RegisterFunc(account)
}

func TestAccountUseCase_Register(t *testing.T) {
	mockRepo := &mockAccountRepository{}
	uc := usecase.New(mockRepo)

	account := &db.Account{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}

	mockRepo.RegisterFunc = func(acc *db.Account) error {
		assert.NotEqual(t, "password123", acc.Password)
		err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte("password123"))
		assert.NoError(t, err)
		return nil
	}

	err := uc.Register(account)

	assert.NoError(t, err)
}
