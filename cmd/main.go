package main

import (
	"github.com/MikelSot/api-ed-paypal/infrastructure/handler"
	"github.com/MikelSot/api-ed-paypal/infrastructure/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	loadEnv()
	startWithEcho()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print("no se pudieron cargar las variables de entorno")
		panic(err)
	}
}

func startWithEcho() {
	db, err := postgres.New()
	if err != nil {
		log.Print("no se pudo cargar la base de datos")
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	handler.Product(e, db)
	handler.Order(e, db)
	handler.Subscription(e, db)
	handler.Invoice(e, db)
	//TODO: aun falta integrar paypal

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Print(e.StartTLS(port, os.Getenv("PUBLIC_KEY"), os.Getenv("PRIVATE_KEY")))
}
