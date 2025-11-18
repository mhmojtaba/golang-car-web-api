package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type TokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)

	return &TokenService{
		logger: logger,
		cfg:    cfg,
	}
}

func (t *TokenService) GenerateToken(token *TokenDto) (*dto.TokenDetails, error) {
	tokenDetail := &dto.TokenDetails{}
	tokenDetail.AccessTokenExpireTime = time.Now().Add(t.cfg.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetail.RefreshTokenExpireTime = time.Now().Add(t.cfg.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims[constants.UserIdKey] = token.UserId
	accessTokenClaims[constants.FirstNameKey] = token.FirstName
	accessTokenClaims[constants.LastNameKey] = token.LastName
	accessTokenClaims[constants.MobileNumberKey] = token.MobileNumber
	accessTokenClaims[constants.EmailKey] = token.Email
	accessTokenClaims[constants.RolesKey] = token.Roles
	accessTokenClaims[constants.ExpireTimeKey] = tokenDetail.AccessTokenExpireTime

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(t.cfg.Jwt.SecretKey))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[constants.UserIdKey] = token.UserId
	refreshTokenClaims[constants.ExpireTimeKey] = tokenDetail.RefreshTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(t.cfg.Jwt.RefreshSecretKey))
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

func (t *TokenService) VerifyToken(tokenString string, isRefreshToken bool) (*jwt.Token, error) {
	secretKey := t.cfg.Jwt.SecretKey
	if isRefreshToken {
		secretKey = t.cfg.Jwt.RefreshSecretKey
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &service_errors.ServiceError{Message: service_errors.UnExpectedError}
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func (t *TokenService) GetClaimsFromToken(token *jwt.Token) (jwt.MapClaims, error) {

// }

func (t *TokenService) GetClaimsFromToken(token string) (ClaimMap map[string]interface{}, err error) {
	ClaimMap = make(map[string]interface{})
	parsedToken, err := t.VerifyToken(token, false)
	if err != nil {
		return nil, err
	}

	verify, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, &service_errors.ServiceError{Message: service_errors.ClaimsNotFound}
	}

	for key, val := range verify {
		ClaimMap[key] = val
	}

	return ClaimMap, nil
}
