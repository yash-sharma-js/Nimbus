package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/yashsharma.js/nimbus/cache"
)

type ReverseProxy struct {
	proxy *httputil.ReverseProxy
	cache *cache.Cache
}

func NewReverseProxy(target string) *ReverseProxy {
	u, _ := url.Parse(target)

	return &ReverseProxy{
		proxy: httputil.NewSingleHostReverseProxy(u),
		cache: cache.New(),
	}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Only cache GET requests
	if r.Method == http.MethodGet {
		key := r.Method + ":" + r.URL.String()

		if item, ok := rp.cache.Get(key); ok {
			for k, v := range item.Headers {
				w.Header().Set(k, v)
			}
			w.Write(item.Response)
			return
		}

		rec := newResponseRecorder(w)
		rp.proxy.ServeHTTP(rec, r)

		rp.cache.Set(key, cache.Item{
			Response: rec.body.Bytes(),
			Headers:  rec.headers,
			Expiry:   time.Now().Add(10 * time.Second),
		})
		return
	}

	rp.proxy.ServeHTTP(w, r)
}
