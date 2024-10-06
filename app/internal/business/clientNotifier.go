package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/Nukie90/my-fluffy/app/internal/shared"
)

type ClientNotifier struct {
	userRepo            *repository.UserRepo
	notificationFactory shared.NotificationFactory
}

// Ensure ClientNotifier implements the Observer interface.
var _ shared.Observer = (*ClientNotifier)(nil)

func NewClientNotifier(ur *repository.UserRepo, nf shared.NotificationFactory) *ClientNotifier {
	return &ClientNotifier{
		userRepo:            ur,
		notificationFactory: nf,
	}
}

// Update is the method that is called when a new user is created.
func (n *ClientNotifier) Update(receiver, sender, notificationType string) error {
	receiverEntity, err := n.userRepo.FindByID(receiver)
	if err != nil {
		return err
	}

	senderEntity, err := n.userRepo.FindByID(sender)
	if err != nil {
		return err
	}

	notification, err := n.notificationFactory.CreateNotification(receiverEntity.Username, senderEntity.Username, notificationType)
	if err != nil {
		return err
	}

	err = n.userRepo.StoreNotification(&entity.Notification{
		UserID:  receiverEntity.ID,
		Message: notification,
	})
	if err != nil {
		return err
	}

	return nil
}
