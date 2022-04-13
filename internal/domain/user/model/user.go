package model

type UserModificationAction string

const (
	UserCreated UserModificationAction = "created"
	UserUpdated UserModificationAction = "updated"
	UserDeleted UserModificationAction = "deleted"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	NickName  string
	Password  string
	Email     string
	Country   string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

type PagedUserRequest struct {
	Nick    string
	Country string
	Email   string
	Page    int
	Size    int
}
