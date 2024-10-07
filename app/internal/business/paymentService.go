package business

import (
	"fmt"

	"github.com/Nukie90/my-fluffy/app/domain/model"
)

type PaymentService struct {
}


func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) ProcessPaymentService(payment *model.PayPalPayment) error {

	fmt.Println("Processing payment for UserID:", payment.UserID)
	fmt.Println("Amount:", payment.Amount)
	fmt.Println("Transaction successful")

	return nil
}
