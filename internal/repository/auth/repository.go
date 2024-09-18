package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/POMBNK/gateway/internal/entity"
	"github.com/POMBNK/gateway/pkg/client/postgres"
)

type Repo struct {
	conn postgres.Client
}

func NewRepo(conn postgres.Client) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) CreateUser(ctx context.Context, user entity.UserToSignUp) (int, error) {
	insertBuilder := sq.Insert("users").PlaceholderFormat(sq.Dollar).
		Columns("last_name", "first_name", "email", "phone", "username", "password").
		Values(user.LastName, user.FirstName, user.Email, user.Phone, user.Username, user.Password).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int
	err = r.conn.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *Repo) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	selectBuilder := sq.Select("id", "last_name", "first_name", "email", "phone", "username", "password").
		PlaceholderFormat(sq.Dollar).
		From("users").
		Where(sq.Eq{"username": username})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return entity.User{}, err
	}

	var user entity.User
	err = r.conn.QueryRow(ctx, query, args...).Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &user.Phone, &user.Username, &user.PasswordHash)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *Repo) FindUserBydID(ctx context.Context, userID int) (entity.User, error) {

	selectBuilder := sq.Select("last_name", "first_name", "email", "phone", "username").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return entity.User{}, err
	}

	var user entity.User
	err = r.conn.QueryRow(ctx, query, args...).Scan(&user.LastName, &user.FirstName, &user.Email, &user.Phone, &user.Username)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *Repo) UpdateUser(ctx context.Context, bannerID int, user entity.User) error {

	updateBuilder := sq.Update("users").PlaceholderFormat(sq.Dollar).
		Set("last_name", user.LastName).
		Set("first_name", user.FirstName).
		Set("email", user.Email).
		Set("phone", user.Phone).
		Set("username", user.Username).
		Where(sq.Eq{"id": bannerID})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repo) DeleteUser(ctx context.Context, userID int) error {

	deleteBuilder := sq.Delete("users").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
