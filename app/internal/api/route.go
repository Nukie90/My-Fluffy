package api

import (
	"github.com/Nukie90/my-fluffy/app/internal/presentation"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	userHandler *presentation.UserHandler
}

func NewRouter(uh *presentation.UserHandler) *Router {
	return &Router{userHandler: uh}
}

func (r *Router) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.Post("/", r.userHandler.CreateUser)
				users.Get("/all", r.userHandler.GetAllUser)
			}
		}
	}
}
