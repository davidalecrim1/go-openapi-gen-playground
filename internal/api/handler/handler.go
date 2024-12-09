package handler

import (
	"net/http"

	"go-openapi-gen-playground/internal/api/generated"
)

type TodoAPI struct{}

func NewTodoAPI() *TodoAPI {
	return &TodoAPI{}
}

// Get all To-Do items
// (GET /todos)
func (t TodoAPI) GetTodos(w http.ResponseWriter, r *http.Request) *generated.Response {
	panic("not implemented") // TODO: Implement
}

// Create a new To-Do item
// (POST /todos)
func (t TodoAPI) PostTodos(w http.ResponseWriter, r *http.Request) *generated.Response {
	panic("not implemented") // TODO: Implement
}

// Delete a specific To-Do item
// (DELETE /todos/{id})
func (t TodoAPI) DeleteTodosID(w http.ResponseWriter, r *http.Request, id string) *generated.Response {
	panic("not implemented") // TODO: Implement
}

// Get a specific To-Do item
// (GET /todos/{id})
func (t TodoAPI) GetTodosID(w http.ResponseWriter, r *http.Request, id string) *generated.Response {
	panic("not implemented") // TODO: Implement
}

// Update a specific To-Do item
// (PUT /todos/{id})
func (t TodoAPI) PutTodosID(w http.ResponseWriter, r *http.Request, id string) *generated.Response {
	panic("not implemented") // TODO: Implement
}
