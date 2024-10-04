package model

type CreatePost struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Picture []byte  `json:"picture"`
	OwnerID string  `json:"owner_id"`
	Reward  float64 `json:"reward"`
}

type Post struct {
	ID      uint    `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Status  string  `json:"status"`
	Picture []byte  `json:"picture"`
	OwnerID string  `json:"owner_id"`
	FoundID string  `json:"found_id"`
	Reward  float64 `json:"reward"`
}

type FoundPost struct {
	ID      uint   `json:"id"`
	FoundID string `json:"found_id"`
}
