package controllers

import (
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

func Register(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
