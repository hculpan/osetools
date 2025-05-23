package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/osetools/internal/dbutils"
)

func KillCharacterHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "characterId")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		slog.Error(fmt.Errorf("invalid id %q: %w", id, err).Error())
	}

	dbCharacter, err := dbutils.GetQueries().GetCharacterById(r.Context(), int32(idNum))
	if err != nil {
		slog.Error("error retrieving character", "msg", err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	dbCharacter.Dead = true
	err = dbutils.UpdateCharacter(r.Context(), dbCharacter)
	if err != nil {
		slog.Error("error updating character", "msg", err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/characters/"+strconv.Itoa(int(dbCharacter.CampaignID)), http.StatusSeeOther)
}
