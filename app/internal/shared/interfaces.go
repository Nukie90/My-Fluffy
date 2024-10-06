package shared

import "github.com/Nukie90/my-fluffy/app/domain/model"

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObserver(userID string, t string)
}

type Observer interface {
	Update(string, string) error
}

type PaymentGateway interface {
    ProcessPayment(payment *model.CreatePayment) error
}
