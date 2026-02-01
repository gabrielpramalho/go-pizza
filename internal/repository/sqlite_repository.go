package repository

import (
	"database/sql"
	"fmt"
	"go-pizza/internal/entity"
	"time"

	_ "modernc.org/sqlite"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS orders (
			id TEXT PRIMARY KEY,
			flavor_id TEXT,
			size TEXT,
			client_id TEXT,
			total_price REAL,
			status TEXT,
			created_at DATETIME,
			updated_at DATETIME
		);
	`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) Create(order entity.Order) (entity.Order, error) {
	stmt := `INSERT INTO orders (id, flavor_id, size, client_id, status, total_price, created_at, updated_at) 
             VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(stmt,
		order.ID,
		order.FlavorID,
		order.Size,
		order.ClientID,
		order.Status,
		order.TotalPrice,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (r *SQLiteRepository) GetByID(id string) (entity.Order, error) {
	stmt := `SELECT id, flavor_id, size, client_id, status, total_price, created_at, updated_at FROM orders WHERE id = ?`

	row := r.db.QueryRow(stmt, id)

	var order entity.Order

	err := row.Scan(
		&order.ID,
		&order.FlavorID,
		&order.Size,
		&order.ClientID,
		&order.Status,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Order{}, fmt.Errorf("order not found")
		}
		return entity.Order{}, err
	}

	return order, nil
}

func (r *SQLiteRepository) UpdateStatus(id string, status string) (entity.Order, error) {
	stmt := `UPDATE orders SET status = ?, updated_at = ? WHERE id = ?`

	now := time.Now()

	_, err := r.db.Exec(stmt, status, now, id)
	if err != nil {
		return entity.Order{}, err
	}

	return r.GetByID(id)
}

func (r *SQLiteRepository) FindAll() ([]entity.Order, error) {
	stmt := `SELECT id, flavor_id, size, client_id, status, total_price, created_at, updated_at FROM orders`

	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {
		var order entity.Order
		err := rows.Scan(
			&order.ID,
			&order.FlavorID,
			&order.Size,
			&order.ClientID,
			&order.Status,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
