package shared

import "fmt"

type UserCreationNotifier struct {
	Observers []Observer
}

func (ucn *UserCreationNotifier) Register(observer Observer) {
	ucn.Observers = append(ucn.Observers, observer)
}

func (ucn *UserCreationNotifier) Unregister(observer Observer) {
	for i, o := range ucn.Observers {
		if o == observer {
			ucn.Observers = append(ucn.Observers[:i], ucn.Observers[i+1:]...)
			break
		}
	}
}

func (ucn *UserCreationNotifier) NotifyObserver(username string, notificationType string) {
	fmt.Println("Notifying observers")
	for _, o := range ucn.Observers {
		err := o.Update(username, notificationType)
		if err != nil {
			return
		}
	}
}
