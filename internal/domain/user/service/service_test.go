package service

import (
	"context"
	"faceit-test/internal/domain/user/entity"
	"faceit-test/internal/domain/user/model"
	mockPublisher "faceit-test/internal/domain/user/publisher/mocks"
	mockRepository "faceit-test/internal/domain/user/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var (
	errDB     = errors.New("db error")
	ctx       = context.Background()
	userCount = 1
	userModel = model.User{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
	}
	userEntity = &entity.User{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
	}
	paginatedRequest = model.PagedUserRequest{
		Nick:    "nick",
		Country: "country",
		Email:   "email",
		Page:    4,
		Size:    10,
	}
	paginatedEntity = &entity.PagedUserRequest{
		Nick:    "nick",
		Country: "country",
		Email:   "email",
		Skip:    30,
		Limit:   10,
	}
	userModels = []*model.User{&userModel}
)

func Test_service_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mockRepository.NewMockUser(ctrl)
	userPublisher := mockPublisher.NewMockUser(ctrl)
	userService := NewUser(userRepository, userPublisher)

	t.Run("failed to create user", func(t *testing.T) {
		userRepository.EXPECT().CreateUser(gomock.Any(), userEntity).Return("", errDB)

		userID, err := userService.CreateUser(ctx, userModel)

		assert.Empty(t, userID)
		assert.Equal(t, ErrUserCreate, err)
	})

	t.Run("success", func(t *testing.T) {
		wg := waitGroup(1)
		userRepository.EXPECT().CreateUser(gomock.Any(), userEntity).Return(userModel.ID, nil)
		userPublisher.EXPECT().Notify(gomock.Any(), userModel, model.UserCreated).Do(publishInterceptor(wg))

		userID, err := userService.CreateUser(ctx, userModel)

		wg.Wait()
		assert.Equal(t, userModel.ID, userID)
		assert.NoError(t, err)
	})
}

func Test_service_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mockRepository.NewMockUser(ctrl)
	userPublisher := mockPublisher.NewMockUser(ctrl)
	userService := NewUser(userRepository, userPublisher)

	t.Run("failed to update user", func(t *testing.T) {
		userRepository.EXPECT().UpdateUser(gomock.Any(), userEntity).Return(errDB)

		err := userService.UpdateUser(ctx, userModel)

		assert.Equal(t, ErrUserUpdate, err)
	})

	t.Run("success", func(t *testing.T) {
		wg := waitGroup(1)
		userRepository.EXPECT().UpdateUser(gomock.Any(), userEntity).Return(nil)
		userPublisher.EXPECT().Notify(gomock.Any(), userModel, model.UserUpdated).Do(publishInterceptor(wg))

		err := userService.UpdateUser(ctx, userModel)

		wg.Wait()
		assert.NoError(t, err)
	})
}

func Test_service_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mockRepository.NewMockUser(ctrl)
	userPublisher := mockPublisher.NewMockUser(ctrl)
	userService := NewUser(userRepository, userPublisher)

	t.Run("failed to update user", func(t *testing.T) {
		userRepository.EXPECT().DeleteUser(gomock.Any(), userEntity.ID).Return(errDB)

		err := userService.DeleteUser(ctx, userModel.ID)

		assert.Equal(t, ErrUserDelete, err)
	})

	t.Run("success", func(t *testing.T) {
		wg := waitGroup(1)
		userRepository.EXPECT().DeleteUser(gomock.Any(), userEntity.ID).Return(nil)
		userPublisher.EXPECT().Notify(gomock.Any(), model.User{ID: userModel.ID}, model.UserDeleted).Do(publishInterceptor(wg))

		err := userService.DeleteUser(ctx, userModel.ID)

		wg.Wait()
		assert.NoError(t, err)
	})
}

func Test_service_PaginatedUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mockRepository.NewMockUser(ctrl)
	userService := NewUser(userRepository, nil)

	userEntities := []*entity.User{{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: time.Unix(0, 0),
	}}

	t.Run("failed to get paginated users", func(t *testing.T) {
		userRepository.EXPECT().PaginatedUsers(gomock.Any(), paginatedEntity).Return(nil, errDB)

		users, count, err := userService.PaginatedUsers(ctx, paginatedRequest)

		assert.Equal(t, ErrUserPaginated, err)
		assert.Nil(t, users)
		assert.Zero(t, count)
	})

	t.Run("failed to count paginated users", func(t *testing.T) {
		userRepository.EXPECT().PaginatedUsers(gomock.Any(), paginatedEntity).Return(userEntities, nil)
		userRepository.EXPECT().CountPaginatedUsers(gomock.Any(), paginatedEntity).Return(0, errDB)

		users, count, err := userService.PaginatedUsers(ctx, paginatedRequest)

		assert.Equal(t, ErrUserPaginated, err)
		assert.Nil(t, users)
		assert.Zero(t, count)
	})

	t.Run("success", func(t *testing.T) {

		userRepository.EXPECT().PaginatedUsers(gomock.Any(), paginatedEntity).Return(userEntities, nil)
		userRepository.EXPECT().CountPaginatedUsers(gomock.Any(), paginatedEntity).Return(userCount, nil)

		users, count, err := userService.PaginatedUsers(ctx, paginatedRequest)

		assert.NoError(t, err)
		assert.Equal(t, userCount, count)
		assert.ElementsMatch(t, userModels, users)
	})
}

func publishInterceptor(wg *sync.WaitGroup) func(ctx context.Context, user model.User, action model.UserModificationAction) {
	return func(ctx context.Context, user model.User, action model.UserModificationAction) {
		wg.Done()
	}
}

func waitGroup(delta int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(delta)
	return wg
}
