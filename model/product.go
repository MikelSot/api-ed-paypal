package model

import "time"

type Product struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Image          string    `json:"image"`
	IsSubscription bool      `json:"is_subscription"`
	Months         uint8     `json:"months"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Products slice of product
type Products []Product
