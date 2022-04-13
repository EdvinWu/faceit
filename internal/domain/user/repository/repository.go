package repository

import (
	"context"
	"faceit-test/internal/domain/user/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User interface {
	CreateUser(ctx context.Context, user *entity.User) (string, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID string) error
	PaginatedUsers(ctx context.Context, user *entity.PagedUserRequest) ([]*entity.User, error)
	CountPaginatedUsers(ctx context.Context, user *entity.PagedUserRequest) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) User {
	return &repository{db: db}
}

func (r *repository) CreateUser(_ context.Context, user *entity.User) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 9)
	if err != nil {
		return "", errors.Wrap(err, "failed to encrypt users password")
	}
	userID := uuid.New().String()
	userMap := mapCreateUserToSQLMap(user, userID, string(password), time.Now())
	_, err = r.db.NamedExec(queryCreateUser, userMap)
	return userID, err
}

func (r *repository) UpdateUser(_ context.Context, user *entity.User) error {
	_, err := r.db.NamedExec(queryUpdateUser, mapUpdateUserToSQLMap(user, time.Now()))
	return err

}

func (r *repository) DeleteUser(_ context.Context, userID string) error {
	_, err := r.db.Exec(queryDeleteUser, userID, time.Now().UTC())
	return err
}

func (r *repository) PaginatedUsers(_ context.Context, user *entity.PagedUserRequest) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.Select(
		&users, queryPagedUsers, "%"+user.Nick+"%", "%"+user.Country+"%", "%"+user.Email+"%", user.Limit, user.Skip)
	return users, err
}

func (r *repository) CountPaginatedUsers(_ context.Context, user *entity.PagedUserRequest) (int, error) {
	var count int
	err := r.db.Get(&count, queryCountPagedUsers, "%"+user.Nick+"%", "%"+user.Country+"%", "%"+user.Email+"%")
	return count, err
}

const (
	queryCreateUser = `
	INSERT INTO "user" (id, email, password, first_name, last_name, nickname, country, created_at)
	VALUES(:id, :email, :password, :first_name, :last_name, :nickname, :country, :created_at)`

	queryUpdateUser = `
	UPDATE "user" SET first_name = :first_name, last_name = :last_name, nickname = :nickname, updated_at = :updated_at
	WHERE "user".id = :id`

	queryDeleteUser = `UPDATE "user" SET deleted_at = $2 where "user".id = $1`

	queryPagedUsers = `
	SELECT id, email, password, first_name, last_name, nickname, country, created_at, updated_at, deleted_at FROM "user"
	WHERE nickname ILIKE $1 AND country ILIKE $2 AND email ILIKE $3 limit $4 offset $5`

	queryCountPagedUsers = `
	SELECT COUNT(*) FROM "user"
	WHERE nickname ILIKE $1 AND country ILIKE $2 AND email ILIKE $3
`
)
