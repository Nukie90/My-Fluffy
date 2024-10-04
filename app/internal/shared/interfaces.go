package shared

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObserver(userID string, t string)
}

type Observer interface {
	Update(string, string) error
}
