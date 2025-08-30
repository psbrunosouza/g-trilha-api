package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	FindAccount(ctx context.Context, arg uuid.UUID) (FindAccountRow, error)
}

var _ Querier = (*Queries)(nil)
