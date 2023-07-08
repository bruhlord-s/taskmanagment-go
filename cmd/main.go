package main

import (
	"log"

	openboardgo "github.com/bruhlord-s/openboard-go"
	"github.com/bruhlord-s/openboard-go/pkg/handler"
	"github.com/bruhlord-s/openboard-go/pkg/repository"
	"github.com/bruhlord-s/openboard-go/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(openboardgo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured when starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
