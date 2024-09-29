package model

type Signup struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password"`
	Role     string `json:"role" example:"admin"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
