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

func GoldForXpHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "campaignId")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		slog.Error(fmt.Errorf("invalid id %q: %w", id, err).Error())
		http.Error(w, fmt.Errorf("invalid id %q: %w", id, err).Error(), http.StatusInternalServerError)
		return
	}

	characters, err := dbutils.GetQueries().GetCharactersForCampaign(r.Context(), int64(idNum))
	if err != nil {
		slog.Error(fmt.Errorf("unable to retrieve characters for campaign id %d: %w", idNum, err).Error())
		http.Error(w, fmt.Errorf("unable to retrieve characters for campaign id %d: %w", idNum, err).Error(), http.StatusInternalServerError)
		return
	}

	comp := templates.GoldForXpTemplate(appTitle, auth.GetCurrentUsername(r), characters, 1)
	comp.Render(r.Context(), w)
}
