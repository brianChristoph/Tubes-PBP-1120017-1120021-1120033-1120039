package controllers

import (
	m "github.com/Tubes-PBP/models"
	"github.com/gin-gonic/gin"
)

func SuccessMessage(c *gin.Context, successCode int, message string) {
	var success m.SuccessResponse
	success.Status = successCode
	if message == "" {
		switch successCode {
		case 200:
			success.Message = "Success"
		case 201:
			success.Message = "Data Has Been Created"
		case 302:
			success.Message = "Found"
		default:
			success.Message = "Undeclared Success"
		}
	} else {
		success.Message = message
	}
	c.JSON(successCode, success)
}

func ErrorMessage(c *gin.Context, errCode int, message string) {
	var errMessage m.ErrorResponse
	errMessage.Status = errCode
	if message == "" {
		switch errCode {
		case 204:
			errMessage.Message = "No Content"
		case 400:
			errMessage.Message = "Bad Request"
		case 401:
			errMessage.Message = "Unauthorized Access"
		case 406:
			errMessage.Message = "Response Doesn't Match Models"
		case 410:
			errMessage.Message = "Token Has Expired"
		default:
			errMessage.Message = "Undeclared Error"
		}
	} else {
		errMessage.Message = message
	}
	c.JSON(errCode, errMessage)
}
