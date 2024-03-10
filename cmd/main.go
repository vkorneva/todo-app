package main

import (
	"github.com/vkorneva/todo-app"
	"github.com/vkorneva/todo-app/pkg/handler"
	"github.com/vkorneva/todo-app/pkg/repository"
	"github.com/vkorneva/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
