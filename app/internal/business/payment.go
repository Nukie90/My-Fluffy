package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/Nukie90/my-fluffy/app/internal/shared"

	"github.com/oklog/ulid/v2"
)

type PaymentUsecase struct {
	PaymentRepo    *repository.PaymentRepo
	PaymentGateway shared.PaymentGateway
	ClientNotifier *shared.UserNotifier
}

func NewPaymentUsecase(pr *repository.PaymentRepo, pg shared.PaymentGateway, cn *shared.UserNotifier) *PaymentUsecase {
	return &PaymentUsecase{PaymentRepo: pr, PaymentGateway: pg, ClientNotifier: cn}
}

func (pu *PaymentUsecase) CreateUserPayment(payment *model.CreatePayment) error {
	err := pu.PaymentGateway.ProcessPayment(payment)
	if err != nil {
		return err
	}

	payInfo := &entity.Payment{
		Amount:     payment.Amount,
		ReceiverID: ulid.MustParse(payment.ReceiverID),
		UserID:     ulid.MustParse(payment.UserID),
	}

	err = pu.PaymentRepo.CreatePayment(payInfo)
	if err != nil {
		return err
	}

	pu.ClientNotifier.NotifyObserver(payment.ReceiverID, payment.UserID, "payment")

	return nil
}

func (pu *PaymentUsecase) GetPaymentsFromSpecificUser(userID string) ([]entity.Payment, error) {

	//to ulid
	useridUlid, err := ulid.Parse(userID)
	if err != nil {
		return nil, err
	}

	payments, err := pu.PaymentRepo.GetPaymentFromSpecificUser(useridUlid)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
