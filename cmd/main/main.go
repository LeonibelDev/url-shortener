package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/leonibeldev/url-shortener/internal/db"
	"github.com/leonibeldev/url-shortener/internal/routes"
)

var tmpl = template.Must(template.ParseFiles("internal/front/index.html"))

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logMessage := fmt.Sprintf("[%s] - %s %s - %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
		fmt.Println(logMessage)
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	if strings.TrimPrefix(r.URL.Path, "/") == "" {

		tmpl.Execute(w, nil)
	} else {
		routes.HandleRedirect(w, r)
	}
}

func main() {
	// Validate DB
	db.ValidateDBExists()

	http.HandleFunc("/shorten", routes.HandleShorten)

	// Logger middleware
	http.Handle("/", Logger(http.HandlerFunc(Home)))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Server running on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
