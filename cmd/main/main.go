package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/leonibeldev/url-shortener/internal/routes"
)

func main() {

	http.HandleFunc("/shorten", routes.HandleShorten)
	http.HandleFunc("/", routes.HandleRedirect)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Server running on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
