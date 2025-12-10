package main

import (
	"context"
	"fmt"
	handler "go-clean-arch/internal/adapter/handler/http"
	"go-clean-arch/internal/adapter/repository"
	"go-clean-arch/internal/core/usecase"
	"go-clean-arch/internal/infraestructure/config"
	"go-clean-arch/internal/infraestructure/postgres"
	"log"
	"os"

	"go.uber.org/zap"
)

func main() {

	//Load Envs
	config, err := config.New()
	if err != nil {
		log.Fatal("Error loading environment variables: ", err)
	}

	//Starting Zap Logs - (Sugar is better for performance)
	var zp *zap.Logger

	switch config.App.Env {
	case "development":
		zp, _ = zap.NewDevelopment()
	default:
		zp, _ = zap.NewProduction()
	}

	defer zp.Sync()
	logger := zp.Sugar()
	logger.Info("Starting the application: ", config.App.Name, "-", config.App.Env)

	ctx := context.Background()

	//Build all external dependencies such as: Database, Message Broker Clients...
	//In this example I will build just the database
	database := postgres.NewDatabase(ctx, config.DB, logger)

	// Dependency Injection: Both this and the process above are SOLID practices,
	// Above we isolated the database creation and now we will inject it in the repository.
	// We are also following a Dependency Inversion Principle from SOLID, where we abstracted the infra
	// By this way, it's easier to test our repository and we don't need to worry with the db client and other dependencies
	// Because our client is on other section (infra), it's also easier to change the DB, we don't need to change de repo, just the infra.
	userRepo := repository.NewRepository(database, logger)

	//Inject the repository into the useCase. (UseCase is responsible for the bussiness rule and don't care about external devices)
	userUseCase := usecase.NewUserService(userRepo, logger)

	//Here you can define your API, if it will be REST,gRPC or other, you just need to inject your useCase.
	h := handler.NewHTTPHandler(userUseCase)

	// Init router
	router, err := handler.NewRouter(
		config.HTTP,
		*h,
	)
	if err != nil {
		logger.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	logger.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		logger.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
