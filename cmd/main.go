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
	merchantRepo := repo.NewMerchantRepo(client)
	customerRepo := repo.NewCustomerRepo(client)
	outletRepo := repo.NewOutletRepo(client)
	priceRepo := repo.NewPriceRepo(client)
	productRepo := repo.NewProductRepo(client)
	purchaseTransactionRepo := repo.NewPurchaseTransactionRepo(client)
	saleTransactionRepo := repo.NewSaleTransactionRepo(client)
	supplierRepo := repo.NewSupplierRepo(client)

	authService := service.NewAuthServie(authRepo)
	merchantService := service.NewMerchantService(merchantRepo, authRepo, productRepo, priceRepo)
	customerService := service.NewCustomerService(customerRepo, authRepo)
	outletService := service.NewOutletService(outletRepo, merchantRepo)
	productService := service.NewProductService(productRepo, merchantRepo, outletRepo, priceRepo)
	purchaseTransactionService := service.NewPurchaseTransactionService(purchaseTransactionRepo, productRepo, merchantRepo, supplierRepo)
	saleTransactionService := service.NewSaleTransactionService(saleTransactionRepo, productRepo, merchantRepo, customerRepo, priceRepo, outletRepo)
	supplierService := service.NewSupplierService(supplierRepo)

	authHandler := handler.AuthHandler{Service: authService}
	merchantHandler := handler.MerchantHandler{Service: merchantService}
	customerHandler := handler.CustomerHandler{Service: customerService}
	outletHandler := handler.OutletHandler{Service: outletService}
	productHandler := handler.ProductHandler{Service: productService}
	purchaseTransactionHandler := handler.PurchaseTransactionHandler{Service: purchaseTransactionService}
	saleTransactionHandler := handler.SaleTransactionHandler{Service: saleTransactionService}
	supplierHandler := handler.SuppplierHandler{Service: supplierService}

	// routing
	authRouter := router.PathPrefix("/auth").Subrouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	authRouter.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/register/merchant", authHandler.RegisterMerchant).Methods(http.MethodPost)
	authRouter.HandleFunc("/register/customer", authHandler.RegisterCustomer).Methods(http.MethodPost)

	apiRouter.HandleFunc("/merchant/me", merchantHandler.GetMerchantDetailMe).Methods(http.MethodGet)
	apiRouter.HandleFunc("/merchant/me", merchantHandler.UpdateMerchantMe).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/merchant/admin", merchantHandler.GetMerchantList).Methods(http.MethodGet)
	apiRouter.HandleFunc("/merchant/{merchant_id}/admin", merchantHandler.GetMerchantDetail).Methods(http.MethodGet)
	apiRouter.HandleFunc("/merchant/{merchant_id}/admin", merchantHandler.RemoveMerchant).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/merchant/{merchant_id}/admin", merchantHandler.UpdateMerchant).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/merchant/{merchant_id}/product", merchantHandler.GetMerchantProductList).Methods(http.MethodGet)

	apiRouter.HandleFunc("/customer/me", customerHandler.GetCustomerDetailMe).Methods(http.MethodGet)
	apiRouter.HandleFunc("/customer/me", customerHandler.UpdateCustomerMe).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/customer/admin", customerHandler.GetCustomerList).Methods(http.MethodGet)
	apiRouter.HandleFunc("/customer/{customer_id}/admin", customerHandler.GetCustomerDetail).Methods(http.MethodGet)
	apiRouter.HandleFunc("/customer/{customer_id}/admin", customerHandler.RemoveCustomer).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/product", productHandler.NewProduct).Methods(http.MethodPost)
	apiRouter.HandleFunc("/product/me", productHandler.GetProductListMe).Methods(http.MethodGet)
	apiRouter.HandleFunc("/product/{sku_id}", productHandler.GetProductDetail).Methods(http.MethodGet)
	apiRouter.HandleFunc("/product/{sku_id}", productHandler.UpdateProduct).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/product/{sku_id}/price", productHandler.UpdateProductPrice).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/product/{sku_id}", productHandler.RemoveProduct).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/outlet", outletHandler.NewOutlet).Methods(http.MethodPost)
	apiRouter.HandleFunc("/outlet/{outlet_id}", outletHandler.UpdateOutlet).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/outlet/{outlet_id}", outletHandler.RemoveOutlet).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/outlet/merchant/{merchant_id}", outletHandler.GetOutletListByMerchantID).Methods(http.MethodGet)

	apiRouter.HandleFunc("/supplier", supplierHandler.GetSuppplierList).Methods(http.MethodGet)

	apiRouter.HandleFunc("/transaction/purchase", purchaseTransactionHandler.NewTransaction).Methods(http.MethodPost)
	apiRouter.HandleFunc("/transaction/sale", saleTransactionHandler.NewTransaction).Methods(http.MethodPost)
	apiRouter.HandleFunc("/transaction/purchase/report", purchaseTransactionHandler.GetTransactionReport).Methods(http.MethodGet)
	apiRouter.HandleFunc("/transaction/sale/report", saleTransactionHandler.GetTransactionReport).Methods(http.MethodGet)
	apiRouter.HandleFunc("/transaction/purchase/product/{sku_id}", purchaseTransactionHandler.GetTransactionProductReport).Methods(http.MethodGet)
	apiRouter.HandleFunc("/transaction/sale/product/{sku_id}", saleTransactionHandler.GetTransactionProductReport).Methods(http.MethodGet)

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
