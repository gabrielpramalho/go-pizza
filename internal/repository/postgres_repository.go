package repository

import (
	"database/sql"
	"go-pizza/internal/entity"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(connectionString string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS orders (
			id TEXT PRIMARY KEY,
			flavor_id TEXT,
			size TEXT,
			client_id TEXT,
			status TEXT,
			total_price DECIMAL,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);
	`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Create(order entity.Order) (entity.Order, error) {
	stmt := `INSERT INTO orders (id, flavor_id, size, client_id, status, total_price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(stmt,
		order.ID,
		order.FlavorID,
		order.Size,
		order.ClientID,
		order.Status,
		order.TotalPrice,
		order.CreatedAt,
		order.UpdatedAt)

	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (r *PostgresRepository) GetByID(id string) (entity.Order, error) {
	stmt := `SELECT id, flavor_id, size, client_id, status, total_price, created_at, updated_at FROM orders WHERE id = $1`

	row := r.db.QueryRow(stmt, id)
	var order entity.Order

	err := row.Scan(
		&order.ID, &order.FlavorID, &order.Size, &order.ClientID,
		&order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt,
	)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (r *PostgresRepository) UpdateStatus(id string, status string) (entity.Order, error) {
	stmt := `UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2 RETURNING id, status, updated_at`

	_, err := r.db.Exec(stmt, status, id)
	if err != nil {
		return entity.Order{}, err
	}
	return r.GetByID(id)
}

func (r *PostgresRepository) FindAll() ([]entity.Order, error) {
	stmt := `SELECT id, flavor_id, size, client_id, status, total_price, created_at, updated_at FROM orders`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []entity.Order{}

	for rows.Next() {
		var order entity.Order
		err := rows.Scan(
			&order.ID, &order.FlavorID, &order.Size, &order.ClientID,
			&order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
