package business

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/Nukie90/my-fluffy/app/internal/shared"
)

type AdminNotifier struct {
	userRepo *repository.UserRepo
}

// Ensure AdminNotifier implements the Observer interface.
var _ shared.Observer = (*AdminNotifier)(nil)

func NewAdminNotifier(ur *repository.UserRepo) *AdminNotifier {
	return &AdminNotifier{
		userRepo: ur,
	}
}

// Update is the method that is called when a new user is created.
func (n *AdminNotifier) Update(username string) error {
	admins, err := n.userRepo.FindAdmin()
	if err != nil {
		return err
	}

	fmt.Println("Notifying admins about user creation")

	for _, admin := range admins {
		notificationMessage := fmt.Sprintf("%s: User %s is created", admin.Username, username)
		notification := &entity.Notification{
			UserID:  admin.ID,
			Message: notificationMessage,
		}

		err := n.userRepo.StoreNotification(notification)
		if err != nil {
			return err
		}

		fmt.Println(notificationMessage)
	}

	return nil

}
