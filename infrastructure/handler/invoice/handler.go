package invoice

import (
	"github.com/MikelSot/api-ed-paypal/domain/invoice"
	"github.com/MikelSot/api-ed-paypal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Invoice struct {
	useCase invoice.Usecase
}

func New(uc invoice.Usecase) Invoice {
	return Invoice{useCase: uc}
}

func (h Invoice) ByEmail(c echo.Context) error {
	email := c.Param("email")

	invoices, err := h.useCase.ByEmail(email)
	if err != nil {
		msg := map[string]string{
			"error":    "no se pudo consultar las facturas",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]model.Invoices{"data": invoices}
	return c.JSON(http.StatusOK, msg)
}
