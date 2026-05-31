package domain

// TodoRepositoryInterface is an interface for a repository that stores todos.
type TodoRepositoryInterface interface {
	GetAll() ([]*Todo, error)
	GetByID(id int) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id int) error
}
