package model

type Signup struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password"`
	Role     string `json:"role" example:""`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Login struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password"`
}

type Notification struct {
	ID       uint   `json:"id"`
	OwnerID  string `json:"owner_id"`
	Message  string `json:"message"`
	CreateAt string `json:"create_at"`
}
