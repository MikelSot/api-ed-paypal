package order

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MikelSot/api-ed-paypal/infrastructure/postgres"
	"github.com/MikelSot/api-ed-paypal/model"
)

const (
	queryInsert = "INSERT INTO orders (email, is_product, is_subscription, product_id, type_subs, price) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	queryByID   = "SELECT id, email, is_product, is_subscription, product_id, type_subs, price, created_at, updated_at FROM orders WHERE id = $1"
)

type Order struct {
	db *sql.DB
}

func New(db *sql.DB) Order {
	return Order{db: db}
}

func (o Order) Create(m *model.Order) error {
	emptyContext := context.Background()
	stmt, err := o.db.PrepareContext(emptyContext, queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.ExecContext(
		emptyContext,
		m.Email,
		m.IsProduct,
		m.IsSubscription,
		m.ProductID,
		m.TypeSubs,
		m.Price,
	)
	if err != nil {
		return err
	}

	got, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if got != 1 {
		return fmt.Errorf("expected 1 row, got %d", got)
	}

	return nil
}

func (o Order) ByID(ID uint) (model.Order, error) {
	emptyContext := context.Background()
	stmt, err := o.db.PrepareContext(emptyContext, queryByID)
	if err != nil {
		return model.Order{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(emptyContext, ID)

	return o.scan(row)
}

func (o Order) scan(row postgres.RowScanner) (model.Order, error) {
	typeSubsNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	m := model.Order{}

	err := row.Scan(
		&m.ID,
		&m.Email,
		&m.IsProduct,
		&m.IsSubscription,
		&m.ProductID,
		&typeSubsNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return model.Order{}, err
	}

	m.TypeSubs = typeSubsNull.String

	return m, nil
}
