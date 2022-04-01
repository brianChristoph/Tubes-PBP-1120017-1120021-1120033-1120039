package main

import (
	c "github.com/Tubes-PBP/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user", c.GetAllUser)
	router.GET("/user/", c.GetUser)
	router.POST("/register", c.Register)
	router.POST("/login", c.Login)
	router.POST("/logout", c.Logout)
	router.PUT("/user/update", c.UpdateUser)
	router.DELETE("/user/delete", c.DeleteUser)

	router.Run("localhost:8080")
}
