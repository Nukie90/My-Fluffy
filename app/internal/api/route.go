package api

import (
	"github.com/Nukie90/my-fluffy/app/internal/presentation"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	userHandler    *presentation.UserHandler
	postHandler    *presentation.PostHandler
	paymentHandler *presentation.PaymentHandler
}

func NewRouter(uh *presentation.UserHandler, ph *presentation.PostHandler, pmh *presentation.PaymentHandler) *Router {
	return &Router{
		userHandler: uh,
		postHandler: ph,
		paymentHandler: pmh,
	}
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
				users.Post("/login", r.userHandler.Login)
			}
			post := v1.Group("/posts")
			{
				post.Post("/", r.postHandler.CreatePost)
				post.Get("/user", r.postHandler.GetPostsFromSpecificUser)
				post.Put("/found", r.postHandler.FoundPet)
			}
			payment := v1.Group("/payments")
			{
				payment.Post("/", r.paymentHandler.CreateUserPayment)
				payment.Get("/user", r.paymentHandler.GetPaymentsFromSpecificUser)
			}
		}
	}
}
