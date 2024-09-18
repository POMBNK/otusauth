package auth

import (
	"context"
	"github.com/POMBNK/gateway/internal/entity"
)

type AuthUserRepo interface {
	AuthRepo
	FindUserBydID(ctx context.Context, userID int) (entity.User, error)
	UpdateUser(ctx context.Context, bannerID int, user entity.User) error
	DeleteUser(ctx context.Context, userID int) error
}

type AuthRepo interface {
	CreateUser(ctx context.Context, user entity.UserToSignUp) (int, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
}
