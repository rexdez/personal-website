package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/rexdez/personal-website/controllers"
)

func InitRouter(handler controllers.Controller) *chi.Mux{
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", handler.Homepage)
	return r
}