package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/hculpan/osetools/internal/db"
	"github.com/hculpan/osetools/internal/dbutils"
)

// LoginPostHandler handles the login form submission
func AddXpPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	campaignId := r.FormValue("campaignId")

	charIds := r.Form["selectedIds[]"]
	xpAmount := r.FormValue("xpAmount")
	xpAward, err := strconv.Atoi(xpAmount)
	if err != nil {
		e := fmt.Errorf("error converting xpAmount value of %q : %w", xpAmount, err)
		slog.Error(e.Error())
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	reason := r.FormValue("reason")
	for _, c := range charIds {
		id, err := strconv.Atoi(c)
		if err != nil {
			e := fmt.Errorf("error converting id value of %q : %w", c, err)
			slog.Error(e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		character, err := dbutils.GetQueries().GetCharacterById(r.Context(), int32(id))
		if err != nil {
			e := fmt.Errorf("error retrieving character with id of %d : %w", id, err)
			slog.Error(e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		xpBonus := float64(character.XpBonus) / 100
		xp := int(float64(xpAward) * (1 + xpBonus))

		err = dbutils.GetQueries().InsertXpAward(r.Context(), db.InsertXpAwardParams{
			XpAward:          int32(xpAward),
			XpAwardWithBonus: int32(xp),
			Reason:           reason,
			CharacterID:      character.ID,
		})
		if err != nil {
			e := fmt.Errorf("error saving xp award for character with id %d : %w", id, err)
			slog.Error(e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		character.TotalXp += int32(xp)
		err = dbutils.UpdateCharacter(r.Context(), character)
		if err != nil {
			e := fmt.Errorf("error saving xp award to character record with id %d : %w", id, err)
			slog.Error(e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Redirect to the welcome page upon successful login
	http.Redirect(w, r, "/characters/"+campaignId+"?msg=xp_awarded", http.StatusSeeOther)
}
