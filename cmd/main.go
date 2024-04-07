package main

import (
	"log"

	"github.com/kutsuna/simple-go-backend-template/internal/application/service"
	"github.com/kutsuna/simple-go-backend-template/internal/infrastructure/database"
	"github.com/kutsuna/simple-go-backend-template/internal/infrastructure/repository"
	"github.com/kutsuna/simple-go-backend-template/internal/interface/controller"
	"github.com/kutsuna/simple-go-backend-template/internal/interface/router"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	productRepo := repository.NewSQLiteProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	e := echo.New()
	router.SetupRoutes(e, productController)
	e.Start(":8080")
}
