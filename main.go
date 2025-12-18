package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/yashsharma.js/nimbus/proxy"
)

func main() {
	p := proxy.NewReverseProxy("https://httpbin.org")

	log.Println("Proxy running on :8080")
	http.ListenAndServe(":8080", p)
}