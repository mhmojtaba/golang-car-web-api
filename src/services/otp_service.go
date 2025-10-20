package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/data/cache"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

type OptDto struct {
	Otp  string
	Used bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redisClient := cache.GetRedis()
	return &OtpService{
		logger: logger,
		cfg:    cfg,
		redis:  redisClient,
	}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OptDto{
		Otp:  otp,
		Used: false,
	}
	res, err := cache.Get[OptDto](s.redis, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{Message: service_errors.OtpExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{Message: service_errors.OtpUsed}
	}
	err = cache.Set[OptDto](s.redis, key, *val, s.cfg.Otp.ExpireTime*time.Second)

	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) SendOtp(mobileNumber string) error {
	return nil
}

func (s *OtpService) VerifyOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)

	res, err := cache.Get[OptDto](s.redis, key)

	if err != nil {
		return err
	}

	if res.Used {
		return &service_errors.ServiceError{Message: service_errors.OtpUsed}
	}

	if !res.Used && res.Otp != otp {
		return &service_errors.ServiceError{Message: service_errors.InvalidOtp}
	}

	if !res.Used && res.Otp == otp {
		res.Used = true

		if err := cache.Set(s.redis, key, res, s.cfg.Otp.ExpireTime*time.Second); err != nil {
			return err
		}
	}
	return nil
}
