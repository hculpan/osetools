package handlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/hculpan/osetools/internal/auth"
	"github.com/hculpan/osetools/internal/db"
	"github.com/hculpan/osetools/internal/dbutils"
)

func AddCampaignPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	description := r.FormValue("description")

	user, err := dbutils.GetQueries().GetUser(r.Context(), auth.GetCurrentUsername(r))
	if err != nil {
		slog.Error(fmt.Errorf("failed to load user: %w", err).Error())
		return
	}

	dbCampaign, _ := dbutils.GetQueries().GetCampaignByName(r.Context(), name)
	if dbCampaign.Name == name {
		slog.Error("error: campaign with that name already exists")
		http.Redirect(w, r, "/campaigns", http.StatusSeeOther)
	}

	err = dbutils.GetQueries().InsertCampaign(r.Context(), db.InsertCampaignParams{
		Name:        name,
		Description: sql.NullString{String: description, Valid: true},
		UserID:      user.ID,
		KeyField:    uuid.New().String(),
	})
	if err != nil {
		slog.Error(fmt.Errorf("failed to insert campaign: %w", err).Error())

	}

	http.Redirect(w, r, "/campaigns", http.StatusSeeOther)

	/*
	   comp := templates.AddCharacterForm(appTitle, auth.GetCurrentUsername(r), id, true)
	   comp.Render(r.Context(), w)
	*/
}
