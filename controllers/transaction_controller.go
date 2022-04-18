package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TransactionBuyTicket(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func BookingSeats(c *gin.Context) {
	db := connect()
	defer db.Close()

	// get status
	msId := c.PostForm("ms_id")
	var seatStatus int
	seatId := c.PostForm("seat_id")
	studioId := c.PostForm("studio_id")
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

	// buat trans jika belum ada
	println(count)
	if count == 0 {
		personId := c.PostForm("person_id")
		_, errQuery := db.Exec("INSERT INTO transactions (person_id) values (?)", personId)
		if errQuery != nil {
			panic(errQuery)
		}
	}

	// insert detail trans
	_, errQuery := db.Exec("INSERT INTO detail_transactions (transaction_id, seat_id, ms_id) values (?, ?, ?)", transId, seatId, msId)
	if errQuery != nil {
		panic(errQuery)
	}
	// Update status seat
	_, errQuery2 := db.Exec("UPDATE studio_seat SET status = 0 WHERE seat_id = ?", seatId)
	if errQuery2 != nil {
		panic(errQuery2)
	}
	// message sukses
	if errQuery2 != nil {
		ErrorMessage(c, http.StatusBadRequest, "Query Error")
	} else {
		SuccessMessage(c, http.StatusCreated, "Sukses")
	}
}
