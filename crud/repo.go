package crud

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/unbxd/go-base/utils/log"
)

var (
	ErrIdNotFound = errors.New("id not found")
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo initializes a new repository instance
func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: logger,
	}, nil
}

// CreateCustomer inserts a new customer with all mandatory fields
func (repo *repo) CreateCustomer(ctx context.Context, customer Customer) error {
	if customer.Customerid == "" || customer.Email == "" || customer.Phone == "" {
		return errors.New("all fields (customerid, email, phone) are required")
	}

	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO Customer(customerid, email, phone) VALUES ($1, $2, $3)",
		customer.Customerid, customer.Email, customer.Phone)

	if err != nil {
		fmt.Println("Error occurred inside CreateCustomer in repo:", err)
		return err
	}

	fmt.Println("User Created:", customer.Email)
	return nil
}

// GetCustomerById retrieves a customer by ID
func (repo *repo) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	customer := Customer{}

	err := repo.db.QueryRowContext(ctx,
		"SELECT customerid, email, phone FROM Customer WHERE customerid = $1", id).
		Scan(&customer.Customerid, &customer.Email, &customer.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}
	return customer, nil
}

// GetAllCustomers retrieves all customers
func (repo *repo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	var customers []Customer
	rows, err := repo.db.QueryContext(ctx, "SELECT customerid, email, phone FROM Customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err = rows.Scan(&customer.Customerid, &customer.Email, &customer.Phone)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

// DeleteCustomer removes a customer by ID
func (repo *repo) DeleteCustomer(ctx context.Context, id string) (string, error) {
	res, err := repo.db.ExecContext(ctx, "DELETE FROM Customer WHERE customerid = $1", id)
	if err != nil {
		return "", err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if rowCnt == 0 {
		return "", ErrIdNotFound
	}

	return "Successfully deleted", nil
}

// UpdateCustomer updates customer details
func (repo *repo) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	if customer.Customerid == "" || customer.Email == "" || customer.Phone == "" {
		return "", errors.New("all fields (customerid, email, phone) are required")
	}

	res, err := repo.db.ExecContext(ctx,
		"UPDATE Customer SET email = $1, phone = $2 WHERE customerid = $3",
		customer.Email, customer.Phone, customer.Customerid)

	if err != nil {
		return "", err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", ErrIdNotFound
	}

	return "Successfully updated", nil
}
