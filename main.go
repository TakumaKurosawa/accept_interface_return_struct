package main

import (
	"fmt"
	"sample/service"
	"sample/usecase"
)

func main() {
	svc := service.New()
	uc := usecase.New(svc)

	user := uc.GetUserByID("takumaron")
	fmt.Println(user)
}
