package shared

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
