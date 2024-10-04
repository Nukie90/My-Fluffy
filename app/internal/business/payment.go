package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/oklog/ulid/v2"
)

type PaymentUsecase struct {
	PaymentRepo *repository.PaymentRepo
}

func NewPaymentUsecase(pr *repository.PaymentRepo) *PaymentUsecase {
	return &PaymentUsecase{PaymentRepo: pr}
}

func (pu *PaymentUsecase) CreateUserPayment(payment *model.CreatePayment) error {
	payInfo := &entity.Payment{
		Amount:      payment.Amount,
		Transaction: payment.Transaction,
		UserID:      ulid.MustParse(payment.UserID),
	}

	err := pu.PaymentRepo.CreatePayment(payInfo)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PaymentUsecase) GetPaymentsFromSpecificUser(userID ulid.ULID) ([]entity.Payment, error) {
	payments, err := pu.PaymentRepo.GetPaymentFromSpecificUser(userID)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
