package shared

import (
	"fmt"
)

type DefaultNotificationFactory struct{}

func (f *DefaultNotificationFactory) CreateNotification(receiver, sender, notificationType string) (string, error) {
	var message string

	switch notificationType {
	case "createUser":
		fmt.Println("Creating user notification")
		message = fmt.Sprintf("%s: User: %s is created", receiver, sender)
	case "createPost":
		fmt.Println("Creating post notification")
		message = fmt.Sprintf("%s has created a post", sender)
	case "foundPet":
		fmt.Println("Creating found pet notification")
		message = fmt.Sprintf("%s: %s has found your pet!", receiver, sender)
	case "payment":
		fmt.Println("Creating payment notification")
		message = fmt.Sprintf("%s: %s has made a payment", receiver, sender)
	case "Confirmation":
		fmt.Println("Creating confirmation notification")
		message = fmt.Sprintf("%s: %s has gifted you a reward! Check your banking account! แต้งกิ้วหลาย", receiver, sender)
	default:
		return "", fmt.Errorf("notification type not found")
	}

	fmt.Println(message)

	return message, nil
}
