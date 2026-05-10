package domain

// TodoRepository is an interface for a repository that stores todos.
type TodoRepository interface {
	GetAll() ([]*Todo, error)
	GetByID(id int) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id int) error
}
