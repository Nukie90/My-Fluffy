package api

import "github.com/gofiber/fiber/v2"

type RouterSubject interface {
	SetupRoutes(app *fiber.App)
}
