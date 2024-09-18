package auth

import (
	"context"
	"github.com/POMBNK/gateway/internal/entity"
)

func (s *Service) FindUserBydID(ctx context.Context, userID int) (entity.User, error) {
	return s.repo.FindUserBydID(ctx, userID)
}

func (s *Service) UpdateUser(ctx context.Context, bannerID int, user entity.User) error {
	return s.repo.UpdateUser(ctx, bannerID, user)
}

func (s *Service) DeleteUser(ctx context.Context, bannerID int) error {
	return s.repo.DeleteUser(ctx, bannerID)
}
