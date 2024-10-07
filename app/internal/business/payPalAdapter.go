package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/model"
)

type PayPalAdapter struct {
	paymentService PaymentService
}
func NewPayPalAdapter(service PaymentService) *PayPalAdapter {
	return &PayPalAdapter{
		paymentService: service,
	}
}

func (p *PayPalAdapter) ProcessPayment(payment *model.CreatePayment) error {
	payPalPayment := &model.PayPalPayment{
		Amount:      payment.Amount,
		ReceiverID:  payment.ReceiverID,
		UserID:      payment.UserID,
	}

	err := p.paymentService.ProcessPaymentService(payPalPayment)
	if err != nil {
		return err
	}

	return nil
}


