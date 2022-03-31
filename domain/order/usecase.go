package order

import "github.com/MikelSot/api-ed-paypal/model"

type Order struct {
	storage Storage
}

func New(s Storage) Order {
	return Order{storage: s}
}

func (o Order) Create(m *model.Order) error {
	return o.storage.Create(m)
}

func (o Order) ByID(ID uint) (model.Order, error) {
	return o.storage.ByID(ID)
}
