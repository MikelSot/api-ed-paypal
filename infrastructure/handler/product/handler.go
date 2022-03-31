package product

import (
	"github.com/MikelSot/api-ed-paypal/domain/product"
	"github.com/MikelSot/api-ed-paypal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase product.UseCase
}

func New(uc product.UseCase) Handler {
	return Handler{useCase: uc}
}

func (h Handler) All(c echo.Context) error {
	data, err := h.useCase.All()
	if err != nil {
		msg := map[string]string{
			"error":    "no pudimos consultar la info",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]model.Products{"data": data}
	return c.JSON(http.StatusOK, msg)
}

func (h Handler) ByID(c echo.Context) error {
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
			"error":    "No se pudo consultar el producto",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]model.Product{"data": data}
	return c.JSON(http.StatusOK, msg)
}
