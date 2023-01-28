package config

import (
	"net/http/httputil"
	"net/url"
)

func NewProxy(rawUrl string) *httputil.ReverseProxy {
	url, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	return httputil.NewSingleHostReverseProxy(url)
}
