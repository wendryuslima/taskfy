package main

import (
	"fmt"
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
}
