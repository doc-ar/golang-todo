package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5"

	db "todo/internal/db/queries"
)

func AddTodo(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(conn)

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		newtodo := r.FormValue("new-todo")
		queries.CreateTodo(r.Context(), newtodo)

		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	}
}
