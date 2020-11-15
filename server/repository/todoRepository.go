package repository

import (
	"github.com/alexthvest/go-todo/database"
)

// TodoRepository ...
type TodoRepository struct {
	database *database.Database
}

// NewTodoRepository ...
func NewTodoRepository(db *database.Database) *TodoRepository {
	return &TodoRepository{database: db}
}

// All ...
func (r *TodoRepository) All() (todos []database.Todo, err error) {
	var results []database.Todo
	err = r.database.Select(&results, "SELECT * FROM todos;")

	if err != nil {
		return nil, err
	}

	return results, nil
}
