package handlers

import (
	"log/slog"
	"net/http"

	"github.com/hculpan/osetools/cmd/web/templates"
	"github.com/hculpan/osetools/internal/auth"
	"github.com/hculpan/osetools/internal/dbutils"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	user, err := dbutils.GetQueries().GetUser(r.Context(), auth.GetCurrentUsername(r))
	if err != nil {
		slog.Error("failed to retrieve user")
		http.Redirect(w, r, "/info", http.StatusSeeOther)
	}

	comp := templates.AccountTemplate(appTitle, user)
	comp.Render(r.Context(), w)
}
