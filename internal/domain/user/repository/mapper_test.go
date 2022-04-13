package repository

import (
	"faceit-test/internal/domain/user/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_mapUpdateUserToSQLMap(t *testing.T) {
	inputUser := &entity.User{
		ID:        "id",
		FirstName: "firstName",
		LastName:  "lastName",
		NickName:  "nickName",
		Password:  "password",
		Email:     "email",
		Country:   "country",
	}
	inputTime := time.Unix(123, 0)
	expected := map[string]interface{}{
		"id":         "id",
		"first_name": "firstName",
		"last_name":  "lastName",
		"nickname":   "nickName",
		"updated_at": time.Unix(123, 0).UTC(),
	}
	got := mapUpdateUserToSQLMap(inputUser, inputTime)
	assert.Equal(t, expected, got)
}

func Test_mapCreateUserToSQLMap(t *testing.T) {
	inputUser := &entity.User{
		ID:        "id",
		FirstName: "firstName",
		LastName:  "lastName",
		NickName:  "nickName",
		Password:  "password",
		Email:     "email",
		Country:   "country",
	}
	id := "id"
	password := "password"
	created := time.Unix(444, 0)
	expected := map[string]interface{}{
		"id":         "id",
		"email":      "email",
		"password":   "password",
		"first_name": "firstName",
		"last_name":  "lastName",
		"nickname":   "nickName",
		"country":    "country",
		"created_at": time.Unix(444, 0).UTC(),
	}
	got := mapCreateUserToSQLMap(inputUser, id, password, created)
	assert.Equal(t, expected, got)
}
