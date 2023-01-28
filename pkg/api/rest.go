package api

import (
	"doduykhang/hermes-gateway/internal/proto"
	"doduykhang/hermes-gateway/pkg/config"
	"doduykhang/hermes-gateway/pkg/controller"
	"doduykhang/hermes-gateway/pkg/middleware"
	"doduykhang/hermes-gateway/pkg/route"
	"doduykhang/hermes-gateway/pkg/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRestApi() {
	app := fiber.New()

 	api := app.Group("/api", logger.New()) // /api
	
	//config
	grpcConn := config.NewGrpcConnection()
	redisClient := config.NewRedis()
	
	//service 
	authService := proto.NewAccountServiceClient(grpcConn)
	cacheService := service.NewCache(redisClient)
	tokenService := service.NewToken(cacheService)
	
	//middleware 
	authMiddleware := middleware.NewAuthenticate(tokenService)

	//controller 
	authController := controller.NewAuth(authService, tokenService)

	//route
	route.AuthRoute(api, authController)
	route.TAuth(api, authController, authMiddleware)

	log.Fatal(app.Listen(":8080"))
}
