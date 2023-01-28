package controller

import (
	"context"
	"doduykhang/hermes-gateway/internal/proto"
	"doduykhang/hermes-gateway/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	service proto.AccountServiceClient
	tokenService service.Token
}

func NewAuth(service proto.AccountServiceClient, tokenService service.Token) *Auth {
	return &Auth{
		service: service,
		tokenService: tokenService,
	}
}

func (c *Auth) Register(ctx *fiber.Ctx) error {
	var req proto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := c.service.Register(context.Background(), &req)
	if err != nil {
		return err
	}

	token, err := c.tokenService.CreateToken(res.UserID)
	if err != nil {
		return err
	}

	var response struct {
		Token string `json:"token"`		
	}

	ctx.Cookie(&fiber.Cookie{
		Name: "session-id",
		Value: token,
		Path: "/api",
		Domain: "localhost",
		Secure: true,
		HTTPOnly: true,
	})

	response.Token = token
	return ctx.JSON(response)
}

func (c *Auth) Login(ctx *fiber.Ctx) error {
	var req proto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := c.service.Login(context.Background(), &req)
	if err != nil {
		return err
	}

	token, err := c.tokenService.CreateToken(res.UserID)
	if err != nil {
		return err
	}

	var response struct {
		Token string `json:"token"`		
	}

	ctx.Cookie(&fiber.Cookie{
		Name: "session-id",
		Value: token,
		Path: "/api",
		Domain: "localhost",
		Secure: true,
		HTTPOnly: true,
	})
	
	response.Token = token
	return ctx.JSON(response)
}
