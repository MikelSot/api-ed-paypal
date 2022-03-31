package handler

import (
	"database/sql"
	useCaseInvoice "github.com/MikelSot/api-ed-paypal/domain/invoice"
	useCaseOrder "github.com/MikelSot/api-ed-paypal/domain/order"
	useCaseProduct "github.com/MikelSot/api-ed-paypal/domain/product"
	useCaseSubscription "github.com/MikelSot/api-ed-paypal/domain/subscription"
	"github.com/MikelSot/api-ed-paypal/infrastructure/handler/invoice"
	"github.com/MikelSot/api-ed-paypal/infrastructure/handler/order"
	"github.com/MikelSot/api-ed-paypal/infrastructure/handler/product"
	"github.com/MikelSot/api-ed-paypal/infrastructure/handler/subscription"
	storageInvoice "github.com/MikelSot/api-ed-paypal/infrastructure/postgres/invoice"
	storageOrder "github.com/MikelSot/api-ed-paypal/infrastructure/postgres/order"
	storageProduct "github.com/MikelSot/api-ed-paypal/infrastructure/postgres/product"
	storageSubscription "github.com/MikelSot/api-ed-paypal/infrastructure/postgres/subscription"
	"github.com/labstack/echo/v4"
)

func Invoice(e *echo.Echo, db *sql.DB) {
	useCase := useCaseInvoice.New(storageInvoice.New(db))
	invoice.NewRouter(e, useCase)
}

func Order(e *echo.Echo, db *sql.DB) {
	useCase := useCaseOrder.New(storageOrder.New(db))
	order.NewRouter(e, useCase)
}

func Product(e *echo.Echo, db *sql.DB) {
	useCase := useCaseProduct.New(storageProduct.New(db))
	product.NewRouter(e, useCase)
}

func Subscription(e *echo.Echo, db *sql.DB) {
	useCase := useCaseSubscription.New(storageSubscription.New(db))
	subscription.NewRouter(e, useCase)
}
