package services

import (
	"context"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type CountryService struct {
	// dataBase *gorm.DB
	// logger   logging.Logger
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]{
			database: db.GetDb(),
			logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "Cities"}},
		},
	}
}

// create
func (s *CountryService) CreateCountry(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	// country := models.Country{Name: req.Name}
	// country.CreatedBy = int(ctx.Value(constants.UserIdKey).(float64))
	// country.CreatedAt = time.Now().UTC()

	// tx := s.dataBase.WithContext(ctx).Begin()
	// err := tx.Create(&country).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
	// 	return nil, err
	// }
	// tx.Commit()
	// dto := &dto.CountryResponse{Name: country.Name}
	// return dto, nil

	result, err := s.base.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// update

func (s *CountryService) UpdateCountry(ctx context.Context, countryId uint, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	// updatedMap := map[string]interface{}{
	// 	"name":       req.Name,
	// 	"updated_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
	// 	"updated_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	// }

	// tx := s.dataBase.WithContext(ctx).Begin()
	// err := tx.
	// 	Model(&models.Country{}).
	// 	Where("id = ?", countryId).
	// 	Updates(updatedMap).
	// 	Error
	// if err != nil {
	// 	tx.Rollback()
	// 	s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
	// 	return nil, err
	// }
	// country := &models.Country{}

	// err = tx.Model(&models.Country{}).
	// 	Where("id = ? AND deleted_by IS NULL", countryId).
	// 	First(country).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
	// 	return nil, err
	// }

	// tx.Commit()
	// dto := &dto.CountryResponse{Name: req.Name, Id: country.Id}
	// return dto, nil

	result, err := s.base.Update(ctx, int(countryId), req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// delete

func (s *CountryService) DeleteCountry(ctx context.Context, countryId uint) error {
	// tx := s.dataBase.WithContext(ctx).Begin()
	// deletedMap := map[string]interface{}{
	// 	"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
	// 	"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	// }
	// err := tx.
	// 	Model(&models.Country{}).
	// 	Where("id = ?", countryId).
	// 	Updates(deletedMap).
	// 	Error
	// if err != nil {
	// 	tx.Rollback()
	// 	s.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
	// 	return err
	// }
	// tx.Commit()
	// return nil
	err := s.base.Delete(ctx, int(countryId))
	if err != nil {
		return err
	}
	return nil
}

// get by id

func (s *CountryService) GetCountryById(ctx context.Context, countryId uint) (*dto.CountryResponse, error) {
	// country := &models.Country{}
	// err := s.dataBase.WithContext(ctx).
	// 	Where("id = ? AND deleted_by IS NULL", countryId).
	// 	First(&country).
	// 	Error
	// if err != nil {
	// 	s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
	// 	return nil, err
	// }
	// dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}
	// return dto, nil
	result, err := s.base.GetByID(ctx, int(countryId))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get all

// func (s *CountryService) GetAllCountries(ctx context.Context) ([]*dto.CountryResponse, error) {
// 	var countries []models.Country
// 	err := s.dataBase.WithContext(ctx).
// 		Where("deleted_by IS NULL").
// 		Find(&countries).
// 		Error
// 	if err != nil {
// 		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
// 		return nil, err
// 	}
// 	var countryDtos []*dto.CountryResponse
// 	for _, country := range countries {
// 		countryDto := &dto.CountryResponse{
// 			Id:   country.Id,
// 			Name: country.Name,
// 		}
// 		countryDtos = append(countryDtos, countryDto)
// 	}
// 	return countryDtos, nil
// }
