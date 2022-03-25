package invoice

import (
	"github.com/MikelSot/api-ed-paypal/model"
	"time"
)

type Invoice struct {
	storage Storage
}

func New(s Storage) Invoice {
	return Invoice{storage: s}
}

func (uc Invoice) Create(order *model.Order, subscriptionID uint) error {
	i := model.Invoice{}

	if order.IsSubscription {
		i.SubscriptionID = subscriptionID
	}

	i.Email = order.Email
	i.InvoiceDate = time.Now()
	i.IsProduct = order.IsProduct
	i.IsSubscription = order.IsSubscription
	i.Price = order.Price
	i.ProductID = order.ProductID

	return uc.storage.Create(&i)
}

func (uc Invoice) ByEmail(email string) (model.Invoices, error) {
	return uc.storage.ByEmail(email)
}
