package controller

import (
	"net/http"
	"net/http/httputil"
)

type Chat struct {
	proxy *httputil.ReverseProxy
}

func NewChat(proxy *httputil.ReverseProxy) *Chat {
	return &Chat{
		proxy: proxy,
	}
}

func (c *Chat) HandleChatProxy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)
	r.Header.Add("x-user-id", userID)
	c.proxy.ServeHTTP(w, r)	
}
