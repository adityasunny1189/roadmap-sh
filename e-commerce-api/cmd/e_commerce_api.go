package cmd

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/config"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/core/services"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/middleware/auth"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/storage/database"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/storage/repository"
	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/transport/http/rest"
	"github.com/gorilla/mux"
)

func RunECommerceAPI() {
	r := mux.NewRouter()

	// Load and Initiate all configs
	cfg := config.NewConfig()
	db := database.Load(cfg)

	// Declare all the repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	checkoutRepo := repository.NewCheckoutRepository(db)

	// Declare all the services
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo)
	checkoutService := services.NewCheckoutService(checkoutRepo)

	// Declare all the middlewares
	r.Use(auth.Middleware(userRepo))

	// Declare all the handlers
	handler := rest.NewECommerceHandler(userService, productService, cartService, checkoutService)

	// Declare all the routes
	authSubroute := r.PathPrefix("/auth").Subrouter()
	authSubroute.HandleFunc("/signup", handler.SignUp).Methods("POST")
	authSubroute.HandleFunc("/login", handler.Login).Methods("POST")

	productSubroute := r.PathPrefix("/products").Subrouter()
	productSubroute.HandleFunc("/", handler.GetAllProductsHandler).Methods("GET")
	productSubroute.HandleFunc("/add", handler.AddNewProductHandler).Methods("POST")
	productSubroute.HandleFunc("/inventory/{productId}", handler.UpdateProductInventoryHandler).Methods("POST")
	productSubroute.HandleFunc("/{category}", handler.GetProductsByCategoryHandler).Methods("GET")
	productSubroute.HandleFunc("/sort", handler.SortAndFilterProductHandler).Methods("POST")
	productSubroute.HandleFunc("/{productId}", handler.GetProductByIdHandler).Methods("GET")
	productSubroute.HandleFunc("/search/{keyword}", handler.SearchProductHandler).Methods("GET")

	cartSubroute := r.PathPrefix("/carts").Subrouter()
	cartSubroute.HandleFunc("/{cartId}", handler.GetCartHandler).Methods("GET")
	cartSubroute.HandleFunc("/create", handler.CreateCartHandler).Methods("POST")
	cartSubroute.HandleFunc("/update", handler.UpdateCartHandler).Methods("PUT")
	cartSubroute.HandleFunc("/delete/{cartId}", handler.DeleteCartHandler).Methods("DELETE")

	checkoutSubroute := r.PathPrefix("/checkout").Subrouter()
	checkoutSubroute.HandleFunc("/orders/create", handler.CreateOrderHandler).Methods("POST")
	checkoutSubroute.HandleFunc("/pay", handler.InitiatePaymentHandler).Methods("POST")
	checkoutSubroute.HandleFunc("/orders", handler.GetAllOrdersHandler).Methods("GET")
	checkoutSubroute.HandleFunc("/orders/{orderId}", handler.GetOrderHandler).Methods("GET")
	checkoutSubroute.HandleFunc("/orders/poll/{orderId}", handler.GetOrderStatusHandler).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic("Error: ", err)
	}
}

/**
	E-commerce-api Features

	1. Login
	2. Signup
	3. View all products
	4. View products by category
	5. Sort products
	6. View product by id
	7. Search products by keyword
	8. Create cart
	9. Update cart by adding or removing product
	10. Create order
	11. Proceed to pay
	12. See all past orders
	13. See a particular order by id
	14. Polling api to get order status


	Auth -> JWT Token based authentication

	Cache -> Use Redis for caching

	Database -> MySQL

	Payment -> Stripe

	Frontend  -> Sveltekit & Tailwind

**/
