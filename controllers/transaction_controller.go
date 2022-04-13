package controllers

import (
	// m "github.com/Tubes-PBP/models"
	"github.com/gin-gonic/gin"
)

func TransactionBuyTicket(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func BookingSeats(c *gin.Context) {
	db := connect()
	defer db.Close()
}
