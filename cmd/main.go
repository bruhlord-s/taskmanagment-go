package main

import (
	"log"

	openboardgo "github.com/bruhlord-s/openboard-go"
	"github.com/bruhlord-s/openboard-go/pkg/handler"
	"github.com/bruhlord-s/openboard-go/pkg/repository"
	"github.com/bruhlord-s/openboard-go/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(openboardgo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error when starting server: %s", err.Error())
	}
}
