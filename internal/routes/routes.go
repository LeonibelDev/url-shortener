package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/leonibeldev/url-shortener/internal/actions"
	"github.com/leonibeldev/url-shortener/internal/models"
)

var urls []models.URL

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("url")
	// handle shorten

	id, _ := actions.GenerateID()
	urls = append(urls, models.URL{
		ID:   id,
		Link: r.URL.Query().Get("url"),
	})

	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", id)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")
	fmt.Println("ID: " + id)

	for _, url := range urls {
		if url.ID == id {
			http.Redirect(w, r, url.Link, http.StatusSeeOther)
			return
		}
	}

}
