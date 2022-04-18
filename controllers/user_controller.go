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

	rows, err := db.Query("SELECT * FROM persons WHERE user_type != ?", LoadEnv("ADMIN"))
	if err != nil {
		log.Fatal(err)
		return
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.UserType, &user.Balance, &user.LastSeen); err != nil {
			ErrorMessage(c, http.StatusNotAcceptable, "")
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		SuccessMessage(c, http.StatusOK, "")
		c.JSON(http.StatusOK, users)
	} else {
		ErrorMessage(c, http.StatusNoContent, "")
	}
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	idUser := c.Query("ID_User")

	_, errQuery := db.Exec("DELETE FROM persons WHERE ID=?", idUser)

	if errQuery == nil {
		SuccessMessage(c, http.StatusOK, "Person Succesfully Deleted")
	} else {
		ErrorMessage(c, http.StatusNoContent, "Delete Person Failed")
	}
}

// MEMBER
func UserProfile(c *gin.Context) {
	db := connect()
	defer db.Close()

	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		if user.ID != 0 {
			SuccessMessage(c, http.StatusOK, "")
			c.IndentedJSON(http.StatusOK, gin.H{
				"user_data": user,
			})
		} else {
			ErrorMessage(c, http.StatusNoContent, "")
		}
	} else {
		ErrorMessage(c, http.StatusGone, "")
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
			ErrorMessage(c, http.StatusInternalServerError, "Failed to Detect Form")
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		if updateProf.Password == updateProf.PasswordConfirm {
			res, errQuery := db.Exec("UPDATE persons SET name=?, password=?, email=? WHERE idUser=?", updateProf.Name, updateProf.Password, updateProf.Email, user.ID)
			num, _ := res.RowsAffected()

			if num == 0 {
				ErrorMessage(c, http.StatusBadRequest, "No Content Updated")
				return
			}

			if errQuery != nil {
				ErrorMessage(c, http.StatusBadRequest, "Query Error")
			} else {
				SuccessMessage(c, http.StatusOK, "User Has Been Updated")
			}
		} else {
			ErrorMessage(c, http.StatusNotAcceptable, "Wrong Password/Email")
		}
	} else {
		ErrorMessage(c, http.StatusNotFound, "")
	}
}

func BuyVIP(c *gin.Context) {
	db := connect()
	defer db.Close()

	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		if user.Balance >= 50000 {
			_, errQuery := db.Exec("UPDATE persons SET status=?, balance=? WHERE idUser=?", "VIP", (user.Balance - 50000), user.ID)
			if errQuery != nil {
				ErrorMessage(c, http.StatusBadRequest, "Query Error")
			}
		} else {
			ErrorMessage(c, http.StatusNotAcceptable, "You Are Too Poor")
		}
	} else {
		ErrorMessage(c, http.StatusNotFound, "")
	}
}

// General Function
func Logout(c *gin.Context) {
	var name string = GetRedis(c)
	isValid, _ := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		s.ResetUserToken(c.Writer)
		SuccessMessage(c, http.StatusOK, "Logged Out")
	} else {
		ErrorMessage(c, http.StatusBadRequest, "You Are Not Logged In Anyway")
	}
}

func Register(c *gin.Context) {
	db := connect()
	defer db.Close()

	var register m.UpdateRegister
	err := c.Bind(&register)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, "Failed to Detect Form")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	_, errQuery := db.Exec("INSERT INTO persons (name, password, email) VALUES (?,?,?)", register.Name, register.Password, register.Email)

	var userEmail = register.Email
	GoMail(userEmail)
	fmt.Print(userEmail)

	if errQuery != nil {
		ErrorMessage(c, http.StatusBadRequest, "Query Error")
	} else {
		SuccessMessage(c, http.StatusCreated, "User Has Been Created")
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
		ErrorMessage(c, http.StatusNoContent, "Wrong Email/Password")
		return
	}

	var loginService s.LoginService = s.StaticLoginService(user.Email, user.Password)
	var jwtService s.JWTService = s.JWTAuthService(user.Name)
	var loginController LoginController = LoginHandler(loginService, jwtService)
	token := loginController.Login(c, user)
	if token != "" {
		SetRedis(c, user.Name)
		c.SetCookie(LoadEnv("TOKEN_NAME"), token, 3600, "/user", "localhost", false, true)
		SuccessMessage(c, http.StatusOK, "Logged In")
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// Background Function
func DeleteUserPeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("DELETE FROM persons WHERE ?-last_seen > 60 AND user_type!='?' AND user_type!='?'",
		time.Now().Format("YYYY-MM-DD"), LoadEnv("ADMIN"), LoadEnv("VIP"),
	)
	num, _ := result.RowsAffected()

	if errQuery != nil {
		if num == 0 {
			fmt.Print(num, " inactive users have been deleted")
			return
		}
	}
}
