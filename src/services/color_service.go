package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *ColorService) Create(ctx context.Context, req *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *ColorService) Update(ctx context.Context, id int, req *dto.UpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *ColorService) GetById(ctx context.Context, id int) (*dto.ColorResponse, error) {
	return s.base.GetByID(ctx, id)
}

// Get By Filter
func (s *ColorService) GetByFilter(ctx context.Context, req *dto.PaginationResultWithFilter) (*dto.Pagination[dto.ColorResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
