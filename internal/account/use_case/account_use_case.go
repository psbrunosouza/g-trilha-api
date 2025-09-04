package usecase

import (
	"trilha-api/internal/account/entity"
	"trilha-api/internal/account/repository"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=account_use_case.go -destination=../mocks/account_use_case_mock.go -package=mocks
type AccountUseCaseInterface interface {
	Register(account *entity.AccountEntity) error
	Find(account *entity.AccountEntity) error
	FindByEmail(account *entity.AccountEntity) error
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

func (uc *AccountUseCase) Find(account *entity.AccountEntity) error {
	return uc.repo.Find(account)
}

func (uc *AccountUseCase) FindByEmail(account *entity.AccountEntity) error {
	return uc.repo.FindByEmail(account)
}
