package main

import (
	"fmt"
	"log"
	"net/http"
	handler "taskfy/handlers"
	"taskfy/internal/domain"
	"taskfy/internal/repository"
	"taskfy/internal/usecase"
)

func main() {

	userInMemoryRepository := repository.NewUserRepositoryInMemory()
	userUseCase := usecase.NewUserUseCase(userInMemoryRepository)

	u, err := userUseCase.CreateUser("teste@example.com", "123")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u)

	t := domain.NewTask("Estudar", u.Id, true)

	fmt.Println(t)

	userHandlers := handler.NewUserHandler(userUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandlers.CreateUser)

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
