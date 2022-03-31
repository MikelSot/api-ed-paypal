package subscription

import "github.com/MikelSot/api-ed-paypal/model"

type Storage interface {
	Create(s *model.Subscription) error
	ByEmail(email string) (model.Subscriptions, error)
}

type UseCase interface {
	ByEmail(email string) (model.Subscriptions, error)
}
