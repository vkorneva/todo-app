package main

import (
	"github.com/vkorneva/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)

	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
