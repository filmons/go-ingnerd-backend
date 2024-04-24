// package repository

// import (
// 	"github.com/filmons/go-ingnerd-backend/domain/todo"

// 	"gorm.io/gorm"
// )

// type TodoRepository struct {
// 	DB *gorm.DB
// }

// func NewTodoRepository(db *gorm.DB) *TodoRepository {
// 	return &TodoRepository{DB: db}
// }

// func (r *TodoRepository) GetAll() ([]todo.Todo, error) {
// 	var todos []todo.Todo
// 	result := r.DB.Find(&todos)
// 	return todos, result.Error
// }
package repository

import (
	"log"

	"github.com/filmons/go-ingnerd-backend/domain/todo"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAll() ([]todo.Todo, error) {
	var todos []todo.Todo
	result := r.DB.Find(&todos)
	if result.Error != nil {
		log.Printf("Failed to fetch todos: %v", result.Error)
		return nil, result.Error
	}
	log.Println("Fetched all todos successfully")
	return todos, nil
}
