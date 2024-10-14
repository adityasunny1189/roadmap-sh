package main

import (
	"github.com/adityasunny1189/roadmap-sh/auth/internal/commons/config"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/services"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/storage/database"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/storage/repository"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/transport/grpc"
)

func main() {
	// load config
	cfg := config.LoadConfig()

	// initialize database
	db := database.Load(cfg)

	// run db migration script

	// initialize repository
	authRepository := repository.NewAuthRepository(db)

	// initialize service
	authService := services.NewAuthService(authRepository)

	// initialize handler
	grpcHandler := grpc.NewGrpcHandler(authService, 6565)

	// start grpc service
	grpcHandler.Start()
}
