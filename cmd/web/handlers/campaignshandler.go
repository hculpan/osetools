package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hculpan/osetools/cmd/web/templates"
	"github.com/hculpan/osetools/internal/auth"
	"github.com/hculpan/osetools/internal/db"
	"github.com/hculpan/osetools/internal/dbutils"
)

func CampaignsHandler(w http.ResponseWriter, r *http.Request) {
	authorized := auth.IsAuthorized(r)

	var user db.User
	if authorized {
		tmpuser, err := dbutils.GetQueries().GetUser(r.Context(), auth.GetCurrentUsername(r))
		if err != nil {
			slog.Error(fmt.Errorf("failed to load user: %w", err).Error())
			return
		}
		user = tmpuser
	}

	campaigns, err := dbutils.GetQueries().GetCampaignsWithCharacterCount(r.Context())
	if err != nil {
		slog.Error(fmt.Errorf("failed to load user: %w", err).Error())
		return
	}

	comp := templates.CampaignsTemplate(appTitle, user.Username, campaigns, authorized)
	comp.Render(r.Context(), w)
}
