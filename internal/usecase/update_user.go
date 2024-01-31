package usecase

import (
	"MyFirstAPIgo/internal/domain"
	"MyFirstAPIgo/internal/pkg"
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

type UserRepository interface {
	FindByIDForUpdate(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}

type UpdateUserUseCase struct {
	users              UserRepository
	transactionManager pkg.TransactionManager
}

type UpdateUserCommand struct {
	ID   uuid.UUID
	Name string
}

func (useCAse *UpdateUserUseCase) Handle(ctx context.Context, command UpdateUserCommand) error {
	return useCAse.transactionManager.Do(ctx, func(ctx context.Context) error {
		user, err := useCAse.users.FindByIDForUpdate(ctx, command.ID)
		if err != nil {
			return err
		}

		if user == nil {
			return errors.New("not found")
		}

		user.Update(command.Name)

		if err := useCAse.users.Update(ctx, user); err != nil {
			return err
		}

		return nil
	})
}
