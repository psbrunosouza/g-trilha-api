//go:build wireinject
// +build wireinject

package wire

import (
	"trilha-api/internal/account/handler"
	"trilha-api/internal/account/repository"
	usecase "trilha-api/internal/account/use_case"
	db "trilha-api/internal/shared/database/sqlc"

	w "github.com/google/wire"
)

var set_account_repository_dependency = w.NewSet(
	repository.New,
	w.Bind(new(repository.AccountRepositoryInterface), new(*repository.AccountRepository)),
)

func NewAccountHandler(db *db.Queries) *handler.AccountHandler {
	w.Build(set_account_repository_dependency, usecase.New, handler.New)
	return &handler.AccountHandler{}
}
