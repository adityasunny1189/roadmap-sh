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
	middleware "github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/middleware/auth"
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
	user := middleware.GetUserFromContext(r.Context())
	if user == nil {
		transport.SendErrorResponse(w, dtos.TOAST, "access denied")
		return
	}

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

func (h *ECommerceHandler) CreateCartHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r.Context())
	if user == nil {
		transport.SendErrorResponse(w, dtos.TOAST, "access denied")
		return
	}

	var createCartReq dtos.CreateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&createCartReq); err != nil {
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	cart, err := h.cartService.CreateCart(createCartReq)
	if err != nil {
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	createCartRes := dtos.CreateCartResponse{
		Cart: cart,
	}

	transport.SendJsonResponse(w, 201, createCartRes, nil)
}

func (h *ECommerceHandler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId := vars["cartId"]

	// call service layer to find the cart
	cart, cartItems, err := h.cartService.GetCartDetails(cartId)
	if err != nil {
		log.Println("GetCartHandler :: Error occured while getting cart details: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// create the response object
	getCartDetailsResponse := dtos.GetCartDetailsResponse{
		CartDetails: dtos.CartDto{
			CartMetadata: cart,
			CartItems:    cartItems,
		},
	}

	// return the response
	transport.SendJsonResponse(w, 200, getCartDetailsResponse, nil)
}

func (h *ECommerceHandler) UpdateCartHandler(w http.ResponseWriter, r *http.Request) {
	var updateCartReq dtos.UpdateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&updateCartReq); err != nil {
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	cart, cartItems, err := h.cartService.UpdateCart(updateCartReq)
	if err != nil {
		log.Println("UpdateCartHandler :: Error occured while updating cart details: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// create the response object
	updateCartRes := dtos.UpdateCartResponse{
		CartDetails: dtos.CartDto{
			CartMetadata: cart,
			CartItems:    cartItems,
		},
	}

	// return the response
	transport.SendJsonResponse(w, 200, updateCartRes, nil)
}

func (h *ECommerceHandler) DeleteCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId := vars["cartId"]

	cart, err := h.cartService.DeleteCart(cartId)
	if err != nil {
		log.Println("DeleteCartHandler :: Error occured while deleting cart: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	deleteCartRes := dtos.DeleteCartResponse{
		Cart: cart,
	}

	transport.SendJsonResponse(w, 200, deleteCartRes, nil)
}

func (h *ECommerceHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var createOrderReq dtos.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&createOrderReq); err != nil {
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	order, err := h.checkoutService.CreateOrder(createOrderReq)
	if err != nil {
		log.Println("CreateOrderHandler :: Error occured while creating order: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	createOrderRes := dtos.CreateOrderResponse{
		Order: order,
	}

	transport.SendJsonResponse(w, 201, createOrderRes, nil)
}

func (h *ECommerceHandler) GetAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r.Context())
	if user == nil {
		transport.SendErrorResponse(w, dtos.TOAST, "access denied")
		return
	}

	orders, err := h.checkoutService.GetAllOrders(string(rune(user.ID)))
	if err != nil {
		log.Println("GetAllOrdersHandler :: Error occured while getting orders: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	getAllOrdersRes := dtos.GetAllOrdersResponse{
		Orders: orders,
	}

	transport.SendJsonResponse(w, 200, getAllOrdersRes, nil)
}

func (h *ECommerceHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	// call service layer to find the cart
	order, items, err := h.checkoutService.GetOrderDetails(orderId)
	if err != nil {
		log.Println("GetOrderHandler :: Error occured while getting order details: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// create the response object
	getOrderDetailsResponse := dtos.GetOrderDetailsResponse{
		OrderDetails: order,
		Items:        items,
	}

	// return the response
	transport.SendJsonResponse(w, 200, getOrderDetailsResponse, nil)
}

func (h *ECommerceHandler) GetOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	// call service layer to find the cart
	order, err := h.checkoutService.GetOrderStatus(orderId)
	if err != nil {
		log.Println("GetOrderStatusHandler :: Error occured while getting order details: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	// create the response object
	getOrderStatusResponse := dtos.OrderStatusPollingResponse{
		Status: string(order.OrderState),
	}

	// return the response
	transport.SendJsonResponse(w, 200, getOrderStatusResponse, nil)
}

func (h *ECommerceHandler) InitiatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var paymentReq dtos.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentReq); err != nil {
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	paymentDetails, err := h.checkoutService.InititatePayment(paymentReq)
	if err != nil {
		log.Println("InitiatePaymentHandler :: Error occured while payment: ", err)
		transport.SendErrorResponse(w, dtos.TOAST, err.Error())
		return
	}

	paymentRes := dtos.PaymentResponse{
		Status: string(paymentDetails.PaymentState),
	}

	transport.SendJsonResponse(w, 200, paymentRes, nil)
}
