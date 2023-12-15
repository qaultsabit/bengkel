package domain

import (
	"context"
	"time"
)

type Vehicle struct {
	ID        int64     `db:"id"`
	VIN       string    `db:"vin"`
	Brand     string    `db:"brand"`
	UpdatedAt time.Time `db:"updated_at"`
}

type VehicleRepository interface {
	FindByID(ctx context.Context, id string) (Vehicle, error)
	FindByVIN(ctx context.Context, vin string) (Vehicle, error)
	Insert(ctx context.Context, vehicle *Vehicle) error
}

type VehicleService interface {
	FindHistorical(ctx context.Context, vin string) ApiResponse
}

type VehicleHistorical struct {
	ID        int64         `json:"id"`
	VIN       string        `json:"vin"`
	Brand     string        `json:"brand"`
	Histories []HistoryData `json:"histories"`
}
