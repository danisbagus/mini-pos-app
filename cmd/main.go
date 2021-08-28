package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danisbagus/mini-pos-app/internal/core/service"
	"github.com/danisbagus/mini-pos-app/internal/handler"
	"github.com/danisbagus/mini-pos-app/internal/middleware"
	"github.com/danisbagus/mini-pos-app/internal/repo"

	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// sql driver
	client := GetClient()

	// multiplexer
	router := mux.NewRouter()

	// wiring
	authRepo := repo.NewAuthRepo(client)
	authService := service.NewAuthServie(authRepo)
	authHandler := handler.AuthHandler{Service: authService}

	merchantRepo := repo.NewMerchantRepo(client)
	merchantService := service.NewMerchantService(merchantRepo)
	merchantHandler := handler.MerchantHandler{Service: merchantService}

	customerRepo := repo.NewCustomerRepo(client)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.CustomerHandler{Service: customerService}

	productRepo := repo.NewProductRepo(client)
	productService := service.NewProductService(productRepo, merchantRepo)
	productHandler := handler.ProductHandler{Service: productService}

	// routing
	authRouter := router.PathPrefix("/auth").Subrouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	authRouter.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/register/merchant", authHandler.RegisterMerchant).Methods(http.MethodPost)
	authRouter.HandleFunc("/register/customer", authHandler.RegisterCustomer).Methods(http.MethodPost)

	apiRouter.HandleFunc("/merchant/me", merchantHandler.GetMerchantDetailMe).Methods(http.MethodGet)

	apiRouter.HandleFunc("/customer/me", customerHandler.GetCustomerDetailMe).Methods(http.MethodGet)

	apiRouter.HandleFunc("/product", productHandler.NewProduct).Methods(http.MethodPost)

	// middleware
	authMiddleware := middleware.AuthMiddleware{Repo: repo.NewAuthRepo(client)}
	apiRouter.Use(authMiddleware.AuthorizationHandler())

	// starting server
	appPort := fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))
}

func GetClient() *sqlx.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
