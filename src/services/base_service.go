package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/mhmojtaba/golang-car-web-api/common"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, err := common.TypeConvertor[T](req)
	tx := s.database.WithContext(ctx).Begin()
	err = tx.Create(&model).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeConvertor[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updatedMap, err := common.TypeConvertor[map[string]interface{}](req)
	(*updatedMap)["updated_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	(*updatedMap)["updated_at"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}
	model := new(T)
	tx := s.database.WithContext(ctx).Begin()
	err = tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(*updatedMap).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return s.GetByID(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.database.WithContext(ctx).Begin()
	model := new(T)

	deletedMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	deletedMap["updated_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	deletedMap["updated_at"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}

	if ctx.Value(constants.UserIdKey) == nil {
		return &service_errors.ServiceError{Message: service_errors.PermissionDenied}
	}

	cnt := tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(deletedMap).
		RowsAffected
	if cnt == 0 {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, service_errors.RecordNotFound, nil)
		return &service_errors.ServiceError{Message: service_errors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetByID(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	err := s.database.
		Where("id = ? AND deleted_by IS NULL", id).
		First(model).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	return common.TypeConvertor[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) GetAll(ctx context.Context) ([]Tr, error) {
	return nil, nil
}
