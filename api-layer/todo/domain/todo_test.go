package domain

import (
	"testing"
)

func TestTodo_NewTodo(t *testing.T) {
	todo := &Todo{
		ID:    1,
		Title: "Test Todo",
		Done:  false,
	}

	if todo.ID != 1 {
		t.Errorf("expected todo ID to be 1, got %d", todo.ID)
	}
	if todo.Title != "Test Todo" {
		t.Errorf("expected todo title to be 'Test Todo', got '%s'", todo.Title)
	}
	if todo.Done != false {
		t.Errorf("expected todo to be not done, got %t", todo.Done)
	}
}
