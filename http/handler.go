package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kkitai/basic-backend-app-in-go/db"
	_ "github.com/kkitai/basic-backend-app-in-go/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

var telephoneRepository db.TelephoneRepository

func NewHandler(tr db.TelephoneRepository) http.Handler {
	r := chi.NewRouter()

	// TODO: pass an arbitrary logger
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
