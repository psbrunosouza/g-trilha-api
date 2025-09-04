package handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"trilha-api/internal/account/dto"
	"trilha-api/internal/account/entity"
	"trilha-api/internal/account/handler"
	"trilha-api/internal/account/mocks"
	sharedDto "trilha-api/internal/shared/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*gin.Engine, *mocks.MockAccountUseCaseInterface) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockAccountUseCaseInterface(ctrl)
	h := handler.New(mock)
	router := gin.Default()

	router.POST("/api/v1/accounts", h.Register)
	router.GET("/api/v1/accounts/:id", h.Find)
	router.GET("/api/v1/accounts/find_by_email/:email", h.FindByEmail)

	return router, mock
}

func TestAccountHandler_Register(t *testing.T) {

	router, mockUseCase := setup(t)

	t.Run("should return status 201 and the created account on success", func(t *testing.T) {

		accountID := uuid.New()
		createAccountReq := dto.CreateAccountRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
			Avatar:   "test",
		}

		mockUseCase.EXPECT().Register(gomock.Any()).DoAndReturn(func(account *entity.AccountEntity) error {
			account.ID = accountID
			return nil
		})

		body, _ := json.Marshal(createAccountReq)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/accounts", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody sharedDto.APIResponse[dto.AccountResponse]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, responseBody.Status)
		assert.Equal(t, accountID, responseBody.Data.ID)
		assert.Equal(t, createAccountReq.Name, responseBody.Data.Name)
		assert.Equal(t, createAccountReq.Email, responseBody.Data.Email)
	})

	t.Run("should return status 400 for invalid json body", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/accounts", bytes.NewBuffer([]byte("invalid-json")))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return status 500 when use case returns an error", func(t *testing.T) {
		expectedError := "database connection failed"
		mockUseCase.EXPECT().Register(gomock.Any()).Return(errors.New(expectedError))

		createAccountReq := dto.CreateAccountRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
			Avatar:   "test",
		}
		body, _ := json.Marshal(createAccountReq)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/accounts", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var responseBody sharedDto.APIResponse[any]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, responseBody.Status)
		assert.Equal(t, expectedError, responseBody.Message)
	})
}

func TestAccountHandler_Find(t *testing.T) {

	router, mockUseCase := setup(t)

	t.Run("should return status 200 and account by id", func(t *testing.T) {

		accountID := uuid.New()
		now := time.Now()
		expectedAccount := &entity.AccountEntity{
			ID:        accountID,
			Name:      "Gandalf",
			Email:     "gandalf@lor.com.br",
			Password:  "123mago",
			Avatar:    "url",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockUseCase.EXPECT().Find(gomock.Any()).DoAndReturn(func(acc *entity.AccountEntity) error {
			*acc = *expectedAccount
			return nil
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s", accountID.String()), nil)
		router.ServeHTTP(w, req)

		var responseBody sharedDto.APIResponse[dto.AccountResponse]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		data := responseBody.Data

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, http.StatusOK, responseBody.Status)
		assert.Equal(t, accountID, data.ID)
		assert.Equal(t, expectedAccount.Name, data.Name)
		assert.Equal(t, expectedAccount.Email, data.Email)
		assert.Equal(t, expectedAccount.Avatar, data.Avatar)
		assert.WithinDuration(t, expectedAccount.CreatedAt, data.CreatedAt, time.Second)
		assert.WithinDuration(t, expectedAccount.UpdatedAt, data.UpdatedAt, time.Second)
	})

	t.Run("should return status 400 for invalid account id", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/accounts/invalid-uuid", nil)
		router.ServeHTTP(w, req)

		var responseBody sharedDto.APIResponse[any]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
	})

	t.Run("should return status 404 when account is not found", func(t *testing.T) {
		accountID := uuid.New()

		mockUseCase.EXPECT().Find(gomock.Any()).Return(sql.ErrNoRows)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s", accountID.String()), nil)
		router.ServeHTTP(w, req)

		var responseBody sharedDto.APIResponse[any]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, responseBody.Status)
	})
}

func TestAccountHandler_FindByEmail(t *testing.T) {
	router, mock := setup(t)

	t.Run("should return status 200 and account by email", func(t *testing.T) {

		accountID := uuid.New()
		now := time.Now()

		expectedAccount := &entity.AccountEntity{
			ID:        accountID,
			Name:      "Gandalf",
			Email:     "gandalf@lor.com.br",
			Password:  "123mago",
			Avatar:    "url",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mock.EXPECT().FindByEmail(gomock.Any()).DoAndReturn(func(acc *entity.AccountEntity) error {
			*acc = *expectedAccount
			return nil
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/find_by_email/%s", expectedAccount.Email), nil)
		router.ServeHTTP(w, req)

		var responseBody sharedDto.APIResponse[dto.AccountResponse]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		data := responseBody.Data

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, http.StatusOK, responseBody.Status)
		assert.Equal(t, expectedAccount.ID, data.ID)
		assert.Equal(t, expectedAccount.Avatar, data.Avatar)
		assert.Equal(t, expectedAccount.Email, data.Email)
		assert.Equal(t, expectedAccount.Name, data.Name)
		assert.WithinDuration(t, expectedAccount.CreatedAt, data.CreatedAt, time.Second)
		assert.WithinDuration(t, expectedAccount.UpdatedAt, data.UpdatedAt, time.Second)
	})

	t.Run("should return status 404 when account not found", func(t *testing.T) {
		accountEmail := "gandalf@lor.com.br"

		mock.EXPECT().FindByEmail(gomock.Any()).Return(sql.ErrNoRows)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/find_by_email/%s", accountEmail), nil)
		router.ServeHTTP(w, req)

		var responseBody sharedDto.APIResponse[dto.AccountResponse]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, http.StatusNotFound, responseBody.Status)
	})

	t.Run("should return status 500 when internal server error", func(t *testing.T) {
		accountEmail := "gandalf@lor.com.br"

		mock.EXPECT().FindByEmail(gomock.Any()).Return(errors.New("db connection error"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/find_by_email/%s", accountEmail), nil)
		router.ServeHTTP(w, req)

		var responseBody sharedDto.APIResponse[dto.AccountResponse]
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, http.StatusInternalServerError, responseBody.Status)
	})
}
