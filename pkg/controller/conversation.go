package controller

import (
	"net/http"
	"net/http/httputil"
)

type Conversation struct {
	proxy *httputil.ReverseProxy
}

func NewConversation(proxy *httputil.ReverseProxy) *Conversation {
	return &Conversation{
		proxy: proxy,
	}
}

func (c *Conversation) HandleConversationProxy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)
	r.Header.Del("x-user-id")
	r.Header.Add("x-user-id", userID)
	c.proxy.ServeHTTP(w, r)	
}
