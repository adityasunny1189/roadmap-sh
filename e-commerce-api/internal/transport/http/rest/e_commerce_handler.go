package rest

import (
	"encoding/json"
	"net/http"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

type ECommerceHandler struct {
	userService    ports.UserService
	productService ports.ProductService
}

func NewECommerceHandler(userService ports.UserService, productService ports.ProductService) *ECommerceHandler {
	return &ECommerceHandler{
		userService:    userService,
		productService: productService,
	}
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

func (h *ECommerceHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ECommerceHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ECommerceHandler) UpdateProductInventoryHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ECommerceHandler) GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ECommerceHandler) SortAndFilterProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ECommerceHandler) SearchProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ECommerceHandler) AddNewProductHandler(w http.ResponseWriter, r *http.Request) {

}

