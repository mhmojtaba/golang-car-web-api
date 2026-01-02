package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/common"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	database *gorm.DB
	logger   logging.Logger
	Preloads []preload
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
	bm, _ := common.TypeConvertor[models.BaseModel](model)
	return s.GetByID(ctx, int(bm.Id))
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updatedMap, err := common.TypeConvertor[map[string]interface{}](req)
	snakeMap := make(map[string]interface{})
	for k, v := range *updatedMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	snakeMap["updated_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	snakeMap["updated_at"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}
	model := new(T)
	tx := s.database.WithContext(ctx).Begin()
	err = tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(snakeMap).Error
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

func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationResultWithFilter) (*dto.Pagination[Tr], error) {
	return Paginate[T, Tr](req, s.Preloads, s.database)

}

func NewPagedList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.Pagination[T] {
	pl := &dto.Pagination[T]{
		Page:      pageNumber,
		TotalRows: count,
		Data:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNext = pl.Page < pl.TotalPages
	pl.HasPrev = pl.Page > 1
	return pl
}

// paginate
func Paginate[T any, Tr any](pagination *dto.PaginationResultWithFilter, preloads []preload, db *gorm.DB) (*dto.Pagination[Tr], error) {
	model := new(T)
	var items *[]T
	var rowItems *[]Tr
	db = Preload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	var totalRows int64 = 0

	db.
		Where(model).
		Where(query).
		Count(&totalRows)

	err := db.
		Where(model).
		Offset(pagination.GetOffsetLimit()).
		Limit(pagination.GetPageSize()).
		Order(sort).
		Find(&items).Error

	if err != nil {
		return nil, err
	}
	rowItems, err = common.TypeConvertor[[]Tr](items)
	if err != nil {
		return nil, err
	}
	return NewPagedList(rowItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err
}

// getQuery
func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > %s", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}
				}
			}
		}
	}
	return strings.Join(query, " AND ")
}

// getSort
func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

// Preload
func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}
	return db
}
