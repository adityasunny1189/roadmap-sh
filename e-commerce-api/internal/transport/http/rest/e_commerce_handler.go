package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/utils"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/utils/auth"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/domain"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/ports"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
	transport "github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/transport/http"
	"github.com/gorilla/mux"
)

type ECommerceHandler struct {
	userService     ports.UserService
	productService  ports.ProductService
	cartService     ports.CartService
	checkoutService ports.CheckoutService
}

func NewECommerceHandler(userService ports.UserService,
	productService ports.ProductService,
	cartService ports.CartService,
	checkoutService ports.CheckoutService) *ECommerceHandler {
	return &ECommerceHandler{
		userService:     userService,
		productService:  productService,
		cartService:     cartService,
		checkoutService: checkoutService,
	}
}

func (h *ECommerceHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUpReq dtos.UserSignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&signUpReq); err != nil {
		log.Println("SignUphandler :: Error occured while parsing req body: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// call service layer
	user, err := h.userService.CreateUser(signUpReq)
	if err != nil {
		log.Println("SignUphandler :: Error occured in creating user: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// create a new jwt token with the new user details just created above
	bearerToken, err := auth.GenerateNewToken(user)
	if err != nil {
		log.Println("SignUphandler :: Error occured while creating jwt token: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	authResponse := dtos.AuthResponse{
		BearerToken: bearerToken,
	}

	// return response back
	transport.SendJsonResponse(w, 201, authResponse, nil)
}

func (h *ECommerceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dtos.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Println("Loginhandler :: Error occured while parsing req body: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// call service layer
	user, err := h.userService.GetUser(loginReq)
	if err != nil {
		log.Println("Loginhandler :: Error occured while getting user: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// create a new jwt token with the user details
	bearerToken, err := auth.GenerateNewToken(user)
	if err != nil {
		log.Println("Loginhandler :: Error occured while creating jwt token: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	authResponse := dtos.AuthResponse{
		BearerToken: bearerToken,
	}

	// return response back
	transport.SendJsonResponse(w, 200, authResponse, nil)
}

func (h *ECommerceHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	// call service layer
	productList, err := h.productService.GetAllProducts()
	if err != nil {
		log.Println("GetAllProductsHandler :: Error occured while getting products: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	productsRes := dtos.ProductsResponse{
		Products: productList,
	}

	// return response back
	transport.SendJsonResponse(w, 200, productsRes, nil)
}

func (h *ECommerceHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	// extract id from request url
	vars := mux.Vars(r)
	productId := vars["productId"]

	// call service layer with product id
	product, err := h.productService.GetProductById(productId)
	if err != nil {
		log.Println("GetProductByIdHandler :: Error occured while getting product: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	productRes := dtos.ProductsResponse{
		Products: []domain.Product{product},
	}

	// return response
	transport.SendJsonResponse(w, 200, productRes, nil)
}

// Auth Required
func (h *ECommerceHandler) UpdateProductInventoryHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header["Authorization"]
	if authHeader != nil {
		bearerToken := authHeader[0]
		userEmail, err := auth.VerifyToken(bearerToken)
		if err != nil {
			transport.SendErrorResponse(w, dtos.TOAST, err.Error())
			return
		}

		if userEmail == utils.ADMIN_EMAIL {
			var updateProductInventoryReq dtos.UpdateInventoryRequest
			if err := json.NewDecoder(r.Body).Decode(&updateProductInventoryReq); err != nil {
				log.Println("UpdateProductInventoryHandler :: Error occured while parsing req body: ", err)
				transport.SendErrorResponse(w, dtos.TOAST, err.Error())
				return
			}

			// call service layer
			productCount, err := h.productService.UpdateProductStock(updateProductInventoryReq)
			if err != nil {
				log.Println("GetProductByIdHandler :: Error occured while updating product inventory: ", err)
				transport.SendErrorResponse(w, dtos.TOAST, err.Error())
				return
			}

			updateInventoryRes := dtos.UpdateInventoryResponse{
				CurrentQuantity: productCount,
			}

			// return response
			transport.SendJsonResponse(w, 200, updateInventoryRes, nil)
		} else {
			transport.SendErrorResponse(w, dtos.TOAST, "you are not authorized to make this change")
		}
	}
}

func (h *ECommerceHandler) GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category"]

	// call service layer with category name
	products, err := h.productService.GetProductsByCategory(categoryName)
	if err != nil {
		log.Println("GetProductsByCategoryHandler :: Error occured while getting products: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	productsRes := dtos.ProductsResponse{
		Products: products,
	}

	// return response
	transport.SendJsonResponse(w, 200, productsRes, nil)
}

func (h *ECommerceHandler) SortAndFilterProductHandler(w http.ResponseWriter, r *http.Request) {
	var sortAndFilterProductRequest dtos.SortAndFilterProductRequest
	if err := json.NewDecoder(r.Body).Decode(&sortAndFilterProductRequest); err != nil {
		log.Println("SortAndFilterProductHandler :: Error occured while parsing req body: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// call service layer
	products, err := h.productService.SortAndFilterProduct(sortAndFilterProductRequest)
	if err != nil {
		log.Println("SortAndFilterProductHandler :: Error occured while getting products: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	productsRes := dtos.ProductsResponse{
		Products: products,
	}

	// send response
	transport.SendJsonResponse(w, 200, productsRes, nil)
}

func (h *ECommerceHandler) SearchProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["keyword"]

	// call service layer with keyword to search
	products, err := h.productService.GetProductsByKeyword(keyword)
	if err != nil {
		log.Println("SearchProductHandler :: Error occured while getting products: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	productsRes := dtos.ProductsResponse{
		Products: products,
	}

	// return response
	transport.SendJsonResponse(w, 200, productsRes, nil)
}

// Auth Required
func (h *ECommerceHandler) AddNewProductHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header["Authorization"]
	if authHeader != nil {
		bearerToken := authHeader[0]
		userEmail, err := auth.VerifyToken(bearerToken)
		if err != nil {
			transport.SendErrorResponse(w, dtos.TOAST, err.Error())
			return
		}

		if userEmail == utils.ADMIN_EMAIL {
			var addProductReq dtos.AddNewProductRequest
			if err := json.NewDecoder(r.Body).Decode(&addProductReq); err != nil {
				log.Println("AddNewProductHandler :: Error occured while parsing req body: ", err)
				transport.SendErrorResponse(w, dtos.TOAST, err.Error())
				return
			}

			// call service layer
			product, err := h.productService.AddProduct(addProductReq)
			if err != nil {
				log.Println("AddNewProductHandler :: Error occured while creating new product: ", err)
				transport.SendErrorResponse(w, dtos.TOAST, err.Error())
				return
			}

			addProductRes := dtos.AddNewProductResponse{
				Product: product,
			}

			// return response
			transport.SendJsonResponse(w, 201, addProductRes, nil)
		} else {
			transport.SendErrorResponse(w, dtos.TOAST, "you are not authorized to make this change")
		}
	}
}
