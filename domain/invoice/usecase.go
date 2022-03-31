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

func (ic Invoice) Create(m *model.Order, subscriptionID uint) error {
	i := model.Invoice{}

	if m.IsSubscription {
		i.SubscriptionID = subscriptionID
	}

	i.Email = m.Email
	i.InvoiceDate = time.Now()
	i.IsProduct = m.IsProduct
	i.IsSubscription = m.IsSubscription
	i.Price = m.Price
	i.ProductID = m.ProductID

	return ic.storage.Create(&i)
}

func (ic Invoice) ByEmail(email string) (model.Invoices, error) {
	return ic.storage.ByEmail(email)
}
