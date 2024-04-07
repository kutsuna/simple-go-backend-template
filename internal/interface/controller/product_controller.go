package controller

import (
	"net/http"
	"strconv"

	"github.com/kutsuna/simple-go-backend-template/internal/application/service"
	"github.com/kutsuna/simple-go-backend-template/internal/domain/model"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) GetProductByID(ctx echo.Context) error {
	id := ctx.Param("id")

	product, err := c.service.GetProductByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Product not found")
	}

	return ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) GetAllProducts(ctx echo.Context) error {
	products, err := c.service.GetAllProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) CreateProduct(ctx echo.Context) error {
	var product model.Product
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid product data")
	}

	newProduct, err := c.service.CreateProduct(product.Name, product.Price)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(http.StatusCreated, newProduct)
}

func (c *ProductController) UpdateProduct(ctx echo.Context) error {
	id := ctx.Param("id")

	var product model.Product
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid product data")
	}

	updatedProduct, err := c.service.UpdateProduct(id, product.Name, product.Price)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(http.StatusOK, updatedProduct)
}

func (c *ProductController) DeleteProduct(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	err = c.service.DeleteProduct(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.NoContent(http.StatusNoContent)
}
