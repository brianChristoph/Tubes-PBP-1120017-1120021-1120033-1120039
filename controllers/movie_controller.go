package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	m "github.com/Tubes-PBP/models"
)

func DeleteMovieSchedulePeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("DELETE FROM movie_schedules")
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

	idStream := c.Query("ID_Stream")

	row := db.QueryRow("SELECT m.movie_name, m.synopsis, sm.movie_path FROM movies m JOIN streaming_movies sm ON m.id = sm.movie_id WHERE sm.id=?", idStream)

	var movieStream m.StreamingMovie
	if err := row.Scan(&movieStream.MovieName, &movieStream.Synopsis, &movieStream.MoviePath); err != nil {
		panic(err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, movieStream)
	}
}

//MOVIES
func TheaterList(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := ("SELECT * FROM theaters")

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	var theater m.Theater
	var theaters []m.Theater

	for rows.Next() {
		err = rows.Scan(&theater.ID, &theater.TheaterName, &theater.LocationID, &theater.Price)
		if err != nil {
			panic(err.Error())
		}
		theaters = append(theaters, theater)
	}

	if len(theaters) != 0 {
		c.IndentedJSON(http.StatusCreated, theaters)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
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
