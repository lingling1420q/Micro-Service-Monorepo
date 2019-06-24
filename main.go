package main

import (
	"demo/routes"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        routes.InitRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
