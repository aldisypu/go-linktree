package route

import (
	"go-linktree/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	LinkController *http.LinkController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
	c.App.Get("/api/users/:username/links", c.LinkController.List)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)

	c.App.Post("/api/links", c.LinkController.Create)
	c.App.Put("/api/links/:linkId", c.LinkController.Update)
	c.App.Get("/api/links/:linkId", c.LinkController.Get)
	c.App.Delete("/api/links/:linkId", c.LinkController.Delete)
}
