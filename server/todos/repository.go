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
func (r *TodoRepository) Add(todo Todo) (*Todo, error) {
	query := "INSERT INTO todos (title, completed) VALUES (:title, :completed) RETURNING id"
	rows, err := r.database.NamedQuery(query, todo)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&todo.Id)
	}

	return &todo, err
}
