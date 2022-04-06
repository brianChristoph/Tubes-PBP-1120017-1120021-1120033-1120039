package controllers

import (
	"fmt"
	"net/http"

	m "github.com/Tubes-PBP/models"
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

func UserProfile(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func BuyVIP(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func Logout(c *gin.Context) {
	db := connect()
	defer db.Close()

	logout, err := db.Query("SELECT * from persons WHERE id=?") // jangan lupa redis
	if err != nil {
		panic(err.Error())
	} else {
		c.SetCookie("name", "Shimin Li", -1, "/", "localhost", false, true)
		c.IndentedJSON(http.StatusOK, logout)
	}
	defer logout.Close()
}

func Register(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func Login(c *gin.Context) {
	db := connect()
	defer db.Close()

	var user m.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	login, err := db.Query("SELECT * from persons WHERE password=? AND email=?", user.Password, user.Email)

	if err != nil {
		panic(err.Error())
	} else {
		// redis
		// ..
		c.SetCookie("name", "Shimin Li", 1, "/", "localhost", false, true)
		c.IndentedJSON(http.StatusOK, login)
	}
	defer login.Close()
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
