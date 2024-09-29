package shared

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObserver(userID string)
}

type Observer interface {
	Update(string) error
}
