package controllers

import (
	"log"
	"strconv"

	"fmt"
	"net/http"

	m "github.com/Tubes-PBP/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	db := connect()
	defer db.Close()

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
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	isAccessTokenValid, userId, email, userType := validateTokenFromCookies(c)
	fmt.Print(email, userType, isAccessTokenValid)

	rows, _ := db.Query("SELECT * FROM persons WHERE id='" + strconv.Itoa(userId) + "'")
	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.JenisKelamin, &user.TanggalLahir, &user.Kewarganegaraan, &user.TipeUser, &user.Status); err != nil {
			log.Print(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if name == "" {
		name = user.name
	}

	if password == "" {
		password = user.password
	}

	if email == "" {
		email = user.email
	}

	// _, errQuery := db.Exec("UPDATE persons SET name = ?, password = ?, email = ? WHERE id=?",
	// 	name,
	// 	password,
	// 	email,
	// 	userId,
	// )

	// var response UserResponse
	// if errQuery == nil {
	// 	response.Message = "Berhasil Memperbaharui Data Pengguna"
	// 	sendSuccessResponse(c, response)
	// } else {
	// 	response.Message = "Gagal Memperbaharui Data Pengguna"
	// 	sendErrorResponse(c, response)
	// }
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
