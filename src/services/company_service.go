package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

// Create
func (s *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CompanyService) Update(ctx context.Context, id int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CompanyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return s.base.GetByID(ctx, id)
}

// Get By Filter
func (s *CompanyService) GetByFilter(ctx context.Context, req *dto.PaginationResultWithFilter) (*dto.Pagination[dto.CompanyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
