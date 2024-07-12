package handlers

import (
	"net/http"

	"github.com/hculpan/osetools/cmd/web/templates"
	"github.com/hculpan/osetools/internal/auth"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	comp := templates.HomeTemplate(appTitle, auth.GetCurrentUsername(r), auth.IsAuthorized(r))
	comp.Render(r.Context(), w)
}
