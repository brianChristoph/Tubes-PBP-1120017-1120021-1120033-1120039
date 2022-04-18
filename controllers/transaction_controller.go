package controllers

import (
	"net/http"

	s "github.com/Tubes-PBP/services"
	"github.com/gin-gonic/gin"
)

func TransactionBuyTicket(c *gin.Context, theater_id int) {
	db := connect()
	defer db.Close()

	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)

	var theaterPrice int
	var amountOfSeat int
	var totalPrice int

	if isValid {
		row1 := db.QueryRow("SELECT price FROM theaters WHERE id = ?", theater_id)

		if err := row1.Scan(&theaterPrice); err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		row2 := db.QueryRow("SELECT COUNT(seat_id) FROM detail_transactions DT "+
			"JOIN transactions T ON DT.transaction_id = T.id "+
			"WHERE T.person_id = ? AND DT.theater_id = ?", user.ID, theater_id)

		if err := row2.Scan(&amountOfSeat); err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		totalPrice = theaterPrice * amountOfSeat

		_, errQuery := db.Exec("UPDATE persons SET balance=? WHERE id=?",(user.Balance - totalPrice), user.ID)
		if errQuery != nil {
			ErrorMessage(c, http.StatusBadRequest, "Query Error")
		}else{
			SuccessMessage(c, http.StatusOK,"Success Buy Ticket")
		}
	} else {
		ErrorMessage(c, http.StatusNotFound, "")
	}
}

func BookingSeats(c *gin.Context) {
	db := connect()
	defer db.Close()
}
