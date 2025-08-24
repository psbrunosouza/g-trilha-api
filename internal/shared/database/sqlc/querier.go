package db

import (
	"context"
)

// Querier is an interface that wraps the Queries struct
// to allow for mocking.
type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
