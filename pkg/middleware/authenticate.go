package middleware

import (
	"doduykhang/hermes-gateway/pkg/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Authenticate struct {
	tokenService service.Token		
}

func NewAuthenticate(tokenService service.Token) *Authenticate {
	return &Authenticate{
		tokenService: tokenService,
	}
}

func (a *Authenticate) Authenticate(c *fiber.Ctx) error {
	sessionId := c.Cookies("session-id", "")
	unAuthCode := http.StatusUnauthorized
	if sessionId == "" {
		return c.Status(unAuthCode).Send([]byte("Who are you ???"))
	}
	token, err := a.tokenService.CheckToken(sessionId)
	if err != nil {
		return c.Status(unAuthCode).Send([]byte("Invalid token"))
	}

	c.Locals("userID", token)
	return c.Next()
}
