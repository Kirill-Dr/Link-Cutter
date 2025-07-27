package main

import (
	"fmt"
	"link-cutter/configs"
	"link-cutter/internal/auth"
	"link-cutter/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	_ = db.NewDB(config)
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
