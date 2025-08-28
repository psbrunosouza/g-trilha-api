package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"trilha-api/internal/account/dto"
	"trilha-api/internal/account/handler"
	usecase "trilha-api/internal/account/use_case"
	db "trilha-api/internal/shared/database/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// mockAccountUseCase é o nosso mock para a interface AccountUseCaseInterface.
// Ele nos permite simular o comportamento do use case em nossos testes.
type mockAccountUseCase struct {
	RegisterFunc func(account *db.Account) error
}

// Register é a implementação mockada do método da interface.
func (m *mockAccountUseCase) Register(account *db.Account) error {
	if m.RegisterFunc != nil {
		return m.RegisterFunc(account)
	}
	return nil
}

// Garante que nosso mock implementa a interface em tempo de compilação.
var _ usecase.AccountUseCaseInterface = (*mockAccountUseCase)(nil)

func TestAccountHandler_Register(t *testing.T) {
	// Configura o Gin para o modo de teste
	gin.SetMode(gin.TestMode)

	// Cenário 1: Registro com sucesso
	t.Run("should return status 201 and the created account on success", func(t *testing.T) {
		// Arrange (Preparação)
		mockUseCase := &mockAccountUseCase{}
		accountID := uuid.New()

		// Configuramos o mock para retornar sucesso
		mockUseCase.RegisterFunc = func(account *db.Account) error {
			// Simulamos que o use case preenche o ID do modelo
			account.ID = accountID
			return nil
		}

		h := handler.New(mockUseCase)
		router := gin.Default()
		router.POST("/register", h.Register)

		// Criamos o corpo da requisição
		createAccountReq := dto.CreateAccountRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(createAccountReq)

		// Criamos a requisição HTTP
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		// Act (Ação)
		router.ServeHTTP(w, req)

		// Assert (Verificação)
		assert.Equal(t, http.StatusCreated, w.Code)

		// Verificamos também o corpo da resposta
		var responseBody map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.Equal(t, "success", responseBody["status"])

		data := responseBody["data"].(map[string]interface{})
		assert.Equal(t, accountID.String(), data["id"])
		assert.Equal(t, createAccountReq.Name, data["name"])
		assert.Equal(t, createAccountReq.Email, data["email"])
	})

	// Cenário 2: Erro de validação (JSON inválido)
	t.Run("should return status 400 for invalid json body", func(t *testing.T) {
		// Arrange
		mockUseCase := &mockAccountUseCase{} // O mock não será chamado
		h := handler.New(mockUseCase)
		router := gin.Default()
		router.POST("/register", h.Register)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("invalid-json")))
		req.Header.Set("Content-Type", "application/json")

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Cenário 3: Erro retornado pelo use case
	t.Run("should return status 500 when use case returns an error", func(t *testing.T) {
		// Arrange
		mockUseCase := &mockAccountUseCase{}
		expectedError := "database connection failed"

		// Configuramos o mock para retornar um erro
		mockUseCase.RegisterFunc = func(account *db.Account) error {
			return errors.New(expectedError)
		}

		h := handler.New(mockUseCase)
		router := gin.Default()
		router.POST("/register", h.Register)

		createAccountReq := dto.CreateAccountRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(createAccountReq)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Verificamos a mensagem de erro
		var responseBody map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedError, responseBody["error"])
	})
}
