package router

import (
	"github.com/kutsuna/simple-go-backend-template/internal/interface/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, productController *controller.ProductController) {
	e.GET("/products", productController.GetAllProducts)
	e.GET("/products/:id", productController.GetProductByID)
	e.POST("/products", productController.CreateProduct)
	e.PUT("/products/:id", productController.UpdateProduct)
	e.DELETE("/products/:id", productController.DeleteProduct)
}
