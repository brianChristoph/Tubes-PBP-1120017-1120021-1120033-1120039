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

// ADMIN
func GetAllUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM persons WHERE user_type != ?", "ADMIN")
	if err != nil {
		log.Fatal(err)
		return
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.UserType, &user.Balance, &user.LastSeen); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{
				"status":  http.StatusNotAcceptable,
				"message": "Number of Column Taken isn't The Same as Models",
				"error":   err,
			})
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"datas":  users,
		})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNoContent,
			"message": "No Data Found",
		})
	}
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()
}

// MEMBER
func UserProfile(c *gin.Context) {
	db := connect()
	defer db.Close()

	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		if user.Name != "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"status":    http.StatusOK,
				"user_data": user,
			})
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "User Data is Nil",
			})
		}
	} else {
		c.IndentedJSON(http.StatusGone, gin.H{
			"status":  http.StatusGone,
			"message": "Token has Expired",
		})
	}
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		var updateProf m.UpdateRegister
		err := c.Bind(&updateProf)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Form not detected",
				"data":    err,
			})
			return
		}
		if updateProf.Password == updateProf.PasswordConfirm {
			_, errQuery := db.Exec("UPDATE persons SET name=?, password=?, email=? WHERE id=?", updateProf.Name, updateProf.Password, updateProf.Email, user.ID)

			if errQuery != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "Query Error",
				})
				return
			} else {
				c.IndentedJSON(http.StatusAccepted, gin.H{
					"status":  http.StatusAccepted,
					"message": "User Has Been Updated",
				})
			}
		} else {
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{
				"status":  http.StatusNotAcceptable,
				"message": "Sam Ting Wong with password and confirm password",
			})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Token Not Found",
		})
	}
}

func BuyVIP(c *gin.Context) {
	db := connect()
	defer db.Close()

	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		if user.Balance >= 50000 {
			_, errQuery := db.Exec("UPDATE persons SET status=?, balance=? WHERE id=?", "VIP", (user.Balance - 50000), user.ID)
			if errQuery != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "Query Error",
				})
			}
		} else {
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{
				"status":  http.StatusNotAcceptable,
				"message": "Either Failed to Get data Or You are Poor",
			})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Token Not Found",
		})
	}
}

// General Function
func Logout(c *gin.Context) {
	s.ResetUserToken(c.Writer)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Logged Out",
	})
}

func Register(c *gin.Context) {
	db := connect()
	defer db.Close()

	var register m.UpdateRegister
	err := c.Bind(&register)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Can't connect to form",
		})
	}

	_, errQuery := db.Exec("INSERT INTO persons (name, password, email) VALUES (?,?,?)", register.Name, register.Password, register.Email)

	if errQuery != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Query Error",
		})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User Has Been Created",
		})
	}
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
			c.JSON(http.StatusNotFound, errData.Error())
			return
		}
	}

	if user.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"status":  http.StatusNoContent,
			"message": "Wrong Email Or Password",
		})
		return
	}

	var loginService s.LoginService = s.StaticLoginService(user.Email, user.Password)
	var jwtService s.JWTService = s.JWTAuthService(user.Name)
	var loginController LoginController = LoginHandler(loginService, jwtService)
	token := loginController.Login(c, user)
	if token != "" {
		SetRedis(c, user.Name)
		c.SetCookie(LoadEnv("TOKEN_NAME"), token, 3600, "/user", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Logged In",
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// Background Function
func DeleteUserPeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("DELETE FROM persons WHERE ?-last_seen > 60 AND user_type!='ADMIN' AND user_type!='VIP'", time.Now().Format("YYYY-MM-DD"))
	num, _ := result.RowsAffected()

	if errQuery != nil {
		if num == 0 {
			return
		}
	}
}
