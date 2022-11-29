package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/Tom-Challenger/go-todo"
	"github.com/Tom-Challenger/go-todo/pkg/handler"
	"github.com/Tom-Challenger/go-todo/pkg/repository"
	"github.com/Tom-Challenger/go-todo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	// Считываем файл конфигурации внутрь объекта viper
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),

		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHendler(service)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRouters()); err != nil {
		logrus.Fatalf("error occured while runing http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
