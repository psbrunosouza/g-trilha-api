//go:build wireinject
// +build wireinject

package wire

import (
	"trilha-api/internal/account/handler"
	"trilha-api/internal/account/repository"
	usecase "trilha-api/internal/account/use_case"
	sqlc "trilha-api/internal/shared/database/sqlc"

	w "github.com/google/wire"
)

var set_account_repository_dependency = w.NewSet(
	repository.New,
	w.Bind(new(repository.AccountRepositoryInterface), new(*repository.AccountRepository)),
)

var set_account_usecase_dependency = w.NewSet(
	usecase.New,
	w.Bind(new(usecase.AccountUseCaseInterface), new(*usecase.AccountUseCase)),
)

func NewAccountHandler(db *sqlc.Queries) *handler.AccountHandler {
	w.Build(
		w.Bind(new(sqlc.Querier), new(*sqlc.Queries)),
		set_account_repository_dependency,
		set_account_usecase_dependency,
		handler.New,
	)
	return &handler.AccountHandler{}
}
