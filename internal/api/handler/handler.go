package handler

import (
	"net/http"

	"go-openapi-gen-playground/internal/api/generated"
)

type TodoAPI struct{}

func NewTodoAPI() *TodoAPI {
	return &TodoAPI{}
}

var ptr = func(s string) *string { return &s }

// Get all To-Do items
// (GET /todos)
func (t TodoAPI) GetTodos(w http.ResponseWriter, r *http.Request) *generated.Response {
	return generated.GetTodosJSON500Response(
		generated.ErrorResponse{
			Error: "Internal Server Error",
		},
	)
}

// Create a new To-Do item
// (POST /todos)
func (t TodoAPI) PostTodos(w http.ResponseWriter, r *http.Request) *generated.Response {
	return generated.PostTodosJSON201Response(
		generated.Todo{
			ID:          "1234",
			Title:       "foo",
			Description: ptr("Bar"),
			Completed:   false,
		},
	)
}

// Delete a specific To-Do item
// (DELETE /todos/{id})
func (t TodoAPI) DeleteTodosID(w http.ResponseWriter, r *http.Request, id string) *generated.Response {
	return &generated.Response{
		Code: 204,
	}
}

// Get a specific To-Do item
// (GET /todos/{id})
func (t TodoAPI) GetTodosID(w http.ResponseWriter, r *http.Request, id string) *generated.Response {
	todo := generated.Todo{
		ID:          id,
		Title:       "foo",
		Description: ptr("bar"),
		Completed:   false,
	}

	return generated.GetTodosIDJSON200Response(todo)
}

// Update a specific To-Do item
// (PUT /todos/{id})
func (t TodoAPI) PutTodosID(w http.ResponseWriter, r *http.Request, id string) *generated.Response {
	return &generated.Response{
		Code: 404,
	}
}
