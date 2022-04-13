package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string       `db:"id"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	NickName  string       `db:"nickname"`
	Password  string       `db:"password"`
	Email     string       `db:"email"`
	Country   string       `db:"country"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type PagedUserRequest struct {
	Nick    string
	Country string
	Email   string
	Skip    int
	Limit   int
}
