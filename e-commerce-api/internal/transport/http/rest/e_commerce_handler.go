package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/utils/auth"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
	transport "github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/transport/http"
	"github.com/gorilla/mux"
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
		log.Println("SignUphandler :: Error occured while parsing req body: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
	}

	// call service layer
	user, err := h.userService.CreateUser(signUpReq)
	if err != nil {
		log.Println("SignUphandler :: Error occured in creating user: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
	}

	// create a new jwt token with the new user details just created above
	bearerToken, err := auth.GenerateNewToken(user)
	if err != nil {
		log.Println("SignUphandler :: Error occured while creating jwt token: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
	}

	authResponse := dtos.AuthResponse{
		BearerToken: bearerToken,
	}

	// return response back
	transport.SendJsonResponse(w, 201, authResponse, nil)
}

func (h *ECommerceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dtos.UserLoginRequest
}

func (h *ECommerceHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	// call service layer

	// return response back
}

func (h *ECommerceHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	// extract id from request url
	vars := mux.Vars(r)
	productId := vars["productId"]

	// call service layer with product id

	// return response
}

func (h *ECommerceHandler) UpdateProductInventoryHandler(w http.ResponseWriter, r *http.Request) {
	var updateProductInventoryReq dtos.UpdateInventoryRequest

}

func (h *ECommerceHandler) GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category"]

	// call service layer with category name
}

func (h *ECommerceHandler) SortAndFilterProductHandler(w http.ResponseWriter, r *http.Request) {
	var sortAndFilterProductRequest dtos.SortAndFilterProductRequest
}

func (h *ECommerceHandler) SearchProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["keyword"]

	// call service layer with keyword to search
}

func (h *ECommerceHandler) AddNewProductHandler(w http.ResponseWriter, r *http.Request) {
	var addProductReq dtos.AddNewProductRequest
}
