package model

import "github.com/oklog/ulid/v2"

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

type PaginatedPostWithUsername struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Status   string    `json:"status"`
	Picture  []byte    `json:"picture"`
	OwnerID  ulid.ULID `json:"owner_id"`
	FoundID  ulid.ULID `json:"found_id"`
	Reward   float64   `json:"reward"`
}

type PaginatedPostResponse struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	Status   string  `json:"status"`
	Picture  []byte  `json:"picture"`
	OwnerID  string  `json:"owner_id"`
	FoundID  string  `json:"found_id"`
	Reward   float64 `json:"reward"`
}

type ConfirmationPost struct {
	ID uint `json:"id"`
}
