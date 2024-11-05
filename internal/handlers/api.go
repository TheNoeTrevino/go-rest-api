package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
  "rest-api-project/internal/middleware"
)

func Handler(r *chi.Mux) {
  // makes it to where a / at the end is not an error
	r.Use(chimiddle.StripSlashes)

  r.Route("/account", func(router chi.Router) {

    router.Use(middleware.Authorization)

    router.Get("/coins", GetCoinBalance)

  })
}
