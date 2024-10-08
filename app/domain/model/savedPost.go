package model

// SuccessResponse represents a successful response with a message.
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response with a message.
type ErrorResponse struct {
	Error string `json:"error"`
}

type SavedPost struct {
	UserID string `json:"user_id"`
	PostID uint   `json:"post_id"`
}

type UnsavePostRequest struct {
	PostID uint `json:"post_id" validate:"required"`
}

type PostWithUsername struct {
	ID       uint    `json:"id"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	Status   string  `json:"status"`
	Picture  []byte  `json:"picture"`
	Reward   float64 `json:"reward"`
	OwnerID  string  `json:"owner_id"`
	FoundID  string  `json:"found_id"`
	Username string  `json:"username"` // Add username field
}
