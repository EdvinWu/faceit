package service

import (
	"context"
	"faceit-test/internal/domain/user/model"
	"faceit-test/internal/domain/user/publisher"
	"faceit-test/internal/domain/user/repository"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

var (
	ErrUserCreate    = errors.New("Failed to create user")
	ErrUserUpdate    = errors.New("Failed to update user")
	ErrUserDelete    = errors.New("Failed to delete user")
	ErrUserPaginated = errors.New("Failed to retrieve paginated users")
)

type User interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, userID string) error
	PaginatedUsers(ctx context.Context, user model.PagedUserRequest) ([]*model.User, int, error)
}

type service struct {
	repo      repository.User
	publisher publisher.User
}

func NewUser(repo repository.User, publisher publisher.User) User {
	return &service{repo: repo, publisher: publisher}
}

func (s *service) CreateUser(ctx context.Context, user model.User) (string, error) {
	userID, err := s.repo.CreateUser(ctx, mapUserModelToEntity(user))
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create user"))
		return "", ErrUserCreate
	}
	go s.publishUser(ctx, user, model.UserCreated)
	return userID, nil
}

func (s *service) UpdateUser(ctx context.Context, user model.User) error {
	err := s.repo.UpdateUser(ctx, mapUserModelToEntity(user))
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to update user"))
		return ErrUserUpdate
	}
	go s.publishUser(ctx, user, model.UserUpdated)
	return nil
}

func (s *service) DeleteUser(ctx context.Context, userID string) error {
	err := s.repo.DeleteUser(ctx, userID)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to delete user"))
		return ErrUserDelete
	}
	go s.publishUser(ctx, model.User{ID: userID}, model.UserDeleted)
	return err
}

func (s *service) PaginatedUsers(ctx context.Context, paginatedRequest model.PagedUserRequest) ([]*model.User, int, error) {
	pagedRequest := mapPagedUserRequestToEntity(paginatedRequest)
	users, err := s.repo.PaginatedUsers(ctx, pagedRequest)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to get paginated users"))
		return nil, 0, ErrUserPaginated
	}
	count, err := s.repo.CountPaginatedUsers(ctx, pagedRequest)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to count paginated users"))
		return nil, 0, ErrUserPaginated
	}
	return mapUserEntitiesToModels(users), count, nil
}

func (s *service) publishUser(ctx context.Context, user model.User, action model.UserModificationAction) {
	err := s.publisher.Notify(ctx, user, action)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to publish user"))
	}
}
