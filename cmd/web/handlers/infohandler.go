package handlers

import (
	"net/http"

	"github.com/hculpan/osetools/cmd/web/templates"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	comp := templates.InfoTemplate(appTitle)
	comp.Render(r.Context(), w)
}
