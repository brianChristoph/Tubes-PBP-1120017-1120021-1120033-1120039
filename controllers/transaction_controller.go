package controllers

import (
	"database/sql"
	"net/http"
<<<<<<< HEAD
	"strconv"
=======
	"time"
>>>>>>> 5ca081309f7786f221758e25b1484ac4c4e9d6f5

	m "github.com/Tubes-PBP/models"
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
	var book_seat m.BookingSeats
	err := c.Bind(&book_seat)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, "Failed to Detect Form")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	var seatStatus int
	row := db.QueryRow("SELECT status FROM studio_seat WHERE seat_id = ? AND studio_id = ? AND theater_id = ?", book_seat.Seat_ID, book_seat.Studio_ID, book_seat.Theater_ID)
	if err := row.Scan(&seatStatus); err != nil {
		ErrorMessage(c, http.StatusBadRequest, "Query Error")
	}
	if seatStatus == 0 {
		ErrorMessage(c, http.StatusConflict, "seats taken")
	}

	// buat transaksi jika belum ada
	var name string = GetRedis(c)
	isValid, user := s.JWTAuthService(name).ValidateTokenFromCookies(c.Request)
	if isValid {
		insertTransaction(db, user.ID)
		// insert detail transaksi
		insertDetailTransaction(db, book_seat.Seat_ID, book_seat.Studio_ID, book_seat.Theater_ID, book_seat.Ms_ID, user.ID)

		// Update status seat
		updateSeatStatus(db, book_seat.Seat_ID, book_seat.Studio_ID, book_seat.Theater_ID)

		// Update quantity
		updateQuantity(db, book_seat.Theater_ID)

		// message sukses
		SuccessMessage(c, http.StatusCreated, "Sukses")
	}
	TransactionBuyTicket(c, book_seat.Theater_ID)
}

func insertTransaction(db *sql.DB, id int) {
	_, errQuery := db.Exec("INSERT INTO transactions (person_id, transaction_date) values (?, ?)", id, time.Now().Format("YYYY-MM-DD"))
	if errQuery != nil {
		panic(errQuery)
	}
}

func insertDetailTransaction(db *sql.DB, seat_id int, studio_id int, theater_id int, ms_id int, user_id int) {
	_, errQuery1 := db.Exec("INSERT INTO detail_transactions (transaction_id, seat_id, studio_id, theater_id, ms_id)"+
		" SELECT MAX(transactions.id), ?, ?, ?, ? FROM transactions"+
		" JOIN persons ON persons.id = transactions.person_id"+
		" WHERE persons.id = ?", seat_id, studio_id, theater_id, ms_id, user_id)
	if errQuery1 != nil {
		panic(errQuery1)
	}
}

func updateSeatStatus(db *sql.DB, seat_id int, studio_id int, theater_id int) {
	_, errQuery2 := db.Exec("UPDATE studio_seat SET status = 0 WHERE seat_id = ? AND studio_id = ? AND theater_id = ?", seat_id, studio_id, theater_id)
	if errQuery2 != nil {
		panic(errQuery2)
	}
}

func updateQuantity(db *sql.DB, theater_id int) {
	_, errQuery3 := db.Exec("UPDATE theater_studio SET quantity = quantity-1 WHERE theater_id = ?", theater_id)
	if errQuery3 != nil {
		panic(errQuery3)
	}
}
