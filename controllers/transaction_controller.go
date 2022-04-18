package controllers

import (
	"net/http"
	"strconv"

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

		_, errQuery := db.Exec("UPDATE persons SET balance=? WHERE id=?", (user.Balance - totalPrice), user.ID)
		if errQuery != nil {
			ErrorMessage(c, http.StatusBadRequest, "Query Error")
		} else {
			SuccessMessage(c, http.StatusOK, "Success Buy Ticket")
		}
	} else {
		ErrorMessage(c, http.StatusNotFound, "")
	}
}

func BookingSeats(c *gin.Context) {
	db := connect()
	defer db.Close()

	// get status
	msId := c.PostForm("ms_id")
	var seatStatus int
	seatId := c.PostForm("seat_id")
	studioId := c.PostForm("studio_id")
	theaterId := c.PostForm("theater_id")
	row := db.QueryRow("SELECT DISTINCT(studio_seat.status) from studio_seat JOIN seats ON studio_seat.seat_id = seats.id WHERE seats.id = ? AND studio_id = ?", seatId, studioId)
	if err := row.Scan(&seatStatus); err != nil {
		panic(err.Error())
	}
	if seatStatus == 0 {
		panic("seat sudah di pesan orang lain")
	}

	// Cek trans id
	var count int
	transId := c.PostForm("trans_id")
	row2 := db.QueryRow("SELECT count(*) from transactions WHERE id = ?", transId)
	if err := row2.Scan(&count); err != nil {
		panic(err.Error())
	}

	// buat transaksi jika belum ada
	println(count)
	if count == 0 {
		personId := c.PostForm("person_id")
		_, errQuery := db.Exec("INSERT INTO transactions (person_id) values (?)", personId)
		if errQuery != nil {
			panic(errQuery)
		}
	}

	// insert detail transaksi
	_, errQuery := db.Exec("INSERT INTO detail_transactions (transaction_id, seat_id, studio_id, theater_id, ms_id) values (?, ?, ?, ?, ?)", transId, seatId, studioId, theaterId, msId)
	if errQuery != nil {
		panic(errQuery)
	}
	// Update status seat
	_, errQuery2 := db.Exec("UPDATE studio_seat SET status = 0 WHERE seat_id = ?", seatId)
	if errQuery2 != nil {
		panic(errQuery2)
	}

	// Update quantity
	_, errQuery3 := db.Exec("UPDATE theater_studio SET quantity = quantity-1 WHERE theater_id = ?", theaterId)
	if errQuery3 != nil {
		panic(errQuery3)
	}

	// message sukses
	if errQuery2 != nil {
		ErrorMessage(c, http.StatusBadRequest, "Query Error")
	} else {
		SuccessMessage(c, http.StatusCreated, "Sukses")
		theater_id, err := strconv.Atoi(theaterId)
		TransactionBuyTicket(c, theaterId)
	}
}
