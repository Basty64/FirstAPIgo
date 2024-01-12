package usecase

import (
	"MyFirstAPIgo/internal/domain"
	"MyFirstAPIgo/internal/repository/postgres"
	"context"
)

type FindUserUseCase struct {
	user postgres.RepositoryUser
}

type FindUserQuery struct {
	Name string
}

type UserSearchParameters struct {
	Name string
}

func (u *FindUserUseCase) Handle(ctx context.Context, q FindUserQuery) ([]*domain.User, error) {
	parameters := UserSearchParameters{
		Name: q.Name,
	}

	return u.user.FindByName(ctx, parameters.Name)
}
