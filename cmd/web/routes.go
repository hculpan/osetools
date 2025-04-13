package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/osetools/cmd/web/handlers"
	"github.com/hculpan/osetools/internal/auth"
)

func routes(router *chi.Mux) *chi.Mux {
	if router != nil {
		router.Get("/", handlers.CampaignsHandler)
		router.Get("/info", handlers.InfoHandler)
		router.Get("/login", handlers.LoginHandler)
		router.Post("/login", handlers.LoginPostHandler)
		router.Get("/logout", handlers.LogoutHandler)

		router.Get("/home", auth.AuthMiddleware(handlers.CampaignsHandler))
		router.Get("/account", auth.AuthMiddleware(handlers.AccountHandler))
		router.Get("/campaigns", auth.AuthMiddleware(handlers.CampaignsHandler))
		router.Get("/characters/{campaignId}", handlers.CharactersHandler)
		router.Get("/character/{characterId}", handlers.CharacterHandler)
		router.Get("/add-xp/{campaignId}", auth.AuthMiddleware(handlers.AddXpHandler))
		router.Post("/add-xp", auth.AuthMiddleware(handlers.AddXpPostHandler))
		router.Get("/add-character/{campaignId}", auth.AuthMiddleware(handlers.AddCharacterFormHandler))
		router.Post("/add-character", auth.AuthMiddleware(handlers.AddCharacterPostHandler))
		router.Get("/deletecharacter/{characterId}", auth.AuthMiddleware(handlers.DeleteCharacterHandler))
		router.Get("/raisecharacter/{characterId}", auth.AuthMiddleware(handlers.RaiseCharacterHandler))
		router.Get("/killcharacter/{characterId}", auth.AuthMiddleware(handlers.KillCharacterHandler))
		router.Get("/add-campaign", auth.AuthMiddleware(handlers.AddCampaignHandler))
		router.Post("/add-campaign", auth.AuthMiddleware(handlers.AddCampaignPostHandler))

		// Serve static files
		fileServer := http.FileServer(http.Dir("./static"))
		router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	}

	return router
}
