package repository

import (
	"faceit-test/internal/domain/user/entity"
	"time"
)

func mapCreateUserToSQLMap(user *entity.User, id, password string, created time.Time) map[string]interface{} {
	return map[string]interface{}{
		"id":         id,
		"email":      user.Email,
		"password":   password,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"nickname":   user.NickName,
		"country":    user.Country,
		"created_at": created.UTC(),
	}
}

func mapUpdateUserToSQLMap(user *entity.User, updated time.Time) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"nickname":   user.NickName,
		"updated_at": updated.UTC(),
	}
}
