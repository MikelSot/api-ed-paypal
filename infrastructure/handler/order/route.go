package order

import (
	"github.com/MikelSot/api-ed-paypal/domain/order"
	"github.com/labstack/echo/v4"
)

const (
	path       = "/v1/orders"
	pathCreate = ""
	pathByID   = "/:id"
)

func NewRouter(e *echo.Echo, useCase order.UseCase) {
	handler := New(useCase)

	g := e.Group(path)
	g.POST(pathCreate, handler.Create)
	g.GET(pathByID, handler.ByID)
}
