package domain

import (
	"context"
	"time"
)

type HistoryDetail struct {
	ID          int64     `db:"id"`
	VehicleID   int64     `db:"vehicle_id"`
	CustomerID  int64     `db:"customer_id"`
	PIC         string    `db:"pic"`
	PlateNumber string    `db:"plate_number"`
	Notes       string    `db:"notes"`
	CreatedAt   time.Time `db:"created_at"`
}

type HistoryRepository interface {
	FindByVehicle(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, detail *HistoryDetail) error
}

type HistoryService interface {
}

type HistoryData struct {
	VehicleID   int64  `json:"vehicle_id"`
	CustomerID  int64  `json:"customer_id"`
	PIC         string `json:"pic"`
	PlateNumber string `json:"plate_number"`
	Notes       string `json:"notes"`
	ComeAt      string `json:"come_at"`
}
