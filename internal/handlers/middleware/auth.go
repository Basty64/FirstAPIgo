package middleware

import (
	"context"
	"github.com/gofrs/uuid"
)

type authKey struct{}

func UserIDFromContext(ctx context.Context) uuid.UUID {
	userID, _ := ctx.Value(authKey{}).(uuid.UUID)
	return userID
}
