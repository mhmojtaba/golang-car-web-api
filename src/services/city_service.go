package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

// Create
func (s *CityService) Create(ctx context.Context, req *dto.CreateCityRequest) (*dto.CityResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CityService) Update(ctx context.Context, cityId int, req *dto.UpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Update(ctx, cityId, req)
}

func (s *CityService) Delete(ctx context.Context, cityId int) error {
	return s.base.Delete(ctx, cityId)
}

func (s *CityService) GetByID(ctx context.Context, cityId int) (*dto.CityResponse, error) {
	return s.base.GetByID(ctx, cityId)
}

func (s *CityService) GetByFilter(ctx context.Context, filter *dto.PaginationResultWithFilter) (*dto.Pagination[dto.CityResponse], error) {
	return s.base.GetByFilter(ctx, filter)
}
