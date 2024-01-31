package domain

import (
	"MyFirstAPIgo/internal/pkg"
	"context"
	"github.com/gofrs/uuid"
)

type Bill struct {
	id       uuid.UUID
	name     string
	balance  int
	isClosed bool
	UserID   uuid.UUID
	Cards    []uuid.UUID
}

func (b *Bill) ID() uuid.UUID  { return b.id }
func (b *Bill) Name() string   { return b.name }
func (b *Bill) Balance() int   { return b.balance }
func (b *Bill) IsClosed() bool { return b.isClosed }

func (b *Bill) Validate() error {
	if b.name == "" {
		return &pkg.EmptyFieldError{Field: "name"}
	}
	return nil
}

func NewBill(name string, userID uuid.UUID) *Bill {
	return &Bill{
		id:       uuid.Must(uuid.NewV7()),
		name:     name,
		UserID:   userID,
		balance:  0,
		isClosed: false,
	}
}

func (b *Bill) Close() {
	b.isClosed = true
}

type BillRepository interface {
	Save(ctx context.Context, bill *Bill) error
}
