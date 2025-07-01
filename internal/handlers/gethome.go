package handlers

import (
	"net/http"

	"todo/internal/templates"
)

func GetHomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := templates.Home()
		layout := templates.Layout(c, "Todo Website")
		layout.Render(r.Context(), w)
	}
}
