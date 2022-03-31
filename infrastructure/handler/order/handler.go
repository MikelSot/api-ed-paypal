package order

import (
	"github.com/MikelSot/api-ed-paypal/domain/order"
	"github.com/MikelSot/api-ed-paypal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Order struct {
	useCase order.UseCase
}

func New(uc order.UseCase) Order {
	return Order{useCase: uc}
}

func (h Order) Create(c echo.Context) error {
	o := model.Order{}
	err := c.Bind(&o)
	if err != nil {
		msg := map[string]string{
			"error":    "la estructura de la orden no es correcta",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, msg)
	}

	err = h.useCase.Create(&o)
	if err != nil {
		msg := map[string]string{
			"error":    "no pudimos crear la orden",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, map[string]model.Order{"message": o})
}

func (h Order) ByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := map[string]string{
			"error": "no pudimos obtener el id",
			"bad":   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	data, err := h.useCase.ByID(uint(id))
	if err != nil {
		msg := map[string]string{
			"error":    "no pudimos consultar la orden",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]model.Order{"data": data}
	return c.JSON(http.StatusOK, msg)
}
