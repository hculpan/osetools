package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/osetools/cmd/web/templates"
	"github.com/hculpan/osetools/internal/auth"
)

func AddCharacterFormHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "campaignId")

	comp := templates.AddCharacterForm(appTitle, auth.GetCurrentUsername(r), id, true)
	comp.Render(r.Context(), w)
}
