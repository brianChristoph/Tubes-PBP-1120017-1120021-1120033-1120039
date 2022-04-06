package controllers

import (
	"log"

	m "github.com/Tubes-PBP/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	email := c.Query("email")
	query := "SELECT * FROM user"

	if email != "" {
		query += " WHERE email ='" + email + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.UserType, &user.Balance, &user.LastSeen); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}
	// var response UsersResponse
	// if len(users) != 0 {
	// 	response.Message = "Berhasil Mendapatkan Data Pengguna"
	// 	response.Data = users
	// 	sendSuccessResponse(c, response)
	// } else {
	// 	response.Message = "Gagal Mendapatkan Data Pengguna"
	// 	sendErrorResponse(c, response)
	// }
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
