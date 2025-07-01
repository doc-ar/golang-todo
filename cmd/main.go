package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	"todo/internal/handlers"
	"todo/static"
)

func main() {
	ctx := context.Background()
	router := http.NewServeMux()
	connection_string := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(ctx, connection_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	router.Handle("/static/", static.StaticHandler())

	router.HandleFunc("/", handlers.GetHomeHandler())
	router.HandleFunc("GET /todos", handlers.GetTodos(conn))
	router.HandleFunc("POST /todos", handlers.AddTodo(conn))
	router.HandleFunc("PATCH /todos", handlers.CompleteTodo(conn))
	router.HandleFunc("PUT /todos", handlers.EditTodo(conn))
	router.HandleFunc("DELETE /todos", handlers.DeleteTodo(conn))

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
