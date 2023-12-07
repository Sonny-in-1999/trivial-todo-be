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
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
