package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shan251197/bookings/pkg/config"
	"github.com/shan251197/bookings/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/generals-quarters", handler.Repo.Generals)
	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Get("/search-availability", handler.Repo.Availability)
	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/make-reservation", handler.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
