package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/hculpan/osetools/internal/db"
	"github.com/hculpan/osetools/internal/dbutils"
)

type XpBonusEntry struct {
	XpBonus int
	Reason  string
}

func AddCharacterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("campaignId")
	characterName := r.FormValue("characterName")
	playerName := r.FormValue("playerName")
	totalXpStr := r.FormValue("totalXp")

	campaignId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("error parsing campaign id", "msg", err.Error())
		http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)
	}
	totalXp, err := strconv.Atoi(totalXpStr)
	if err != nil {
		slog.Error("error parsing total xp", "msg", err.Error())
		http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)
	}

	xpBonuses := r.Form["xpBonuses[]"]
	xpBonusReasons := r.Form["xpBonusReasons[]"]

	if len(xpBonuses) != len(xpBonusReasons) {
		http.Error(w, "Mismatched XP Bonuses and Reasons", http.StatusBadRequest)
		return
	}

	xpBonusTotal := 0
	var xpBonusEntries []XpBonusEntry
	for i := range xpBonuses {
		xpBonus, err := strconv.Atoi(xpBonuses[i])
		if err != nil {
			http.Error(w, "Invalid XP Bonus", http.StatusBadRequest)
			return
		}
		reason := xpBonusReasons[i]
		xpBonusEntries = append(xpBonusEntries, XpBonusEntry{
			XpBonus: xpBonus,
			Reason:  reason,
		})
		xpBonusTotal += xpBonus
	}

	err = dbutils.GetQueries().InsertCharacter(r.Context(), db.InsertCharacterParams{
		Name:       characterName,
		PlayerName: playerName,
		XpBonus:    int32(xpBonusTotal),
		CampaignID: int32(campaignId),
	})
	if err != nil {
		slog.Error("error persisting character", "msg", err.Error())
		http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)
	}
	dbCharacter, err := dbutils.GetQueries().GetLatestCharacterByID(r.Context())
	if err != nil {
		slog.Error("error retrieving persisted character", "msg", err.Error())
		http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)
	}

	err = dbutils.GetQueries().InsertXpAward(r.Context(), db.InsertXpAwardParams{
		XpAward:          int32(totalXp),
		XpAwardWithBonus: int32(totalXp),
		Reason:           "Initial",
		CharacterID:      dbCharacter.ID,
	})
	if err != nil {
		slog.Error("error persisting character", "msg", err.Error())
		http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)
	}

	for _, e := range xpBonusEntries {
		err := dbutils.GetQueries().InsertXpBonusForCharacter(r.Context(), db.InsertXpBonusForCharacterParams{
			XpBonus:     int32(e.XpBonus),
			Reason:      e.Reason,
			CharacterID: dbCharacter.ID,
		})
		if err != nil {
			slog.Error("error persisting character", "msg", err.Error())
			http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)
		}
	}

	http.Redirect(w, r, "/characters/"+id, http.StatusSeeOther)

	/*
	   comp := templates.AddCharacterForm(appTitle, auth.GetCurrentUsername(r), id, true)
	   comp.Render(r.Context(), w)
	*/
}
