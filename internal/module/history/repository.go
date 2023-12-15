package history

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/qaultsabit/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

// FindByVehicle implements domain.HistoryRepository.
func (r *repository) FindByVehicle(ctx context.Context, id int64) (result []domain.HistoryDetail, err error) {
	dataset := r.db.From("history_details").Where(goqu.Ex{
		"vehicle_id": id,
	}).Order(goqu.I("id").Asc())

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// Inser implements domain.HistoryRepository.
func (r *repository) Insert(ctx context.Context, detail *domain.HistoryDetail) error {
	detail.CreatedAt = time.Now()

	executor := r.db.Insert("history_details").Rows(goqu.Record{
		"vehicle_id":  detail.VehicleID,
		"customer_id": detail.CustomerID,
		"notes":       detail.Notes,
		"pic":         detail.PIC,
		"created_at":  detail.CreatedAt,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, detail)
	return err
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{db: goqu.New("default", con)}
}
