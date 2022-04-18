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
	print(msId)
	var seatStatus int
	seatId := c.PostForm("seat_id")
	print(seatId)
	studioId := c.PostForm("studio_id")
	print(studioId)
	row := db.QueryRow("SELECT status from studio_seat JOIN seats ON studio_seat.seat_id = seats.id WHERE seats.id = ? AND studio_id = ?", seatId, studioId)
	if err := row.Scan(&seatStatus); err != nil {
		print("a")
		panic(err.Error())
	}
	if seatStatus == 0 {
		panic("seat sudah di pesan orang lain")
	}

	// Cek trans id
	var count int
	transId := c.PostForm("trans_id")
	print(transId)
	row2 := db.QueryRow("SELECT count(*) from transactions WHERE id = ?", transId)
	if err := row2.Scan(&count); err != nil {
		print("b")
		panic(err.Error())
	}

	// buat trans jika belum ada
	println(count)
	if count == 0 {
		personId := c.PostForm("person_id")
		_, errQuery := db.Exec("INSERT INTO transactions (person_id) values (?)", personId)
		if errQuery != nil {
			print("c")
			panic(errQuery)
		}
	}

	// insert detail trans
	_, errQuery := db.Exec("INSERT INTO detail_transactions (transaction_id, seat_id, ms_id) values (?, ?, ?)", transId, seatId, msId)
	if errQuery != nil {
		print("d")
		panic(errQuery)
	}
	// Update status seat
	_, errQuery2 := db.Exec("UPDATE studio_seat SET status = 0 WHERE seat_id = ?", seatId)
	if errQuery2 != nil {
		print("e")
		panic(errQuery2)
	}
	// cara masukin message sukses?
	if errQuery != nil {
		ErrorMessage(c, http.StatusBadRequest, "Query Error")
	} else {
		SuccessMessage(c, http.StatusCreated, "Sukses")
	}
}
