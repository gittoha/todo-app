package main

import (
	todo_app "github.com/gittoha/todo-app"
	"github.com/gittoha/todo-app/pkg/handler"
	"github.com/gittoha/todo-app/pkg/repository"
	"github.com/gittoha/todo-app/pkg/service"
	"log"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
