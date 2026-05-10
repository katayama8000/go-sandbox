package infrastructure

import (
	"errors"
	"sync"

	"example.com/api-layer/todo/domain"
)

// InMemoryTodoRepository is an in-memory implementation of the TodoRepository.
type InMemoryTodoRepository struct {
	mu      sync.RWMutex
	todos   map[int]*domain.Todo
	counter int
}

// NewInMemoryTodoRepository creates a new InMemoryTodoRepository.
func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos:   make(map[int]*domain.Todo),
		counter: 0,
	}
}

// GetAll returns all todos.
func (r *InMemoryTodoRepository) GetAll() ([]*domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	todos := make([]*domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// GetByID returns a todo by its ID.
func (r *InMemoryTodoRepository) GetByID(id int) (*domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	todo, ok := r.todos[id]
	if !ok {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

// Create creates a new todo.
func (r *InMemoryTodoRepository) Create(todo *domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	todo.ID = r.counter
	r.todos[todo.ID] = todo
	return nil
}

// Update updates an existing todo.
func (r *InMemoryTodoRepository) Update(todo *domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todos[todo.ID]; !ok {
		return errors.New("todo not found")
	}
	r.todos[todo.ID] = todo
	return nil
}

// Delete deletes a todo by its ID.
func (r *InMemoryTodoRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todos[id]; !ok {
		return errors.New("todo not found")
	}
	delete(r.todos, id)
	return nil
}
