package main

import (
	"log"

	"github.com/hipzz/orm-practice/api"
	"github.com/hipzz/orm-practice/controller"
	"github.com/hipzz/orm-practice/database"
	"github.com/hipzz/orm-practice/pkg/config"
	"github.com/hipzz/orm-practice/repository/postgres"
	"github.com/hipzz/orm-practice/service"
)

func main() {
	config := config.New("dev")
	dbConfig := database.Config{
		Host: config.GetString("DB_HOST"),
		Port: config.GetString("DB_PORT"),
		User: config.GetString("DB_USER"),
		Pass: config.GetString("DB_PASS"),
		Name: config.GetString("DB_NAME"),
	}
	pool, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("error connect to database: %s", err)
	}
	userRepository := postgres.New(pool)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	router := api.NewRouter(userController)
	router.Run()
}
