package middleware

import (
	"context"
	"doduykhang/hermes-gateway/pkg/service"
	"log"
	"net/http"
)

type Authenticate struct {
	tokenService service.Token		
}

func NewAuthenticate(tokenService service.Token) *Authenticate {
	return &Authenticate{
		tokenService: tokenService,
	}
}


func (a *Authenticate) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session-id")
		if err != nil {
			log.Printf("cookie not valid, %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)	
			w.Write([]byte(err.Error()))
			return
		}

		userID, err := a.tokenService.CheckToken(cookie.Value)
		if err != nil {
			log.Printf("token not valid, %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)	
			w.Write([]byte(err.Error()))
			return
		}

    		ctx := context.WithValue(r.Context(), "userID", userID)

    		next.ServeHTTP(w, r.WithContext(ctx))
  	})
}
