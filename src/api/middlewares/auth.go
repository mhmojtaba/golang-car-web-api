package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mhmojtaba/golang-car-web-api/api/helper"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
	"github.com/mhmojtaba/golang-car-web-api/services"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)
	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		// get token from header in request
		tokenString := c.GetHeader(constants.AuthorizationHeaderKey)

		token := strings.Split(tokenString, " ")
		if tokenString == "" {
			err = &service_errors.ServiceError{
				Code:    401,
				Message: service_errors.TokenRequired,
				Err:     nil,
			}
			return
		} else {
			_, err := tokenService.VerifyToken(token[1], false)
			if err != nil {
				err = &service_errors.ServiceError{
					Code:    401,
					Message: service_errors.InvalidToken,
					Err:     nil,
				}
				return
			}

			claimMap, err = tokenService.GetClaimsFromToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{
						Code:    401,
						Message: service_errors.TokenExpired,
						Err:     nil,
					}
				default:
					err = &service_errors.ServiceError{
						Code:    401,
						Message: service_errors.InvalidToken,
						Err:     nil,
					}
				}
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, false, helper.AuthError, err, err.Error()),
			)
			return
		}
		c.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		c.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		c.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		c.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		c.Set(constants.EmailKey, claimMap[constants.EmailKey])
		c.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		c.Set(constants.RolesKey, claimMap[constants.RolesKey])
		c.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(
				http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError, service_errors.Forbidden),
			)
			return
		}
		rolesVal := c.Keys[constants.RolesKey]
		fmt.Println(rolesVal)

		if rolesVal == nil {
			c.AbortWithStatusJSON(
				http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError, service_errors.Forbidden),
			)
			return
		}

		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, role := range roles {
			val[role.(string)] = 1
		}

		for _, validRole := range validRoles {
			if _, ok := val[validRole]; ok {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError, service_errors.Forbidden))
	}
}
