package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/Nukie90/my-fluffy/app/internal/presentation"
	"github.com/Nukie90/my-fluffy/app/internal/repository"

	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Home
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	uh := presentation.NewUserHandler(business.NewUserUsecase(repository.NewUserRepo(db)))
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.Post("/", uh.CreateUser)
				users.Get("/all", uh.GetAllUser)
			}
		}
	}
}
