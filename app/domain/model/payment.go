package model

type CreatePayment struct {
	Amount      float64 `json:"amount"`
	ReceiverID  string  `json:"receiver_id"`
	UserID      string  `json:"user_id"`
}
type Payment struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	ReceiverID  string  `json:"receiver_id"`
	UserID      string  `json:"user_id"`
}

type PayPalPayment struct {
    Amount      float64 
    ReceiverID  string	
    UserID      string
}
