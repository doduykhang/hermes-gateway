package route

import (
	"doduykhang/hermes-gateway/pkg/controller"
	"doduykhang/hermes-gateway/pkg/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TAuth(r chi.Router, authController *controller.Auth, authMidddleware *middleware.Authenticate) {
	r.Route("/test", func (r chi.Router) {
		r.With(authMidddleware.Authenticate).Get("/", func (w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := ctx.Value("userID").(string)
			w.Write([]byte(userID))
		})
	})
}
