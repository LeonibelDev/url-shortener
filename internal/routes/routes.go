package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/leonibeldev/url-shortener/internal/actions"
	"github.com/leonibeldev/url-shortener/internal/db"
	"github.com/leonibeldev/url-shortener/internal/models"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	// Get User URL
	var userUrl models.URL
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &userUrl)

	// Generate ID
	id, _ := actions.GenerateID()

	// Data
	data := models.URL{
		ID:   id,
		Link: userUrl.Link,
	}

	// add to db
	db.AddNewUrl(data.ID, data.Link)

	// format data to json
	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", jsonData)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")

	// Get URL
	link := db.GetUrl(id)
	http.Redirect(w, r, link, http.StatusSeeOther)
}
