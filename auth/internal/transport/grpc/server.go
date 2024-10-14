package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	authpb "github.com/adityasunny1189/protorepo/protogen/go/auth/v1"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/ports"
)

type GrpcHandler struct {
	authService ports.AuthService
	grpcPort    int
	server      *grpc.Server
	authpb.AuthServiceServer
}

func NewGrpcHandler(authService ports.AuthService, grpcPort int) *GrpcHandler {
	return &GrpcHandler{
		authService: authService,
		grpcPort:    grpcPort,
	}
}

func (h *GrpcHandler) Start() {
	var err error

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", h.grpcPort))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Server up and running on port: ", h.grpcPort)

	grpcServer := grpc.NewServer()
	h.server = grpcServer
	reflection.Register(grpcServer)

	authpb.RegisterAuthServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lst); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}

func (h *GrpcHandler) Stop() {
	h.server.Stop()
}
