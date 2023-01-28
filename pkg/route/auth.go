package route

import (
	"doduykhang/hermes-gateway/pkg/controller"

	"github.com/go-chi/chi/v5"
)

func AuthRoute(r chi.Router, authController *controller.Auth) {
	r.Route("/auth", func (r chi.Router) {
		r.Post("/register", authController.Register)
		r.Post("/login", authController.Login)
	})
}
