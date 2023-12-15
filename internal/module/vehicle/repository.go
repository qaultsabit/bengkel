package vehicle

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/qaultsabit/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

// FindByID implements domain.VehicleRepository.
func (r *repository) FindByID(ctx context.Context, id string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicle").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

// FindByVIN implements domain.VehicleRepository.
func (r *repository) FindByVIN(ctx context.Context, vin string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicle").Where(goqu.Ex{
		"vin": vin,
	})

	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

// Insert implements domain.VehicleRepository.
func (r *repository) Insert(ctx context.Context, vehicle *domain.Vehicle) error {
	executor := r.db.Insert("vehicles").Rows(goqu.Record{
		"vin":        vehicle.VIN,
		"brand":      vehicle.Brand,
		"updated_at": vehicle.UpdatedAt,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, vehicle)
	return err
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}
