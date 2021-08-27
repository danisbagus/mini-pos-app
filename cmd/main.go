package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danisbagus/mini-pos-app/internal/core/service"
	"github.com/danisbagus/mini-pos-app/internal/handler"
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

	// routing
	router.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)

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
