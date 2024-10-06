package model

type CreatePayment struct {
	Amount      float64 `json:"amount"`
	Transaction string  `json:"transaction"`
	UserID      string  `json:"user_id"`
}
type Payment struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Transaction string  `json:"transaction"`
	UserID      string  `json:"user_id"`
}

type PayPalPayment struct {
    Amount      float64 
    Transaction string
    UserID      string
}
