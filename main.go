package main

import (
	api_tools "github.com/Tubes-PBP/api-tools"
	// controllers
	c "github.com/Tubes-PBP/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	go api_tools.RunBackgroundFunc()

	router.GET("/user", c.GetAllUser)
	router.GET("/user/", c.GetUser)
	router.POST("/register", c.Register)
	router.POST("/login", c.Login)
	router.POST("/logout", c.Logout)
	router.PUT("/user/update/", c.UpdateUser)
	router.DELETE("/user/delete/", c.DeleteUser)

	router.Run("localhost:8080")
}
