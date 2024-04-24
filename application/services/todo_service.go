package services

import (
	"github.com/filmons/go-ingnerd-backend/application/dto"
	"github.com/filmons/go-ingnerd-backend/application/mapper"
	"github.com/filmons/go-ingnerd-backend/domain/todo"
)

type TodoService struct {
	repository todo.Repository
}

func NewTodoService(repo todo.Repository) *TodoService {
	return &TodoService{
		repository: repo,
	}
}

func (s *TodoService) GetAllTodos() ([]dto.TodoDTO, error) {
	todos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return mapper.ToTodoDTOs(todos), nil
}
