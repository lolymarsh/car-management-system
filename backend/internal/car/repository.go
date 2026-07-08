package car

import (
	"context"

	"github.com/lolymarsh/car-management-system/internal/model"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, car *model.Car) error {
	_, err := r.db.NewInsert().Model(car).Returning("*").Exec(ctx)
	return err
}

func (r *Repository) GetByID(ctx context.Context, id string) (*model.Car, error) {
	car := new(model.Car)
	err := r.db.NewSelect().Model(car).Where("car_id = ?", id).Where("deleted_at IS NULL").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (r *Repository) List(ctx context.Context, q *bun.SelectQuery) ([]model.Car, int, error) {
	var cars []model.Car
	count, err := q.ScanAndCount(ctx, &cars)
	if err != nil {
		return nil, 0, err
	}
	return cars, count, nil
}

func (r *Repository) Update(ctx context.Context, car *model.Car) error {
	_, err := r.db.NewUpdate().Model(car).WherePK().Returning("*").Exec(ctx)
	return err
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*model.Car)(nil)).
		Set("deleted_at = NOW()").
		Where("car_id = ?", id).
		Exec(ctx)
	return err
}

func (r *Repository) NewSelect() *bun.SelectQuery {
	return r.db.NewSelect()
}
