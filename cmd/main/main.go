package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/leonibeldev/url-shortener/internal/db"
	"github.com/leonibeldev/url-shortener/internal/routes"
)

func main() {
	// Validate DB
	db.ValidateDBExists()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/") == "" {
			tmpl := template.Must(template.ParseFiles("internal/front/index.html"))

			tmpl.Execute(w, nil)
		} else {
			routes.HandleRedirect(w, r)
		}
	})

	http.HandleFunc("/shorten", routes.HandleShorten)
	//http.HandleFunc("/r/", routes.HandleRedirect)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Server running on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
