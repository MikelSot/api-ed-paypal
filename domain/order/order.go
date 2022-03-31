package order

import "github.com/MikelSot/api-ed-paypal/model"

type Storage interface {
	Create(o *model.Order) error
	ByID(ID uint) (model.Order, error)
}

type UseCase interface {
	Create(o *model.Order) error
	ByID(ID uint) (model.Order, error)
}
