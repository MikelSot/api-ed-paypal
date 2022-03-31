package subscription

import (
	"github.com/MikelSot/api-ed-paypal/model"
	"time"
)

const (
	month = 1
	year  = 12
)

type Subscription struct {
	storage Storage
}

func New(s Storage) Subscription {
	return Subscription{storage: s}
}

func (s Subscription) Create(m *model.Subscription) error {
	months := month
	if m.TypeSubs == model.Annual {
		months = year
	}

	m.BeginsAt = time.Now()
	m.EndsAt = m.BeginsAt.AddDate(0, months, 0)
	m.Status = model.Active

	return s.storage.Create(m)
}

func (s Subscription) ByEmail(email string) (model.Subscriptions, error) {
	return s.storage.ByEmail(email)
}
