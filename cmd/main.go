package main

import (
	"log"

	"github.com/Tom-Challenger/go-todo"
	"github.com/Tom-Challenger/go-todo/pkg/handler"
	"github.com/Tom-Challenger/go-todo/pkg/repository"
	"github.com/Tom-Challenger/go-todo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// Считываем файл конфигурации во внутрь объекта viper
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHendler(service)

	srv := new(todo.Server)
	
	if err := srv.Run(viper.GetString("port"), handler.InitRouters()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}