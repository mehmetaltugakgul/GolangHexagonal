package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golangHexagonal/internal/app/handler"
	"golangHexagonal/internal/app/repository"
	"golangHexagonal/internal/app/service"
	"golangHexagonal/internal/infrastructure/database"
	"golangHexagonal/internal/infrastructure/router"
	"os"
)

type Config struct {
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	DBHost string `json:"db_host"`
	DBPort string `json:"db_port"`
	DBName string `json:"db_name"`
}

func main() {
	app := fiber.New()

	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {

		}
	}(configFile)

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService)

	router.SetupRoutes(app, userHandler, authHandler)

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
