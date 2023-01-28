package route

import (
	"doduykhang/hermes-gateway/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, authController *controller.Auth) {
	user := r.Group("/auth")
	user.Post("/register", authController.Register)
	user.Post("/login", authController.Login)
}
