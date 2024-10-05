package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type RouterProxy struct {
	realRouter  *Router
	maxAllowed  int
	rateLimiter map[string]int
}

func NewRouterProxy(realRouter *Router) *RouterProxy {
	return &RouterProxy{
		realRouter:  realRouter,
		maxAllowed:  999,
		rateLimiter: make(map[string]int),
	}
}

func (p *RouterProxy) SetupRoutes(app *fiber.App) {
	fmt.Println("Proxy: Setting up routes.")
	app.Use(p.RateLimitingMiddleware)
	p.realRouter.SetupRoutes(app)
}

func (p *RouterProxy) RateLimitingMiddleware(c *fiber.Ctx) error {
	fmt.Println("Proxy: Rate limiting middleware.")
	ip := c.IP()
	p.rateLimiter[ip]++
	if p.rateLimiter[ip] > p.maxAllowed {
		return c.SendStatus(fiber.StatusTooManyRequests)
	}
	return c.Next()
}
