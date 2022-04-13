package handler

import (
	"faceit-test/internal/domain/user/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var (
	testUserModel = &model.User{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: 123,
		UpdatedAt: 123,
		DeletedAt: 123,
	}
	testHandlerUserModel = &userModel{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: 123,
		UpdatedAt: 123,
		DeletedAt: 123,
	}
)

func Test_mapErrorToErrorResponse(t *testing.T) {
	message := "error msg"
	response := mapErrorToErrorResponse(errors.New(message))
	assert.Equal(t, errorResponse{Error: message}, response)
}

func Test_mapUsersToUserModels(t *testing.T) {
	input := []*model.User{testUserModel}
	expected := []*userModel{testHandlerUserModel}
	got := mapUsersToUserModels(input)
	assert.Equal(t, expected, got)
}

func Test_mapToPagedUsersResponse(t *testing.T) {
	total := 12
	input := []*model.User{testUserModel}
	expected := paginatedResponse{
		Users: []*userModel{testHandlerUserModel},
		Page:  page{Total: 12, Count: 1}}
	got := mapToPagedUsersResponse(input, total)
	assert.Equal(t, expected, got)
}

func Test_mapUserToUserModel(t *testing.T) {
	got := mapUserToUserModel(testUserModel)
	assert.Equal(t, testHandlerUserModel, got)
}

func Test_mapToUserModel(t *testing.T) {
	got := mapToUserModel(*testHandlerUserModel)
	assert.Equal(t, *testUserModel, got)
}

func Test_mapUpdateToUserModel(t *testing.T) {
	got := mapUpdateToUserModel(*testHandlerUserModel, id)
	expected := model.User{
		ID:        "id1",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: 123,
		UpdatedAt: 123,
		DeletedAt: 123,
	}
	assert.Equal(t, expected, got)
}

func Test_mapPagingParamToInt(t *testing.T) {
	tests := []struct {
		name  string
		param string
		want  int
	}{
		{name: "empty", param: "", want: 0},
		{name: "not a number", param: "string", want: 0},
		{name: "number", param: "12", want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mapPagingParamToInt(tt.param))
		})
	}
}

func Test_mapToPagedUsersRequest(t *testing.T) {
	input := url.Values{}
	nickValue := "nick"
	countryValue := "country"
	emailValue := "email"
	input.Add("nick", nickValue)
	input.Add("country", countryValue)
	input.Add("email", emailValue)
	input.Add("page", "3")
	input.Add("size", "13")
	expected := model.PagedUserRequest{
		Nick:    nickValue,
		Country: countryValue,
		Email:   emailValue,
		Page:    3,
		Size:    13,
	}
	got := mapToPagedUsersRequest(input)
	assert.Equal(t, expected, got)
}
