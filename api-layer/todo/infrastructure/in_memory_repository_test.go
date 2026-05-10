package infrastructure

import (
	"testing"

	"example.com/api-layer/todo/domain"
)

func TestInMemoryTodoRepository_Create(t *testing.T) {
	repo := NewInMemoryTodoRepository()
	todo := &domain.Todo{Title: "Test Todo"}
	err := repo.Create(todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if todo.ID != 1 {
		t.Errorf("expected todo ID to be 1, got %d", todo.ID)
	}
}

func TestInMemoryTodoRepository_GetAll(t *testing.T) {
	repo := NewInMemoryTodoRepository()
	repo.Create(&domain.Todo{Title: "Test Todo 1"})
	repo.Create(&domain.Todo{Title: "Test Todo 2"})

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(todos) != 2 {
		t.Fatalf("expected 2 todos, got %d", len(todos))
	}
}

func TestInMemoryTodoRepository_GetByID(t *testing.T) {
	repo := NewInMemoryTodoRepository()
	todo := &domain.Todo{Title: "Test Todo"}
	repo.Create(todo)

	foundTodo, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if foundTodo.ID != 1 {
		t.Errorf("expected todo ID to be 1, got %d", foundTodo.ID)
	}

	_, err = repo.GetByID(2)
	if err == nil {
		t.Error("expected an error for not found todo, got nil")
	}
}

func TestInMemoryTodoRepository_Update(t *testing.T) {
	repo := NewInMemoryTodoRepository()
	todo := &domain.Todo{Title: "Test Todo"}
	repo.Create(todo)

	todo.Title = "Updated Title"
	err := repo.Update(todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	foundTodo, _ := repo.GetByID(1)
	if foundTodo.Title != "Updated Title" {
		t.Errorf("expected updated title, got '%s'", foundTodo.Title)
	}

    // Test updating a non-existent todo
    nonExistentTodo := &domain.Todo{ID: 99, Title: "Non-existent"}
    err = repo.Update(nonExistentTodo)
    if err == nil {
        t.Error("expected an error when updating a non-existent todo, got nil")
    }
}

func TestInMemoryTodoRepository_Delete(t *testing.T) {
	repo := NewInMemoryTodoRepository()
	todo := &domain.Todo{Title: "Test Todo"}
	repo.Create(todo)

	err := repo.Delete(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = repo.GetByID(1)
	if err == nil {
		t.Error("expected an error for deleted todo, got nil")
	}

    // Test deleting a non-existent todo
    err = repo.Delete(99)
    if err == nil {
        t.Error("expected an error when deleting a non-existent todo, got nil")
    }
}
