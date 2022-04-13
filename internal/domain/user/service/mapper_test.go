package service

import (
	"database/sql"
	"faceit-test/internal/domain/user/entity"
	"faceit-test/internal/domain/user/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_mapNullTimeToInt(t *testing.T) {
	tests := []struct {
		name string
		time sql.NullTime
		want int64
	}{
		{name: "empty", time: sql.NullTime{}, want: 0},
		{name: "valid zero time", time: sql.NullTime{Time: time.Unix(0, 0)}, want: 0},
		{name: "valid not zero time", time: sql.NullTime{Valid: true, Time: time.Unix(123, 0)}, want: 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mapNullTimeToInt(tt.time))
		})
	}
}

func Test_mapUserModelToEntity(t *testing.T) {
	input := model.User{
		ID:        "id",
		FirstName: "firstName",
		LastName:  "lastName",
		NickName:  "nickname",
		Password:  "pwd",
		Email:     "email",
		Country:   "country",
		CreatedAt: 1,
		UpdatedAt: 2,
		DeletedAt: 3}
	expected := &entity.User{
		ID:        "id",
		FirstName: "firstName",
		LastName:  "lastName",
		NickName:  "nickname",
		Password:  "pwd",
		Email:     "email",
		Country:   "country",
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
		DeletedAt: sql.NullTime{},
	}
	got := mapUserModelToEntity(input)
	assert.Equal(t, expected, got)
}

func Test_mapPaging(t *testing.T) {
	type args struct {
		page int
		size int
	}
	tests := []struct {
		name       string
		args       args
		wantLimit  int
		wantOffset int
	}{
		{name: "empty", args: args{}, wantLimit: 10, wantOffset: 0},
		{name: "size > 1, page < 2", args: args{size: 3}, wantLimit: 3, wantOffset: 0},
		{name: "size > 1, page > 2", args: args{size: 3, page: 3}, wantLimit: 3, wantOffset: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimit, gotOffset := mapPaging(tt.args.page, tt.args.size)
			assert.Equal(t, tt.wantLimit, gotLimit)
			assert.Equal(t, tt.wantOffset, gotOffset)
		})
	}
}

func Test_mapPagedUserRequestToEntity(t *testing.T) {
	input := model.PagedUserRequest{
		Nick:    "nick",
		Country: "country",
		Email:   "email",
		Page:    1,
		Size:    3,
	}
	expected := &entity.PagedUserRequest{
		Nick:    "nick",
		Country: "country",
		Email:   "email",
		Limit:   3,
		Skip:    0,
	}
	got := mapPagedUserRequestToEntity(input)
	assert.Equal(t, expected, got)
}

func Test_mapUserEntityToModel(t *testing.T) {
	input := &entity.User{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: time.Unix(123, 0),
		UpdatedAt: sql.NullTime{},
		DeletedAt: sql.NullTime{},
	}
	exptected := &model.User{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: 123,
		UpdatedAt: 0,
		DeletedAt: 0,
	}
	got := mapUserEntityToModel(input)
	assert.Equal(t, exptected, got)
}

func Test_mapUserEntitiesToModels(t *testing.T) {
	input := []*entity.User{{ID: "id"}, {ID: "id2"}}
	expected := []*model.User{{ID: "id", CreatedAt: time.Time{}.Unix()}, {ID: "id2", CreatedAt: time.Time{}.Unix()}}
	got := mapUserEntitiesToModels(input)
	assert.ElementsMatch(t, expected, got)
}
