package main

import (
	c "github.com/Tubes-PBP/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user", c.GetAllUser)
	router.GET("/user/{user_id}", c.GetUser)
	router.POST("/user/register", c.Register)
	router.PUT("/user/update/{user_id}", c.UpdateUser)
	router.DELETE("/user/delete/{user_id}", c.DeleteUser)

	router.Run("localhost:8080")
}
