package route

import (
	"doduykhang/hermes-gateway/pkg/controller"
	"doduykhang/hermes-gateway/pkg/middleware"

	"github.com/go-chi/chi/v5"
)

func ConversationRoute(r chi.Router, conversationController *controller.Conversation, authMiddleware *middleware.Authenticate) {
	r.With(authMiddleware.Authenticate).Route("/", func (r  chi.Router)  {
		r.HandleFunc("/*", conversationController.HandleConversationProxy)
	})
}
