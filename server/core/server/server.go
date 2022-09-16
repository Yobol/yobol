package server

import (
	"net/http"
)

func initServer(address string, router http.Handler) server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
