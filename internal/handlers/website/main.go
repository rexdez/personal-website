package website

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Homepage *template.Template
	ContactPage *template.Template
	BlogViewPge *template.Template
)

type Webpages struct {
	Pool *pgxpool.Pool
	Logger *log.Logger
	Mux *chi.Mux
}

// LoadTemplates initializes the templates declared as global variables
func (conn *Webpages) LoadTemplates() {
	htmlRootPath := "web/html/"

	Homepage = template.Must(template.ParseFiles(fmt.Sprintf("%shome.html", htmlRootPath)))
	ContactPage = template.Must(template.ParseFiles(fmt.Sprintf("%scontact.html", htmlRootPath)))
	BlogViewPge = template.Must(template.ParseFiles(fmt.Sprintf("%sblogview.html", htmlRootPath)))
}

// Homepage serves the homepage at "/"
func (conn *Webpages) Homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/html")
	err := Homepage.Execute(w, nil)
	if err != nil {
		go conn.Logger.Println(err)
	}
}
