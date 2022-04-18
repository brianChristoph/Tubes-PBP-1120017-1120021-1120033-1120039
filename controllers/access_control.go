package controllers

import (
	"net/http"

	s "github.com/Tubes-PBP/services"
	"github.com/gin-gonic/gin"
)

// Wrapper
func Authentication(f func(c *gin.Context), accessType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var name string = GetRedis(c)
		isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
		if isValid {
			if user.UserType != accessType {
				ErrorMessage(c, http.StatusUnauthorized, "You Don't Have Access to This Endpoint")
				return
			} else {
				f(c)
			}
		} else {
			ErrorMessage(c, http.StatusGone, "Token Expired/Not Logged In")
			return
		}
	}
}
