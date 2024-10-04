package shared

import "github.com/Nukie90/my-fluffy/app/domain/entity"

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObserver(userID string, notificationType string)
}

type Observer interface {
	Update(string, string) error
}

type NotificationFactory interface {
	CreateNotification(adminUsername, username, notificationType string) (*entity.Notification, error)
}
