package pkg


import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)


type TransactionManager interface {
	Do(ctx context.Context, f func(ctx context.Context) error) error
}


type PoolConnection struct {
	pool *pgxpool.Pool
}


type MockTransactionManager struct {
}

func (m MockTransactionManager) Do(ctx context.Context, f func(ctx context.Context) error) error {
	return f(ctx)
}
