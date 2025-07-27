package main

import (
	"fmt"
	"link-cutter/configs"
	"link-cutter/internal/auth"
	"link-cutter/internal/link"
	"link-cutter/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	database := db.NewDB(config)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(database)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running")
	server.ListenAndServe()
}
