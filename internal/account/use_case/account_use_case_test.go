package usecase_test

import (
	"testing"
	"time"
	"trilha-api/internal/account/entity"
	usecase "trilha-api/internal/account/use_case"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockAccountRepository struct {
	RegisterFunc func(account *entity.AccountEntity) error
	FindFunc     func(account *entity.AccountEntity) error
}

func (m *mockAccountRepository) Register(account *entity.AccountEntity) error {
	return m.RegisterFunc(account)
}

func (m *mockAccountRepository) Find(account *entity.AccountEntity) error {
	return m.FindFunc(account)
}

func TestAccountUseCase_Register(t *testing.T) {
	mockRepo := &mockAccountRepository{}
	uc := usecase.New(mockRepo)

	account := &entity.AccountEntity{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Avatar:   "test image",
	}

	mockRepo.RegisterFunc = func(acc *entity.AccountEntity) error {
		assert.NotEqual(t, "password123", acc.Password)
		err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte("password123"))
		assert.NoError(t, err)
		return nil
	}

	err := uc.Register(account)

	assert.NoError(t, err)
}

func TestAccountUseCase_Find(t *testing.T) {
	mockRepo := &mockAccountRepository{}
	uc := usecase.New(mockRepo)

	account := &entity.AccountEntity{
		ID:        uuid.New(),
		Name:      "Gandalf",
		Email:     "gandalf@lor.com",
		Password:  "mago123",
		Avatar:    "http://my_huge_staff.com/3121034",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: func() *time.Time { t := time.Now(); return &t }(),
	}

	err := uc.Find(account)

	mockRepo.FindFunc = func(acc *entity.AccountEntity) error {
		assert.NotNil(t, acc.ID)
		return nil
	}

	assert.NoError(t, err)
}
