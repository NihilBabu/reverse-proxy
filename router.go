package main

import (
	"proxy/handlers"

	"github.com/gorilla/mux"
)

//ProxyRouter is ...
func ProxyRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ReverseProxy)
	return r
}
