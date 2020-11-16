package todos

import (
	"github.com/alexthvest/go-todo/database"
)

// TodoRepository ...
type TodoRepository struct {
	database *database.Database
}

// NewTodoRepository ...
func NewRepository(db *database.Database) *TodoRepository {
	return &TodoRepository{database: db}
}

// All ...
func (r *TodoRepository) All() (todos []Todo, err error) {
	err = r.database.Select(&todos, "SELECT * FROM todos;")
	return todos, err
}

// Add ...
func (r *TodoRepository) Add(todo Todo) error {
	_, err := r.database.NamedExec("INSERT INTO todos (title, completed) VALUES (:title, :completed)", todo)
	return err
}
