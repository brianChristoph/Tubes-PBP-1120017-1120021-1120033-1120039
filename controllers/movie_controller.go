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
func UpdateStreaming(c *gin.Context) {
	db := connect()
	defer db.Close()
}

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

	idMovie := c.Query("ID_Movie")

	//Get Movie Data
	rows := db.QueryRow("SELECT * FROM movies WHERE id = ?", idMovie)

	var movie m.Movie
	if err := rows.Scan(&movie.ID, &movie.Movie_name, &movie.Thumbnail_path, &movie.Synopsis, &movie.Last_premier, &movie.Streamable); err != nil {
		panic(err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, movie)
	}
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
