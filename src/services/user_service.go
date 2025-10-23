package services

import (
	"fmt"

	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/common"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	database := db.GetDb()
	return &UserService{
		logger:     logger,
		cfg:        cfg,
		otpService: NewOtpService(cfg),
		database:   database,
	}
}

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
