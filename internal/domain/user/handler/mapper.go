package handler

import (
	"faceit-test/internal/domain/user/model"
	"net/url"
	"strconv"
)

func mapToPagedUsersRequest(queryParams url.Values) model.PagedUserRequest {
	return model.PagedUserRequest{
		Nick:    queryParams.Get("nick"),
		Country: queryParams.Get("country"),
		Email:   queryParams.Get("email"),
		Page:    mapPagingParamToInt(queryParams.Get("page")),
		Size:    mapPagingParamToInt(queryParams.Get("size")),
	}
}

func mapPagingParamToInt(param string) int {
	if param == "" {
		return 0
	}
	numParam, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return numParam
}

func mapUpdateToUserModel(user userModel, userID string) model.User {
	return model.User{
		ID:        userID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		NickName:  user.NickName,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func mapToUserModel(user userModel) model.User {
	return model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		NickName:  user.NickName,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func mapUserToUserModel(user *model.User) *userModel {
	return &userModel{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		NickName:  user.NickName,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func mapToPagedUsersResponse(users []*model.User, total int) paginatedResponse {
	return paginatedResponse{
		Users: mapUsersToUserModels(users),
		Page:  page{Total: total, Count: len(users)},
	}
}

func mapUsersToUserModels(users []*model.User) []*userModel {
	res := make([]*userModel, 0, len(users))
	for _, user := range users {
		res = append(res, mapUserToUserModel(user))
	}
	return res
}

func mapErrorToErrorResponse(err error) errorResponse {
	return errorResponse{Error: err.Error()}
}
