package model

type CreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Picture []byte `json:"picture"`
	OwnerID string `json:"owner_id"`
}

type Post struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Picture []byte `json:"picture"`
	OwnerID string `json:"owner_id"`
	FoundID string `json:"found_id"`
}
