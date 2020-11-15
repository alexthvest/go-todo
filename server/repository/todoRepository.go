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
	err = r.database.Select(&todos, "SELECT * FROM todos;")

	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Add ...
func (r *TodoRepository) Add(todo database.Todo) error {
	_, err := r.database.NamedExec("INSERT INTO todos (title, completed) VALUES (:title, :completed)", todo)
	return err
}
