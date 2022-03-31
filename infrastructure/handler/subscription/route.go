package subscription

import (
	"github.com/MikelSot/api-ed-paypal/domain/subscription"
	"github.com/labstack/echo/v4"
)

const (
	path        = "/v1/subscriptions"
	pathByEmail = "/:email"
)

func NewRouter(e *echo.Echo, useCase subscription.UseCase) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathByEmail, handler.ByEmail)
}
