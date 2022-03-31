package product

import (
	"github.com/MikelSot/api-ed-paypal/domain/product"
	"github.com/labstack/echo/v4"
)

const (
	path     = "/v1/products"
	pathAll  = ""
	pathByID = "/:id"
)

func NewRouter(e *echo.Echo, useCase product.UseCase) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathAll, handler.All)
	g.GET(pathByID, handler.ByID)
}
