package car

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lolymarsh/car-management-system/internal/model"
	"github.com/uptrace/bun"
)

type FilterRequest struct {
	Search   string `json:"search"`
	Brand    string `json:"brand,omitempty"`
	Model    string `json:"model,omitempty"`
	Color    string `json:"color,omitempty"`
	YearFrom *int   `json:"year_from,omitempty"`
	YearTo   *int   `json:"year_to,omitempty"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	SortBy   string `json:"sort_by"`
	SortDir  string `json:"sort_dir"`
}

type FilterResponse struct {
	Cars     []model.Car `json:"cars"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, car *model.Car) (*model.Car, error) {
	if car.RegistrationNumber == "" || car.Brand == "" || car.Model == "" {
		return nil, fmt.Errorf("registration_number, brand, and model are required")
	}

	car.CarID = uuid.New().String()
	car.CreatedAt = time.Now()
	car.UpdatedAt = time.Now()

	if err := s.repo.Create(ctx, car); err != nil {
		return nil, err
	}
	return car, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*model.Car, error) {
	car, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("car not found")
	}
	return car, nil
}

func (s *Service) List(ctx context.Context, req FilterRequest) (*FilterResponse, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 10
	}

	q := s.repo.NewSelect().Model((*model.Car)(nil)).Where("deleted_at IS NULL")

	if req.Search != "" {
		like := "%" + req.Search + "%"
		q = q.WhereGroup("AND", func(qq *bun.SelectQuery) *bun.SelectQuery {
			return qq.
				WhereOr("registration_number ILIKE ?", like).
				WhereOr("brand ILIKE ?", like).
				WhereOr("model ILIKE ?", like).
				WhereOr("color ILIKE ?", like)
		})
	}

	if req.Brand != "" {
		q = q.Where("brand ILIKE ?", "%"+req.Brand+"%")
	}
	if req.Model != "" {
		q = q.Where("model ILIKE ?", "%"+req.Model+"%")
	}
	if req.Color != "" {
		q = q.Where("color ILIKE ?", "%"+req.Color+"%")
	}
	if req.YearFrom != nil {
		q = q.Where("year >= ?", *req.YearFrom)
	}
	if req.YearTo != nil {
		q = q.Where("year <= ?", *req.YearTo)
	}

	allowedSorts := map[string]bool{
		"registration_number": true,
		"brand":               true,
		"model":               true,
		"year":                true,
		"created_at":          true,
	}
	sortBy := "created_at"
	if req.SortBy != "" && allowedSorts[req.SortBy] {
		sortBy = req.SortBy
	}
	sortDir := "DESC"
	if strings.EqualFold(req.SortDir, "asc") {
		sortDir = "ASC"
	}
	q = q.OrderExpr(fmt.Sprintf("%s %s", sortBy, sortDir))

	offset := (req.Page - 1) * req.PageSize
	q = q.Limit(req.PageSize).Offset(offset)

	cars, total, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	return &FilterResponse{
		Cars:     cars,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *Service) Update(ctx context.Context, id string, input *model.Car) (*model.Car, error) {
	car, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("car not found")
	}

	if input.RegistrationNumber != "" {
		car.RegistrationNumber = input.RegistrationNumber
	}
	if input.Brand != "" {
		car.Brand = input.Brand
	}
	if input.Model != "" {
		car.Model = input.Model
	}
	if input.Color != "" {
		car.Color = input.Color
	}
	if input.Year != 0 {
		car.Year = input.Year
	}
	if input.Notes != "" {
		car.Notes = input.Notes
	}

	car.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, car); err != nil {
		return nil, err
	}
	return car, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("car not found")
	}
	return s.repo.Delete(ctx, id)
}
