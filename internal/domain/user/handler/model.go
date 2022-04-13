package handler

type userModel struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Country   string `json:"country,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	DeletedAt int64  `json:"deleted_at,omitempty"`
}

type page struct {
	Total int
	Count int
}

type paginatedResponse struct {
	Users []*userModel `json:"data"`
	Page  page         `json:"Page"`
}

type errorResponse struct {
	Error string `json:"error"`
}
