package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/filmons/go-ingnerd-backend/application/services"
	"github.com/filmons/go-ingnerd-backend/pkg/utils"
)

type TodoController struct {
	Service *services.TodoService
}

func NewTodoController(s *services.TodoService) *TodoController {
	return &TodoController{
		Service: s,
	}
}

func (c *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.Service.GetAllTodos()
	if err != nil {
		utils.Logger.Println("Error getting todos:", err)
		http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
