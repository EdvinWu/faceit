package service

import (
	"database/sql"
	"faceit-test/internal/domain/user/entity"
	"faceit-test/internal/domain/user/model"
)

func mapUserEntitiesToModels(users []*entity.User) []*model.User {
	res := make([]*model.User, 0, len(users))
	for _, user := range users {
		res = append(res, mapUserEntityToModel(user))
	}
	return res
}

func mapUserEntityToModel(user *entity.User) *model.User {
	return &model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		NickName:  user.NickName,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: mapNullTimeToInt(user.UpdatedAt),
		DeletedAt: mapNullTimeToInt(user.DeletedAt),
	}
}

func mapPagedUserRequestToEntity(request model.PagedUserRequest) *entity.PagedUserRequest {
	limit, skip := mapPaging(request.Page, request.Size)
	return &entity.PagedUserRequest{
		Nick:    request.Nick,
		Country: request.Country,
		Email:   request.Email,
		Skip:    skip,
		Limit:   limit,
	}
}

func mapPaging(page, size int) (limit, offset int) {
	if size < 1 {
		limit = 10
	} else {
		limit = size
	}
	if page < 2 {
		return limit, 0
	}
	return limit, size * (page - 1)
}

func mapUserModelToEntity(user model.User) *entity.User {
	return &entity.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		NickName:  user.NickName,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	}
}

func mapNullTimeToInt(time sql.NullTime) int64 {
	if time.Valid {
		return time.Time.Unix()
	}
	return 0
}
