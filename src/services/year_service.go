package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type PersianYearService struct {
	base *BaseService[models.CountryYear, dto.CreateYearRequest, dto.UpdateYearRequest, dto.YearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.CountryYear, dto.CreateYearRequest, dto.UpdateYearRequest, dto.YearResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *PersianYearService) Create(ctx context.Context, req *dto.CreateYearRequest) (*dto.YearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *PersianYearService) Update(ctx context.Context, id int, req *dto.UpdateYearRequest) (*dto.YearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PersianYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PersianYearService) GetById(ctx context.Context, id int) (*dto.YearResponse, error) {
	return s.base.GetByID(ctx, id)
}

// Get By Filter
func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationResultWithFilter) (*dto.Pagination[dto.YearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
