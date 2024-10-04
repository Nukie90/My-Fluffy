package business

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/Nukie90/my-fluffy/app/internal/shared"
)

type AdminNotifier struct {
	userRepo            *repository.UserRepo
	notificationFactory shared.NotificationFactory
}

// Ensure AdminNotifier implements the Observer interface.
var _ shared.Observer = (*AdminNotifier)(nil)

func NewAdminNotifier(ur *repository.UserRepo, nf shared.NotificationFactory) *AdminNotifier {
	return &AdminNotifier{
		userRepo:            ur,
		notificationFactory: nf,
	}
}

// Update is the method that is called when a new user is created.
func (n *AdminNotifier) Update(username, notificationType string) error {
	admins, err := n.userRepo.FindAdmin()
	if err != nil {
		return err
	}

	for _, admin := range admins {
		notification, err := n.notificationFactory.CreateNotification(admin.Username, username, notificationType)
		if err != nil {
			return err
		}

		err = n.userRepo.StoreNotification(notification)
		if err != nil {
			return err
		}

		fmt.Println(notification.Message)
	}

	return nil
}
