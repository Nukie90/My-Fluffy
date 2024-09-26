package presentation

import (
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/gofiber/fiber/v2"

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
//		@Summary		Create a new user
//		@Description	Create a new user
//		@Tags			users
//		@Accept			json
//		@Param			user	body	model.Signup	true	"Signup information"
//		@Produce		json
//		@Success		200	{string}	string	"User created successfully"
//	 @Failure		400	{string}	string	"Bad request"
//		@Router			/users [post]
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

// GetAllUser godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	[]model.User
//	@Router			/users/all [get]
func (uh *UserHandler) GetAllUser(c *fiber.Ctx) error {
	users, err := uh.UserUsecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
