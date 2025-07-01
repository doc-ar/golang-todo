package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	db "todo/internal/db/queries"
)

func EditTodo(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(conn)

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		completed, err := strconv.ParseBool(r.FormValue("completed"))
		if err != nil {
			println(err)
			http.Error(w, "Invalid value for 'completed'", http.StatusBadRequest)
			return
		}

		updatedTodo := db.UpdateTodoParams{
			ID:        uuid.MustParse(r.FormValue("id")),
			Value:     r.FormValue("value"),
			Completed: completed,
		}

		queries.UpdateTodo(r.Context(), updatedTodo)
	}
}
