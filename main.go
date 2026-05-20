package main

import (
	"fmt"
	"taskfy/internal/domain"
)

func main() {
	u := domain.NewUser("teste@example.com", "123")

	fmt.Println(u)

	t := domain.NewTask("Estudar", u.Id, true)

	fmt.Println(t)
}
