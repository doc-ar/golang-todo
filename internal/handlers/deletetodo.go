package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	db "todo/internal/db/queries"
)

func DeleteTodo(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(conn)

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		id, err := uuid.Parse(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Unable to parse uuid", http.StatusBadRequest)
			return
		}

		queries.DeleteTodo(r.Context(), id)
	}
}
