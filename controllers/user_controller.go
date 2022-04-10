package controllers

import (
	"fmt"
	"log"
	"time"

	"net/http"

	m "github.com/Tubes-PBP/models"
	s "github.com/Tubes-PBP/services"
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

func UserProfile(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func BuyVIP(c *gin.Context) {
	db := connect()
	defer db.Close()
}

// General Function
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

	row, err := db.Query("SELECT * from persons WHERE password=? AND email=?", user.Password, user.Email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	for row.Next() {
		if errData := row.Scan(
			&user.ID,
			&user.Name,
			&user.Password,
			&user.Email,
			&user.UserType,
			&user.Balance,
			&user.LastSeen); errData != nil {
			c.IndentedJSON(http.StatusNotFound, errData.Error())
			return
		}
	}

	var loginService s.LoginService = s.StaticLoginService(user.Email, user.Password)
	var jwtService s.JWTService = s.JWTAuthService(user.Name)
	var loginController LoginController = LoginHandler(loginService, jwtService)
	token := loginController.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// Background Function
func DeleteUserPeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("DELETE FROM persons WHERE ?-last_seen > 60 AND user_type!='ADMIN'", time.Now().Format("YYYY-MM-DD"))
	num, _ := result.RowsAffected()

	if errQuery != nil {
		if num == 0 {
			return
		}
	}
}
