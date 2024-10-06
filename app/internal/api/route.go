package api

import (
	"github.com/Nukie90/my-fluffy/app/internal/presentation"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	userHandler      *presentation.UserHandler
	postHandler      *presentation.PostHandler
	savedPostHandler *presentation.SavedPostHandler
}

func NewRouter(uh *presentation.UserHandler, ph *presentation.PostHandler, sph *presentation.SavedPostHandler) *Router {
	return &Router{
		userHandler:      uh,
		postHandler:      ph,
		savedPostHandler: sph,
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
				users.Get("/:id", r.userHandler.GetUserByID)
				users.Post("/login", r.userHandler.Login)
			}
			post := v1.Group("/posts")
			{
				post.Post("/", r.postHandler.CreatePost)
				post.Get("/user", r.postHandler.GetPostsFromSpecificUser)
				post.Put("/found", r.postHandler.FoundPet)
				post.Get("/feed", r.postHandler.GetPaginatedPosts)
			}
			savedPosts := v1.Group("/saved_posts")
			{
				savedPosts.Post("/", r.savedPostHandler.CreateSavedPost)
				savedPosts.Get("/", r.savedPostHandler.GetAllSavedPostsByUser)
				savedPosts.Delete("/", r.savedPostHandler.UnsavePost)
			}
			noti := v1.Group("/notifications")
			{
				noti.Get("/", r.userHandler.GetNotifications)
				noti.Delete("/:id", r.userHandler.DeleteNotification)
			}
		}
	}
}
