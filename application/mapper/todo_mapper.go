package mapper

import (
	"github.com/filmons/go-ingnerd-backend/application/dto"
	"github.com/filmons/go-ingnerd-backend/domain/todo"
)

// ToTodoDTO maps Todo domain model to TodoDTO.
func ToTodoDTO(t todo.Todo) dto.TodoDTO {
	return dto.TodoDTO{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}
}

// ToTodoDTOs maps slice of Todo to slice of TodoDTO.
func ToTodoDTOs(ts []todo.Todo) []dto.TodoDTO {
	dtos := make([]dto.TodoDTO, len(ts))
	for i, t := range ts {
		dtos[i] = ToTodoDTO(t)
	}
	return dtos
}
