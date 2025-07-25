package main

import (
	"fmt"
	"link-cutter/internal/auth"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running")
	server.ListenAndServe()
}
