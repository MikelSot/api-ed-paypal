package invoice

import "github.com/MikelSot/api-ed-paypal/model"

type Storage interface {
	Create(i *model.Invoice) error
	ByEmail(email string) (model.Invoices, error)
}

type Usecase interface {
	ByEmail(email string) (model.Invoices, error)
}
