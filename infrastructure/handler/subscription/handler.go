package subscription

import (
	"github.com/MikelSot/api-ed-paypal/domain/subscription"
	"github.com/MikelSot/api-ed-paypal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Subscription struct {
	useCase subscription.UseCase
}

func New(uc subscription.UseCase) Subscription {
	return Subscription{useCase: uc}
}

func (h Subscription) ByEmail(c echo.Context) error {
	email := c.Param("email")

	subscriptions, err := h.useCase.ByEmail(email)
	if err != nil {
		msg := map[string]string{
			"error":    "no se pudo consultar las suscripciones",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]model.Subscriptions{"data": subscriptions}
	return c.JSON(http.StatusOK, msg)
}
