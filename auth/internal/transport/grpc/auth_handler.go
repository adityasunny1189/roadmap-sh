package grpc

import (
	"context"
	"log"

	authpb "github.com/adityasunny1189/protorepo/protogen/go/auth/v1"
	"github.com/adityasunny1189/roadmap-sh/auth/internal/core/domain"
	"github.com/google/uuid"
)

func (h *GrpcHandler) SignIn(ctx context.Context, req *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	// validate the request

	// call service layer
	bearerToken, err := h.authService.SignIn(req.Email, req.Password)
	if err != nil {
		log.Println("Error occured during signin: ", err)
		return &authpb.SignInResponse{}, err
	}

	// return response
	return &authpb.SignInResponse{
		BearerToken: bearerToken,
	}, nil
}

func (h *GrpcHandler) SignUp(ctx context.Context, req *authpb.SignUpRequest) (*authpb.SignUpResponse, error) {
	// validate the request

	// create user object
	user := domain.User{
		UserId:   uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password: ", err)
		return &authpb.SignUpResponse{}, err
	}

	user.Password = hashedPassword

	// call service layer
	bearerToken, err := h.authService.SignUp(user)
	if err != nil {
		log.Println("Error occured during signup: ", err)
		return &authpb.SignUpResponse{}, err
	}

	// return response
	return &authpb.SignUpResponse{
		BearerToken: bearerToken,
	}, nil
}
