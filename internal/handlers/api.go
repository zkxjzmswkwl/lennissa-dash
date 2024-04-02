package handlers

import (
	"main/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Auth)
		router.Get("/", GetRunescapeAccount)
		router.Get("/name/{name}", GetRunescapeAccountByName)
		router.Post("/", CreateRunescapeAccount)

		router.Route("/items", func(itemsRouter chi.Router) {
			itemsRouter.Post("/", GotItem)
			itemsRouter.Get("/name/{name}", GetItems)
		})
	})
}
