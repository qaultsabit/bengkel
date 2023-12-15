package customer

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/qaultsabit/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

func (r *repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Order(goqu.I("name").Asc())

	if err := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, err
	}
	return
}

func (r *repository) FindByID(ctx context.Context, id int64) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": id,
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}

	return
}

func (r *repository) FindByIDs(ctx context.Context, ids []int64) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": ids,
	})

	if err := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, err
	}

	return
}

func (r *repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := r.db.From("customer").Where(goqu.Ex{
		"phone": phone,
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}

	return
}

func (r *repository) Insert(ctx context.Context, customer *domain.Customer) error {
	executor := r.db.Insert("customers").
		Rows(goqu.Record{
			"name":       customer.Name,
			"phone":      customer.Phone,
			"created_at": customer.CreatedAt,
		}).
		Returning("id").
		Executor()

	var customerdb domain.Customer
	_, err := executor.ScanStructContext(ctx, customerdb)
	if err != nil {
		return err
	}
	customer.ID = customerdb.ID
	return err
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{db: goqu.New("default", con)}
}
