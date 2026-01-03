package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "User"},
			},
		},
	}
}

// Create
func (s *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelCommentService) Update(ctx context.Context, id int, req *dto.UpdateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelCommentService) GetById(ctx context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return s.base.GetByID(ctx, id)
}

// Get By Filter
func (s *CarModelCommentService) GetByFilter(ctx context.Context, req *dto.PaginationResultWithFilter) (*dto.Pagination[dto.CarModelCommentResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
