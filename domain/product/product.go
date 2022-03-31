package product

import "github.com/MikelSot/api-ed-paypal/model"

type Storage interface {
	All() (model.Products, error)
	ByID(ID uint) (model.Product, error)
}

type UseCase interface {
	All() (model.Products, error)
	ByID(ID uint) (model.Product, error)
}
