package business

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/entity"
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
func (n *AdminNotifier) Update(receiver, sender, notificationType string) error {
	admins, err := n.userRepo.FindAdmin()
	if err != nil {
		fmt.Println("1: ", err)
		return err
	}

	for _, admin := range admins {
		notification, err := n.notificationFactory.CreateNotification(admin.Username, sender, notificationType)
		if err != nil {
			fmt.Println("2: ", err)
			return err
		}

		err = n.userRepo.StoreNotification(&entity.Notification{
			UserID:  admin.ID,
			Message: notification,
		})
		if err != nil {
			fmt.Println("3: ", err)
		}
	}

	return nil
}
