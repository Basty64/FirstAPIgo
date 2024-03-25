package postgres

import (
	"MyFirstAPIgo/internal/domain"
	"MyFirstAPIgo/internal/pkg"
	"context"
	"fmt"
)

type RepositoryUser struct {
	connection pkg.Connection
}

func NewUserRepository(connection pkg.Connection) *RepositoryUser {
	return &RepositoryUser{
		connection: connection,
	}
}

func (r *RepositoryUser) Save(ctx context.Context, user *domain.User) error {
	query, args, err := psql.
		Insert(`auth.app_user`).
		Columns("id", "name", "password").
		Values(user.Id(), user.Name(), user.Password()).
		ToSql()
	if err != nil {
		return fmt.Errorf("build insert user query: %w", err)
	}

	_, err = r.connection.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

func (r *RepositoryUser) FindByName(ctx context.Context, name string) ([]*domain.User, error) {

	//to do

	return nil, nil
}

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByName(ctx context.Context, name string) (*domain.User, error)
}
