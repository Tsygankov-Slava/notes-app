package main

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/Tsygankov-Slava/notes-app/pkg/handler"
	"github.com/Tsygankov-Slava/notes-app/pkg/repository"
	"github.com/Tsygankov-Slava/notes-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// @title Notes App API
// @version 1.0
// @description API Server for Notes Application

// @host localhost:8001
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) // json format more convenient for log collection systems

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil { // file `.env` by default
		logrus.Fatalf("Error loading env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Error initialization database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(notes.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Http server run error: %s", err.Error())
	}
}

// Use `github.com/spf13/viper` library
func initConfig() error {
	configsPath := "configs" // The path to our directory with configuration files
	viper.AddConfigPath(configsPath)
	configName := "config"
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
