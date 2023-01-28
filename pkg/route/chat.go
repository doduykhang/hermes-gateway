package route

import (
	"doduykhang/hermes-gateway/pkg/controller"
	"doduykhang/hermes-gateway/pkg/middleware"

	"github.com/go-chi/chi/v5"
)

func ChatRoute(r chi.Router, chatController *controller.Chat, authMiddleware *middleware.Authenticate) {
	r.With(authMiddleware.Authenticate).Route("/chat", func (r chi.Router) {
		r.HandleFunc("/*", chatController.HandleChatProxy)
	})
}
