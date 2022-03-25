package model

import "time"

// Invoice structure
type Invoice struct {
	ID             uint      `json:"id"`
	InvoiceDate    time.Time `json:"invoice_date"`
	Email          string    `json:"email"`
	IsProduct      bool      `json:"is_product"`
	IsSubscription bool      `json:"is_subscription"`
	ProductID      uint      `json:"product_id"`
	SubscriptionID uint      `json:"subscription_id"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Invoices slice of invoice
type Invoices []Invoice
