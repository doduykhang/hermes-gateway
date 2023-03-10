package api

import (
	"doduykhang/hermes-gateway/internal/proto"
	"doduykhang/hermes-gateway/pkg/config"
	"doduykhang/hermes-gateway/pkg/controller"
	"doduykhang/hermes-gateway/pkg/middleware"
	"doduykhang/hermes-gateway/pkg/route"
	"doduykhang/hermes-gateway/pkg/service"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMid "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)


func NewRestApi() {
	r := chi.NewRouter()
	r.Use(chiMid.Logger)
		
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000", "http://localhost:3000"},
    		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    		AllowCredentials: true,
    		MaxAge:           300, // Maximum value not ignored by any of major browsers
  	}))

	
	//config
	conf := config.LoadConfig()
	grpcConn := config.NewGrpcConnection(conf.GRPC.Account)
	redisClient := config.NewRedis(conf)
	chatProxy := config.NewProxy(conf.Proxy.Chat)
	conversationProxy := config.NewProxy(conf.Proxy.Conversation)
	
	//service 
	authService := proto.NewAccountServiceClient(grpcConn)
	cacheService := service.NewCache(redisClient)
	tokenService := service.NewToken(cacheService)
	
	//middleware 
	authMiddleware := middleware.NewAuthenticate(tokenService)

	//controller 
	authController := controller.NewAuth(authService, tokenService)
	chatController := controller.NewChat(chatProxy)
	conversationController := controller.NewConversation(conversationProxy)

	//route
	r.Route("/api", func (r chi.Router) {
		route.AuthRoute(r, authController)
		route.TAuth(r, authController, authMiddleware)
		route.ChatRoute(r, chatController, authMiddleware)
		route.ConversationRoute(r, conversationController, authMiddleware)
	})

	log.Printf("Gateway starting at port %s", conf.Server.Port)	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Server.Port), r))
}
