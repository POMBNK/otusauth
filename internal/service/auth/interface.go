package auth

import (
	"context"
	"github.com/POMBNK/gateway/internal/entity"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, user entity.UserToSignUp) (int, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
}
