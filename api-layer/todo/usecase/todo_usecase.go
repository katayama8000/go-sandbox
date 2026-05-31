package usecase

import (
	"example.com/api-layer/todo/domain"
)

// TodoUsecase provides the business logic for todos.
type TodoUsecase struct {
	repo domain.TodoRepositoryInterface
}

// NewTodoUsecase creates a new TodoUsecase.
func NewTodoUsecase(repo domain.TodoRepositoryInterface) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

// GetAllTodos returns all todos.
func (u *TodoUsecase) GetAllTodos() ([]*domain.Todo, error) {
	return u.repo.GetAll()
}

// CreateTodo creates a new todo.
func (u *TodoUsecase) CreateTodo(title string) error {
	todo := &domain.Todo{
		Title: title,
		Done:  false,
	}
	return u.repo.Create(todo)
}

// CompleteTodo marks a todo as done.
func (u *TodoUsecase) CompleteTodo(id int) error {
	todo, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	todo.Done = true
	return u.repo.Update(todo)
}
