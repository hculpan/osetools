package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/osetools/cmd/web/handlers"
	"github.com/hculpan/osetools/internal/auth"
)

func routes(router *chi.Mux) *chi.Mux {
	if router != nil {
		router.Get("/", handlers.HomeHandler)
		router.Get("/info", handlers.InfoHandler)
		router.Get("/login", handlers.LoginHandler)
		router.Post("/login", handlers.LoginPostHandler)
		router.Get("/logout", handlers.LogoutHandler)

		router.Get("/home", auth.AuthMiddleware(handlers.HomeHandler))
		router.Get("/account", auth.AuthMiddleware(handlers.AccountHandler))
		router.Get("/campaigns", auth.AuthMiddleware(handlers.CampaignsHandler))
		router.Get("/characters/{campaignId}", handlers.CharactersHandler)
		router.Get("/character/{characterId}", handlers.CharacterHandler)
		router.Get("/add-xp/{campaignId}", auth.AuthMiddleware(handlers.AddXpHandler))
		router.Post("/add-xp", auth.AuthMiddleware(handlers.AddXpPostHandler))
		router.Get("/gold-for-xp/{campaignId}", auth.AuthMiddleware(handlers.GoldForXpHandler))
		router.Get("/add-character/{campaignId}", auth.AuthMiddleware(handlers.AddCharacterFormHandler))
		router.Post("/add-character", auth.AuthMiddleware(handlers.AddCharacterPostHandler))

		// Serve static files
		fileServer := http.FileServer(http.Dir("./static"))
		router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	}

	return router
}
