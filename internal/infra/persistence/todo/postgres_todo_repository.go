package todo

import (
	"database/sql"
	"time"

	"github.com/ismael3s/go-todo/internal/core/domain"
	"github.com/ismael3s/go-todo/internal/infra/utils"
)

// r PostgresTodoRepository ports.TodoRepository

type PostgresTodoRepository struct {
	db *sql.DB
}

func NewPostgresTodoRepository(db *sql.DB) *PostgresTodoRepository {
	return &PostgresTodoRepository{
		db: db,
	}
}

func (r *PostgresTodoRepository) GetTodos() ([]*domain.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todo")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*domain.Todo{}

	for rows.Next() {
		var todo domain.Todo
		var deletedAt sql.NullString

		err := rows.Scan(&todo.ID, &todo.Title, &todo.CreatedAt, &todo.UpdatedAt, &deletedAt)
		if err != nil {
			return nil, err
		}

		todo.DeletedAt = deletedAt.String

		todos = append(todos, &todo)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *PostgresTodoRepository) GetTodo(id string) (*domain.Todo, error) {
	stmnt, err := r.db.Prepare("SELECT * FROM todo WHERE id = $1")

	if err != nil {
		return nil, err
	}

	defer stmnt.Close()

	var todo domain.Todo

	var deletedAt sql.NullString

	err = stmnt.QueryRow(id).Scan(&todo.ID, &todo.Title, &todo.CreatedAt, &todo.UpdatedAt, &deletedAt)

	if err != nil {
		return nil, err
	}

	todo.DeletedAt = deletedAt.String

	return &todo, nil
}

func (r *PostgresTodoRepository) Save(todo *domain.Todo) (*domain.Todo, error) {
	if todo.ID == "" {
		return r.create(todo)
	}

	return r.update(todo)
}

func (r *PostgresTodoRepository) create(createTodo *domain.Todo) (*domain.Todo, error) {
	todo := domain.NewTodo(createTodo.Title)
	stmt, err := r.db.Prepare("INSERT INTO todo (id, title) VALUES ($1, $2)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(todo.ID, todo.Title)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *PostgresTodoRepository) update(todo *domain.Todo) (*domain.Todo, error) {
	stmt, err := r.db.Prepare("UPDATE todo SET title = $2, updated_at = $3 WHERE id = $1")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	todo.UpdatedAt = utils.NowTime()

	_, err = stmt.Exec(todo.ID, todo.Title, todo.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *PostgresTodoRepository) Remove(id string) error {
	stmt, err := r.db.Prepare("UPDATE todo SET deleted_at = $1 WHERE id = $2")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, time.Now().Format(time.RFC3339))

	if err != nil {
		return err
	}

	return nil
}
