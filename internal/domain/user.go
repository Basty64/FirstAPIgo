package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type User struct {
	id       uuid.UUID `json:"id"`
	name     string    `json:"name"`
	email    string    `json:"email"`
	password []byte    `json:"password"`
}

func (u *User) SetId(id uuid.UUID) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password []byte) {
	u.password = password
}

func (u *User) Id() uuid.UUID    { return u.id }
func (u *User) Name() string     { return u.name }
func (u *User) Email() string    { return u.email }
func (u *User) Password() []byte { return u.password }

func NewUser(name string, email string, passwordHash []byte) *User {
	return &User{
		id:       uuid.Must(uuid.NewV7()),
		name:     name,
		email:    email,
		password: passwordHash,
	}
}

func (u *User) Update(name string) {
	u.name = name
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name string) (*User, error)
	// FindByParameters(ctx context.Context, parameters UserSearchParameters) ([]*User, error)
}
