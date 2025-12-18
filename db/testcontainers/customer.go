package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Customer struct {
	Id    int
	Name  string
	Email string
}

type CustomerRepository struct {
	conn *pgx.Conn
}

func NewRepository(ctx context.Context, connectionString string) (*CustomerRepository, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &CustomerRepository{conn: conn}, nil
}

func (r CustomerRepository) CreateCustomer(ctx context.Context, customer Customer) (Customer, error) {
	query := "INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING id"
	err := r.conn.QueryRow(ctx, query, customer.Name, customer.Email).Scan(&customer.Id)
	return customer, err
}

func (r CustomerRepository) GetCustomerByEmail(ctx context.Context, email string) (Customer, error) {
	var customer Customer
	query := "SELECT id, name, email FROM customers WHERE email = $1"
	err := r.conn.QueryRow(ctx, query, email).Scan(&customer.Id, &customer.Name, &customer.Email)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}
