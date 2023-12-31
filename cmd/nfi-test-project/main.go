package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	httpHandlers "github.com/vikasd145/nfi-test-project/pkg/http-handlers"
	"github.com/vikasd145/nfi-test-project/pkg/repository"
	"github.com/vikasd145/nfi-test-project/pkg/service"
	"log"
	"net/http"
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
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read configuration file:", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Failed to unmarshal configuration:", err)
	}
	dataSourceName := fmt.Sprintf(
		"postgres://%v:%v@%v/%v?sslmode=disable",
		config.Database.User,
		config.Database.Password,
		"postgres:5432",
		config.Database.DBName,
	)
	// Connect to the database
	db, err := sqlx.Connect("postgres", dataSourceName)
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
	//gin.SetMode(gin.ReleaseMode)
	httpHandlers.RegisterHandlers(router, userService, transactionService)

	addr := fmt.Sprintf(":%d", config.Server.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
