package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Color"},
			},
		},
	}
}

// Create
func (s *CarModelColorService) Create(ctx context.Context, req *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelColorService) Update(ctx context.Context, id int, req *dto.UpdateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelColorService) GetById(ctx context.Context, id int) (*dto.CarModelColorResponse, error) {
	return s.base.GetByID(ctx, id)
}

// Get By Filter
func (s *CarModelColorService) GetByFilter(ctx context.Context, req *dto.PaginationResultWithFilter) (*dto.Pagination[dto.CarModelColorResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
