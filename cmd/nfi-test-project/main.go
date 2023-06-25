package main

import (
	"fmt"
	httpHandlers "github.com/vikasd145/nfi-test-project/pkg/http-handlers"
	"github.com/vikasd145/nfi-test-project/pkg/repository"
	"github.com/vikasd145/nfi-test-project/pkg/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	Server struct {
		Port int
	}
}

func main() {
	var config Config

	// Load configuration from file
	viper.SetConfigFile("../../configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read configuration file:", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Failed to unmarshal configuration:", err)
	}

	// Connect to the database
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port,
		config.Database.User, config.Database.Password, config.Database.DBName,
	))
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	transactionService := service.NewTransactionService(userRepository)

	router := gin.Default()

	httpHandlers.RegisterHandlers(router, userService, transactionService)

	addr := fmt.Sprintf(":%d", config.Server.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
