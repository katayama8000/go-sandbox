package usecase

import (
	"errors"
	"testing"

	"example.com/api-layer/todo/domain"
)

// MockTodoRepository is a mock implementation of the TodoRepository.
type MockTodoRepository struct {
	GetAllFunc   func() ([]*domain.Todo, error)
	GetByIDFunc  func(id int) (*domain.Todo, error)
	CreateFunc   func(todo *domain.Todo) error
	UpdateFunc   func(todo *domain.Todo) error
	DeleteFunc   func(id int) error
}

func (m *MockTodoRepository) GetAll() ([]*domain.Todo, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, nil
}

func (m *MockTodoRepository) GetByID(id int) (*domain.Todo, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, errors.New("not implemented")
}

func (m *MockTodoRepository) Create(todo *domain.Todo) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(todo)
	}
	return nil
}

func (m *MockTodoRepository) Update(todo *domain.Todo) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(todo)
	}
	return nil
}

func (m *MockTodoRepository) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func TestTodoUsecase_GetAllTodos(t *testing.T) {
	mockRepo := &MockTodoRepository{}
	uc := NewTodoUsecase(mockRepo)

	expectedTodos := []*domain.Todo{
		{ID: 1, Title: "Test Todo 1", Done: false},
	}

	mockRepo.GetAllFunc = func() ([]*domain.Todo, error) {
		return expectedTodos, nil
	}

	todos, err := uc.GetAllTodos()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	if todos[0].Title != "Test Todo 1" {
		t.Errorf("expected todo title to be 'Test Todo 1', got '%s'", todos[0].Title)
	}
}

func TestTodoUsecase_CreateTodo(t *testing.T) {
	mockRepo := &MockTodoRepository{}
	uc := NewTodoUsecase(mockRepo)

	var createdTodo *domain.Todo
	mockRepo.CreateFunc = func(todo *domain.Todo) error {
		createdTodo = todo
		return nil
	}

	err := uc.CreateTodo("New Todo")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdTodo == nil {
		t.Fatal("expected todo to be created")
	}

	if createdTodo.Title != "New Todo" {
		t.Errorf("expected todo title to be 'New Todo', got '%s'", createdTodo.Title)
	}
}

func TestTodoUsecase_CompleteTodo(t *testing.T) {
	mockRepo := &MockTodoRepository{}
	uc := NewTodoUsecase(mockRepo)

	todoToUpdate := &domain.Todo{ID: 1, Title: "Test Todo", Done: false}

	mockRepo.GetByIDFunc = func(id int) (*domain.Todo, error) {
		if id == 1 {
			return todoToUpdate, nil
		}
		return nil, errors.New("not found")
	}

	var updatedTodo *domain.Todo
	mockRepo.UpdateFunc = func(todo *domain.Todo) error {
		updatedTodo = todo
		return nil
	}

	err := uc.CompleteTodo(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedTodo == nil {
		t.Fatal("expected todo to be updated")
	}

	if !updatedTodo.Done {
		t.Error("expected todo to be done")
	}
}
