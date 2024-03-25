package inmemory

import (
	"MyFirstAPIgo/internal/domain"
	"context"
	"github.com/gofrs/uuid"
)

type BillRepository struct {
	bills map[uuid.UUID]*domain.Bill
}

func NewBillRepository() *BillRepository {
	return &BillRepository{}
}

func (b *BillRepository) Save(ctx context.Context, bill *domain.Bill) error {
	//TODO implement me
	panic("implement me")
}
