package main

import (
	"fmt"
	"log"
	"net/http"
	"taskfy/internal/domain"
	handler "taskfy/internal/handlers"
	"taskfy/internal/middleware"
	"taskfy/internal/repository"
	"taskfy/internal/usecase"
)

func main() {

	userInMemoryRepository := repository.NewUserRepositoryInMemory()
	userUseCase := usecase.NewUserUseCase(userInMemoryRepository)
	taskInMemoryRepository := repository.NewTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskInMemoryRepository)

	u, err := userUseCase.CreateUser("teste@example.com", "123")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u)

	t := domain.NewTask("Estudar", u.Id, true)

	fmt.Println(t)

	userHandlers := handler.NewUserHandler(userUseCase)
	taskHandlers := handler.NewTaskHandler(*taskUseCase)

	mux := http.NewServeMux()

	authMiddlware := middleware.NewAuthMiddleware(userInMemoryRepository)
	mux.HandleFunc("POST /users", userHandlers.CreateUser)
	mux.HandleFunc("POST /login", userHandlers.Login)
	mux.HandleFunc("GET /tasks", authMiddlware.VerifyAuthentication(taskHandlers.ListTasks))

	port := ":8081"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
