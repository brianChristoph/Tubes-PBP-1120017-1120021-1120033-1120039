package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	m "github.com/Tubes-PBP/models"
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

	query := ("SELECT id,movie_name,thumbnail_path FROM movies")

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	var movie m.Movie
	var movies []m.Movie

	for rows.Next() {
		err = rows.Scan(&movie.ID, &movie.Movie_name, &movie.Thumbnail_path)
		if err != nil {
			panic(err.Error())
		}
		movies = append(movies, movie)
	}

	if len(movies) != 0 {
		c.IndentedJSON(http.StatusOK, movies)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
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
