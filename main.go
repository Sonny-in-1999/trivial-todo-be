package main

import (
	"trivial-todo-be/controller"
	"trivial-todo-be/db"
	"trivial-todo-be/repository"
	"trivial-todo-be/router"
	"trivial-todo-be/usecase"
	"trivial-todo-be/validator"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
