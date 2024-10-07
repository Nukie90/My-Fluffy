package shared

import "github.com/Nukie90/my-fluffy/app/domain/model"

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObserver(userID string, notificationType string)
}

type Observer interface {
	Update(string, string, string) error
}

type NotificationFactory interface {
	CreateNotification(adminUsername, username, notificationType string) (string, error)
}

type PaymentGateway interface {
    ProcessPayment(payment *model.CreatePayment) error
}

type PaymentService interface {

    ProcessPaymentService(payment *model.PayPalPayment) error

}

