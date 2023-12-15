package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, id int64) (Customer, error)
	FindByIDs(ctx context.Context, id []int64) ([]Customer, error)
	FindByPhone(ctx context.Context, phone string) (Customer, error)
	Insert(ctx context.Context, customer *Customer) error
}

type CustomerService interface {
	All(ctx context.Context) ApiResponse
	Save(ctx context.Context, customer CustomerData) ApiResponse
}

type CustomerData struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
