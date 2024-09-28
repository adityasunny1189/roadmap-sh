package rest

import (
	dtos "command-line-arguments/Users/adityapathak/Desktop/learning/golang/roadmap-sh/e-commerce-api/internal/dtos/user_dto.go"
	"encoding/json"
	"net/http"
)

type ECommerceHandler struct {

}

func NewECommerceHandler() *ECommerceHandler {
	return &ECommerceHandler{}
}

func (h *ECommerceHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUpReq dtos.UserSignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&signUpReq); err != nil {
		// handle error
	}

	// call service layer

	// return response back
}

func (h *ECommerceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dtos.UserLoginRequest
}



