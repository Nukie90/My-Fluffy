package presentation

import (
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/Nukie90/my-fluffy/app/internal/business"
)

// UserHandler is the handler for users
type UserHandler struct {
	UserUsecase *business.UserUsecase
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(uu *business.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: uu}
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Param			user	body	model.Signup	true	"Signup information"
//	@Produce		json
//	@Success		200	{string}	string	"User created successfully"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/users [post]
func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var signup model.Signup
	if err := c.BodyParser(&signup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := uh.UserUsecase.Create(&signup)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			users
//	@Accept			json
//	@Param			user	body	model.Login	true	"Login information"
//	@Produce		json
//	@Success		200	{string}	string	"Login successful"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/users/login [post]
func (uh *UserHandler) Login(c *fiber.Ctx) error {
	var login model.Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user model.User

	user, err := uh.UserUsecase.Login(&login)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	//set to session
	c.Cookie(&fiber.Cookie{
		Name:  "session",
		Value: user.ID,
	})

	return c.JSON(fiber.Map{"message": "Login successful"})
}

// GetAllUser godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	[]model.User	"List of users"
//	@Router			/users/all [get]
func (uh *UserHandler) GetAllUser(c *fiber.Ctx) error {
	users, err := uh.UserUsecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetNotifications godoc
//
//	@Summary		Get all notifications
//	@Description	Get all notifications for current user
//	@Tags			notifications
//	@Produce		json
//	@Success		200	{object}	[]model.Notification	"List of notifications"
//	@Router			/notifications [get]
func (uh *UserHandler) GetNotifications(c *fiber.Ctx) error {
	userID := c.Cookies("session")
	notifications, err := uh.UserUsecase.GetNotifications(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(notifications)
}

// DeleteNotification godoc
//
//	@Summary		Delete a notification
//	@Description	Delete a notification
//	@Tags			notifications
//	@Param			id	path	string	true	"Notification ID"
//	@Produce		json
//	@Success		200	{string}	string	"Notification deleted"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/notifications/{id} [delete]
func (uh *UserHandler) DeleteNotification(c *fiber.Ctx) error {
	notificationID := c.Params("id")
	notificationIDUint, err := strconv.ParseUint(notificationID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = uh.UserUsecase.DeleteNotification(uint(notificationIDUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Notification deleted"})
}
