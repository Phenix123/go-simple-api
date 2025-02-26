package repositories

import (
	"database/sql"
	"fmt"
	"orders/internal/models"
)

type OrderRepositoryInterface interface {
	GetOrders() ([]models.Order, error)
	GetOrderById(id int64) (*models.Order, error)
}

type OrderRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetOrders() ([]models.Order, error) {
	rows, err := r.db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.Customer, &order.Total); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderById(id int64) (*models.Order, error) {
	res, err := r.db.Query("SELECT * FROM orders WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	// Assuming `models.Order` has fields: ID, Customer (string), and Total (float64)
	var order models.Order
	if res.Next() {
		if err := res.Scan(&order.ID, &order.Customer, &order.Total); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
	} else {
		return nil, fmt.Errorf("order with id %d not found", id)
	}

	return &order, nil
}
