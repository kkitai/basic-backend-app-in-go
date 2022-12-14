package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/kkitai/basic-backend-app-in-go/docs"
	"github.com/kkitai/basic-backend-app-in-go/repository"

	httpSwagger "github.com/swaggo/http-swagger"
)

var telephoneRepository repository.TelephoneRepository

// NewHandler returns http.Handler.
// this function configure a route path and middleware using go-chi/chi.
func NewHandler(tr repository.TelephoneRepository) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// api document
	r.Get("/swagger/*", httpSwagger.Handler())

	r.Route("/telephones", func(r chi.Router) {
		r.Get("/", listTelephone)
		r.Route("/{number}", func(r chi.Router) {
			r.Get("/", getTelephone)
			//r.Put("/", putTelephone)
			r.Post("/", postTelephone)
		})
	})

	telephoneRepository = tr

	return r
}
