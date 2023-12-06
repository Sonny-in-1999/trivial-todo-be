package main

import (
	"fmt"
	"log"
	"trivial-todo-be/db"
	"trivial-todo-be/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated!")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalln(err)
	}
}
