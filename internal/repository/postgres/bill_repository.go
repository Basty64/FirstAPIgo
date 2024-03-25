package postgres

import (
	"MyFirstAPIgo/internal/domain"
	"context"

	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BillRepository struct {
	pool *pgxpool.Pool
}

func NewBillRepository(pool *pgxpool.Pool) *BillRepository {
	return &BillRepository{pool: pool}
}


func (b *BillRepository) Save(ctx context.Context, bill *domain.Bill) error {
	_, err := b.pool.Exec(ctx, "INSERT INTO bank.bill (id, name) VALUES ($1, $2)", bill.ID())

	if err != nil {
		return fmt.Errorf("insert bill: %w", err)
	}

	return nil
}
