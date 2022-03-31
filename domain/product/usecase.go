package product

import "github.com/MikelSot/api-ed-paypal/model"

type Product struct {
	storage Storage
}

func New(s Storage) Product {
	return Product{storage: s}
}

func (p Product) All() (model.Products, error) {
	return p.storage.All()
}

func (p Product) ByID(ID uint) (model.Product, error) {
	return p.storage.ByID(ID)
}
