package services

import (
	"fmt"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/common"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	otpService   *OtpService
	tokenService *TokenService
	database     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	database := db.GetDb()
	return &UserService{
		logger:       logger,
		cfg:          cfg,
		otpService:   NewOtpService(cfg),
		tokenService: NewTokenService(cfg),
		database:     database,
	}
}

// login by username
func (u *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (tokenDetails *dto.TokenDetails, err error) {
	var dbUser models.User
	err = u.database.Model(&models.User{}).
		Where("username = ?", req.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&dbUser).Error
	if err != nil {
		u.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	if err != nil {
		u.logger.Error(logging.General, logging.PasswordValidation, err.Error(), nil)
		return nil, &service_errors.ServiceError{Message: service_errors.PermissionDenied}
	}

	tokenDto := TokenDto{UserId: dbUser.Id, FirstName: dbUser.FirstName, LastName: dbUser.LastName, Email: dbUser.Email, MobileNumber: dbUser.Mobile}

	if len(*dbUser.UserRoles) > 0 {
		for _, userRole := range *dbUser.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
		}
	}
	token, err := u.tokenService.GenerateToken(&tokenDto)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// register by username
func (u *UserService) RegisterUserByUsername(req *dto.RegisterUserByUsernameRequest) error {
	user := &models.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Active:    true,
	}
	exist, err := u.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exist {
		return &service_errors.ServiceError{Message: service_errors.EmailExists}
	}

	exist, err = u.existsByUserName(req.Username)
	if err != nil {
		return err
	}
	if exist {
		return &service_errors.ServiceError{Message: service_errors.UsernameExists}
	}

	pass := []byte(req.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return &service_errors.ServiceError{Message: service_errors.UnExpectedError}
	}

	user.Password = string(hashedPassword)

	roleId, err := u.getDefaultRole()
	if err != nil {
		u.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	// transaction start
	tx := u.database.Begin()
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		u.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	if err = tx.Create(&models.UserRole{
		UserId: user.Id,
		RoleId: roleId,
	}).Error; err != nil {
		tx.Rollback()
		u.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	// transaction end

	return nil
}

// verify otp and login or register
func (u *UserService) RegisterLoginByMobileNumber(req *dto.RegisterLoginByMobileRequest) (tokenDetails *dto.TokenDetails, err error) {
	err = u.otpService.VerifyOtp(req.MobileNumber, req.Otp)
	if err != nil {
		return nil, err
	}

	exist, err := u.existsByMobileNumber(req.MobileNumber)
	if err != nil {
		return nil, err
	}

	user := &models.User{Username: req.MobileNumber, Mobile: req.MobileNumber}

	if exist {
		// login flow
		var dbUser models.User
		err = u.database.Model(&models.User{}).
			Where("username = ?", user.Username).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).Find(&dbUser).Error
		if err != nil {
			u.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
			return nil, err
		}
		tokenDto := TokenDto{UserId: dbUser.Id, FirstName: dbUser.FirstName, LastName: dbUser.LastName, Email: dbUser.Email, MobileNumber: dbUser.Mobile}

		if len(*dbUser.UserRoles) > 0 {
			for _, userRole := range *dbUser.UserRoles {
				tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
			}
		}
		token, err := u.tokenService.GenerateToken(&tokenDto)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
	// register flow
	pass := []byte(common.GeneratePassword())
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, &service_errors.ServiceError{Message: service_errors.UnExpectedError}
	}

	user.Password = string(hashedPassword)

	roleId, err := u.getDefaultRole()
	if err != nil {
		u.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return nil, err
	}

	// transaction start
	tx := u.database.Begin()
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		u.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	if err = tx.Create(&models.UserRole{
		UserId: user.Id,
		RoleId: roleId,
	}).Error; err != nil {
		tx.Rollback()
		u.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	// transaction end
	var dbUser models.User
	err = u.database.Model(&models.User{}).
		Where("mobileNumber = ?", req.MobileNumber).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&dbUser).Error
	if err != nil {
		u.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	tokenDto := TokenDto{UserId: dbUser.Id, FirstName: dbUser.FirstName, LastName: dbUser.LastName, Email: dbUser.Email, MobileNumber: dbUser.Mobile}

	if len(*dbUser.UserRoles) > 0 {
		for _, userRole := range *dbUser.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
		}
	}
	token, err := u.tokenService.GenerateToken(&tokenDto)
	if err != nil {
		return nil, err
	}
	return token, nil

}

// send otp
func (u *UserService) SendOtp(req *dto.GetOtpRequest) error {
	// Generate OTP
	otp := common.GenerateOtp()
	// Store OTP
	err := u.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	// Send OTP (e.g., via SMS) - Placeholder for actual sending logic

	// u.logger.Infof("otp %s - has been sent to %s", otp, req.MobileNumber)
	fmt.Printf("otp %s - has been sent to %s\n", otp, req.MobileNumber)

	return nil
}

func (u *UserService) existsByEmail(email string) (bool, error) {
	var exists bool

	if err := u.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).Error; err != nil {
		u.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}

	return exists, nil
}

func (u *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool

	if err := u.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile = ?", mobileNumber).
		Find(&exists).Error; err != nil {
		u.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}

	return exists, nil
}

func (u *UserService) existsByUserName(username string) (bool, error) {
	var exists bool

	if err := u.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).Error; err != nil {
		u.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}

	return exists, nil
}

func (u *UserService) getDefaultRole() (roleId int, err error) {
	if err := u.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil

}
