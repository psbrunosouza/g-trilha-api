package repository

import (
	"context"
	"errors"
	"testing"
	"trilha-api/internal/account/entity"
	db "trilha-api/internal/shared/database/sqlc"
	"trilha-api/internal/shared/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// DBMock é o nosso mock para a interface Querier.
// Ele nos permite simular o comportamento do banco de dados.
type DBMock struct {
	mock.Mock
}

// CreateAccount é a implementação mockada do método da interface.
func (m *DBMock) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	// m.Called(ctx, arg) registra que o método foi chamado com esses argumentos.
	args := m.Called(ctx, arg)
	// Retornamos os valores que configuramos no teste.
	return args.Get(0).(db.Account), args.Error(1)
}

func TestAccountRepository_Register(t *testing.T) {
	// Cenário 1: Registro com sucesso
	t.Run("should register a new account with success", func(t *testing.T) {
		// Arrange (Preparação)
		dbMock := new(DBMock)
		repo := New(dbMock)

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

		// Dizemos ao mock o que esperar e o que retornar
		dbMock.On("CreateAccount", context.Background(), createAccountParams).Return(expectedAccount, nil)

		// Act (Ação)
		err := repo.Register(account)

		// Assert (Verificação)
		assert.NoError(t, err)                          // Verificamos que não houve erro
		assert.Equal(t, expectedAccount.ID, account.ID) // Verificamos que o ID foi atualizado
		dbMock.AssertExpectations(t)                    // Verificamos se o mock foi chamado como esperado
	})

	// Cenário 2: Erro no registro
	t.Run("should return an error when fails to register a new account", func(t *testing.T) {
		// Arrange (Preparação)
		dbMock := new(DBMock)
		repo := New(dbMock)

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

		// Configuramos o mock para retornar um erro
		dbMock.On("CreateAccount", context.Background(), createAccountParams).Return(db.Account{}, errors.New("database error"))

		// Act (Ação)
		err := repo.Register(account)

		// Assert (Verificação)
		assert.Error(t, err)         // Verificamos que um erro foi retornado
		dbMock.AssertExpectations(t) // Verificamos se o mock foi chamado como esperado
	})
}
