package handlers

import (
	"net/http"
	"time"

	"github.com/hculpan/osetools/cmd/web/templates"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	comp := templates.LogoutTemplate(appTitle)
	comp.Render(r.Context(), w)
}
