package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/POMBNK/gateway/internal/entity"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo AuthUserRepo
}

func NewService(repo AuthUserRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) SingUp(ctx context.Context, user entity.UserToSignUp) (int, error) {

	existedUser, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("failed get user by username due error:%w", err)
		}
	}
	if existedUser.Username == user.Username {
		return 0, fmt.Errorf("user with this email has already exist:%w", err)
	}

	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("failled due hashing password error:%w", err)
	}

	user.Password = passwordHash
	uuid, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return uuid, err
	}

	return uuid, nil
}

func (s *Service) SignIn(ctx context.Context, user entity.UserToSignIn) (entity.User, error) {
	existedUser, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed get user by username due error:%w", err)
	}

	if err := validatePassword(user.Password, existedUser.PasswordHash); err != nil {
		return entity.User{}, fmt.Errorf("incorrect password:%w", err)
	}

	return existedUser, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("can't generate hash for given password")
	}
	return string(hash), nil
}

func validatePassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("incorrect password")
	}
	return nil
}
