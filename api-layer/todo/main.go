package main

import (
	"fmt"
	"log"

	"example.com/api-layer/todo/infrastructure"
	"example.com/api-layer/todo/usecase"
)

func main() {
	// Initialize the repository
	repo := infrastructure.NewInMemoryTodoRepository()

	// Initialize the usecase
	uc := usecase.NewTodoUsecase(repo)

	// --- Demonstrate the use cases ---

	// Create a new todo
	err := uc.CreateTodo("Buy milk")
	if err != nil {
		log.Fatal(err)
	}
	err = uc.CreateTodo("Walk the dog")
	if err != nil {
		log.Fatal(err)
	}

	// Get all todos
	todos, err := uc.GetAllTodos()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Todos:")
	for _, todo := range todos {
		fmt.Printf("- ID: %d, Title: %s, Done: %t\n", todo.ID, todo.Title, todo.Done)
	}

	// Complete a todo
	err = uc.CompleteTodo(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("After completing a todo:")
	todos, err = uc.GetAllTodos()
	if err != nil {
		log.Fatal(err)
	}
	for _, todo := range todos {
		fmt.Printf("- ID: %d, Title: %s, Done: %t\n", todo.ID, todo.Title, todo.Done)
	}
}
