package controller

import (
	"context"
	"doduykhang/hermes-gateway/internal/proto"
	"doduykhang/hermes-gateway/pkg/service"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (c *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var req proto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		w.Write([]byte(err.Error()))
		return
	}

	res, err := c.service.Register(context.Background(), &req)
	if err != nil {
		st, _:= status.FromError(err)
		if st.Code() == codes.Unavailable {
			w.WriteHeader(http.StatusConflict)	
			w.Write([]byte("Email already used"))
			return
		}	
	}

	token, err := c.tokenService.CreateToken(res.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		w.Write([]byte(err.Error()))
		return
	}
		
	http.SetCookie(w, &http.Cookie{
		Name: "session-id",
		Value: token,
		Path: "/",
		HttpOnly: true,
	})

	w.Write([]byte("Ok"))
}

func (c *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var req proto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		w.Write([]byte(err.Error()))
		return
	}


	res, err := c.service.Login(context.Background(), &req)
	if err != nil {
		st, _:= status.FromError(err)
		if st.Code() == codes.Unauthenticated {
			w.WriteHeader(http.StatusUnauthorized)	
			w.Write([]byte("Wrong user name or password"))
			return
		}	
	}

	token, err := c.tokenService.CreateToken(res.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		w.Write([]byte(err.Error()))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session-id",
		Value: token,
		Path: "/api",
		HttpOnly: true,
	})

	w.Write([]byte("Ok"))
}
