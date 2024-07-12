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

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	authorized := auth.IsAuthorized(r)

	id := chi.URLParam(r, "characterId")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		slog.Error(fmt.Errorf("invalid id %q: %w", id, err).Error())
	}

	character, err := dbutils.GetQueries().GetCharacterWithXp(r.Context(), int64(idNum))
	if err != nil {
		slog.Error(fmt.Errorf("unable to retrieve characters for campaign id %d: %w", idNum, err).Error())
	}

	xpBonusReasons, err := dbutils.GetQueries().GetXpBonusForCharacter(r.Context(), character.ID)
	if err != nil {
		slog.Error(fmt.Errorf("unable to retrieve xp bonus reasons for character id %d: %w", character.ID, err).Error())
	}

	xpAwards, err := dbutils.GetQueries().GetXpAwardsForCharacter(r.Context(), character.ID)
	if err != nil {
		slog.Error(fmt.Errorf("unable to retrieve xp awards for character id %d: %w", character.ID, err).Error())
	}

	comp := templates.CharacterTemplate(appTitle, auth.GetCurrentUsername(r), authorized, character, xpBonusReasons, xpAwards)
	comp.Render(r.Context(), w)
}
