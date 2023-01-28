package route

import (
	"doduykhang/hermes-gateway/pkg/controller"
	"doduykhang/hermes-gateway/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func TAuth(r fiber.Router, authController *controller.Auth, authMidddleware *middleware.Authenticate) {
	user := r.Group("/t")
	user.Get("/", authMidddleware.Authenticate, func (c *fiber.Ctx) error {
		return c.Send([]byte("You are in"))	
	})
}
