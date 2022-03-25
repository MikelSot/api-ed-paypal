package model

import "time"

const (
	Monthly = "Monthly"
	Annual  = "Annual"
)

const (
	Active     = "ACTIVE"
	Terminated = "TERMINATED"
)

// Subscription structure
type Subscription struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	TypeSubs  string    `json:"type_subs"`
	BeginsAt  time.Time `json:"begins_at"`
	EndsAt    time.Time `json:"ends_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Subscriptions slice of subscription
type Subscriptions []Subscription
