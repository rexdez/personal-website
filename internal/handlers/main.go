package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rexdez/personal-website/internal/handlers/website"
	"github.com/rexdez/personal-website/pkg/common"
)

type Config struct {
	Pool *pgxpool.Pool
	Mux *chi.Mux
	Logger *log.Logger
}

func (cfg *Config) StaticHandler(root string) http.Handler{
	fs := http.FileServer(http.Dir(root))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		splits := strings.Split(root, ".")
		extension := splits[len(splits)-1]
		w.Header().Set("Content-Type", common.MimeTypes[extension])
		fs.ServeHTTP(w, r)
	})
}

func (cfg *Config) InitRouter() {
	
	pages := website.Webpages{Pool:cfg.Pool, Logger:cfg.Logger, Mux:cfg.Mux}
	// Initializing templates
	pages.LoadTemplates()

	// Initializing router
	r := chi.NewRouter()
	fs := cfg.StaticHandler("web/static")
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", pages.Homepage)
	cfg.Mux = r
}