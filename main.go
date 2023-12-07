package main

import (
	"trivial-todo-be/controller"
	"trivial-todo-be/db"
	"trivial-todo-be/repository"
	"trivial-todo-be/router"
	"trivial-todo-be/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewUserRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
