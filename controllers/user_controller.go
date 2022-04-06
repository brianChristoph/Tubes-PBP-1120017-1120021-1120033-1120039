package controllers

import (
	// m "github.com/Tubes-PBP/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	db := connect()
	defer db.Close()
	
}

func GetAllUser(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func Login(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func Register(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func Logout(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func UserProfile(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func BuyVIP(c *gin.Context) {
	db := connect()
	defer db.Close()
}

// Background Function
func DeleteUserPeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("")
	num, _ := result.RowsAffected()

	if errQuery != nil {
		if num == 0 {
			return
		}
	}
}
