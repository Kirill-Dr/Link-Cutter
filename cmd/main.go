package main

import (
	"fmt"
	"link-cutter/configs"
	"link-cutter/internal/auth"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running")
	server.ListenAndServe()
}
