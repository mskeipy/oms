package main

import (
	"context"
	router "dropx/internal/api"
	"dropx/internal/api/handlers"
	"dropx/internal/api/middleware"
	"dropx/internal/infrastructure/database"
	"dropx/internal/repositories"
	"dropx/internal/services"
	"dropx/pkg/config"
	"dropx/pkg/constants"
	"dropx/pkg/job/rabbitmq"
	"dropx/pkg/job/rabbitmq/consumer"
	"dropx/pkg/logger"
	"dropx/pkg/validators"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitApp()
}

func main() {
	logger.InitLogger(config.Global.AppName)

	db := database.PostgresqlDB(context.Background())

	server := setupServer()

	rabbit := rabbitmq.InitRabbitMQ()

	// Init Repositories, Services, Handlers
	initRepository := repositories.NewRepositories(db)
	initService := services.NewAppServices(initRepository, rabbit)
	initHandler := handlers.NewAppHandlers(initService)

	rabbit.Consume(constants.ResetPasswordQueue, consumer.HandleResetPasswordMessage)
	router.SetupRoutes(server, initHandler)

	restPort := fmt.Sprintf(":%s", config.Global.RestPort)
	logger.Logger.Info().Msg("Starting server on port " + restPort)
	if err := server.Run(restPort); err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to start server")
	}
}

func setupServer() *gin.Engine {
	r := gin.Default()
	validators.SetupValidator()
	r.Use(middleware.CORS())
	r.Use(middleware.LoggingMiddleware())
	return r
}
