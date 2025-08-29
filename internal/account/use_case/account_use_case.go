package usecase

import (
	"trilha-api/internal/account/entity"
	"trilha-api/internal/account/repository"

	"golang.org/x/crypto/bcrypt"
)

type AccountUseCaseInterface interface {
	Register(account *entity.AccountEntity) error
}

type AccountUseCase struct {
	repo repository.AccountRepositoryInterface
}

func New(repo repository.AccountRepositoryInterface) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}

func (uc *AccountUseCase) Register(account *entity.AccountEntity) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	account.Password = string(hashedPassword)

	return uc.repo.Register(account)
}
