package controllers

import (

	"github.com/gin-gonic/gin"
)

func DeleteMovieSchedulePeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("")
	num, _ := result.RowsAffected()

	if errQuery != nil {
		if num == 0 {
			return
		}
	}
}

//STREAMING
func ShowStreamingList(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func StreamingMovie(c *gin.Context) {
	db := connect()
	defer db.Close()
}

//MOVIES
func TheaterList(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func ViewMovieDescription(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func ShowMovieList(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func ShowTheaterForCertainMovie(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func ChangePrice(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func UpdateMovieDescription(c *gin.Context) {
	db := connect()
	defer db.Close()
}
