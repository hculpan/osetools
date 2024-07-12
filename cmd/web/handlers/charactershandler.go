package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/osetools/cmd/web/templates"
	"github.com/hculpan/osetools/internal/auth"
	"github.com/hculpan/osetools/internal/dbutils"
)

func CharactersHandler(w http.ResponseWriter, r *http.Request) {
	authorized := auth.IsAuthorized(r)

	id := chi.URLParam(r, "campaignId")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		slog.Error(fmt.Errorf("invalid id %q: %w", id, err).Error())
		http.Error(w, fmt.Errorf("invalid id %q: %w", id, err).Error(), http.StatusInternalServerError)
		return
	}

	msg := r.URL.Query().Get("msg")

	characters, err := dbutils.GetQueries().GetCharactersForCampaignWithXp(r.Context(), int64(idNum))
	if err != nil {
		slog.Error(fmt.Errorf("unable to retrieve characters for campaign id %d: %w", idNum, err).Error())
		http.Error(w, fmt.Errorf("unable to retrieve characters for campaign id %d: %w", idNum, err).Error(), http.StatusInternalServerError)
		return
	}

	comp := templates.XpTemplate(appTitle, auth.GetCurrentUsername(r), characters, authorized, msg)
	comp.Render(r.Context(), w)
}
