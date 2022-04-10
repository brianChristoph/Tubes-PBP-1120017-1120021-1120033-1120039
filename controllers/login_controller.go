package controllers

import (
	m "github.com/Tubes-PBP/models"
	s "github.com/Tubes-PBP/services"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context, user m.User) string
}

type loginController struct {
	loginService s.LoginService
	jWtService   s.JWTService
}

func LoginHandler(loginService s.LoginService,
	jWtService s.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context, user m.User) string {
	isUserAuthenticated := controller.loginService.LoginUser(user.Email, user.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(user.Password, user.Email, user.UserType, user.Balance)
	}
	return ""
}
