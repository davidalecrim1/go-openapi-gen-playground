package main

import (
	"log"
	"net/http"

	"go-openapi-gen-playground/internal/api/generated"
	"go-openapi-gen-playground/internal/api/handler"
)

func main() {
	todoAPI := handler.NewTodoAPI()
	router := http.NewServeMux()
	router.Handle("/", generated.Handler((todoAPI)))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
