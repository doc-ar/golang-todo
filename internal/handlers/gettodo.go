package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	db "todo/internal/db/queries"
	"todo/internal/templates"
)

func GetTodos(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(conn)
		todos, err := queries.ListTodos(r.Context())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching todos: %v\n", err)
		}
		list := templates.Todo_list(todos)
		list.Render(r.Context(), w)
	}
}
