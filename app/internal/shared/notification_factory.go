package shared

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/oklog/ulid/v2"
)

type DefaultNotificationFactory struct{}

func (f *DefaultNotificationFactory) CreateNotification(adminUsername, username, notificationType string) (*entity.Notification, error) {
	var message string

	switch notificationType {
	case "createUser":
		message = fmt.Sprintf("%s: User: %s is created", adminUsername, username)
	default:
		return nil, fmt.Errorf("unknown notification type: %s", notificationType)
	}

	adminULID, err := ulid.Parse(adminUsername)
	if err != nil {
		return nil, err
	}

	return &entity.Notification{
		UserID:  adminULID,
		Message: message,
	}, nil
}
