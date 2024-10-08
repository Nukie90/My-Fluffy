package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RouterProxy struct {
	realRouter  *Router
	maxAllowed  int
	rateLimiter map[string]int
	resetTime   time.Duration
}

func NewRouterProxy(realRouter *Router) *RouterProxy {
	proxy := &RouterProxy{
		realRouter:  realRouter,
		maxAllowed:  50,
		rateLimiter: make(map[string]int),
		resetTime:   30 * time.Second,
	}
	go proxy.resetRateLimiter()
	return proxy
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

// Reset the rateLimiter map every resetTime duration
func (p *RouterProxy) resetRateLimiter() {
	for {
		//if you want to reset the rate limiter every 30 seconds only if the rate limiter is not empty
		if len(p.rateLimiter) > 0 {
			time.Sleep(p.resetTime)
			fmt.Println("Proxy: Resetting rate limiter.")
			p.rateLimiter = make(map[string]int)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
