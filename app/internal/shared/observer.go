package shared

import "fmt"

type UserNotifier struct {
	Observers []Observer
}

var _ Subject = (*UserNotifier)(nil)

func (ucn *UserNotifier) Register(observer Observer) {
	ucn.Observers = append(ucn.Observers, observer)
}

func (ucn *UserNotifier) Unregister(observer Observer) {
	for i, o := range ucn.Observers {
		if o == observer {
			ucn.Observers = append(ucn.Observers[:i], ucn.Observers[i+1:]...)
			break
		}
	}
}

func (ucn *UserNotifier) NotifyObserver(receiver string, sender string, notificationType string) {
	fmt.Println("Notifying observers")
	for _, o := range ucn.Observers {
		err := o.Update(receiver, sender, notificationType)
		if err != nil {
			fmt.Println(err)
		}
	}
}
