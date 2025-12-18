package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	proxy *httputil.ReverseProxy
}

func NewReverseProxy(target string) *ReverseProxy {
	url, _ := url.Parse(target)
	return &ReverseProxy{
		proxy: httputil.NewSingleHostReverseProxy(url),
	}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.proxy.ServeHTTP(w, r)
}
