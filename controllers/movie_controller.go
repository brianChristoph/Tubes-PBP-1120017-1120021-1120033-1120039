package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	m "github.com/Tubes-PBP/models"
)

func DeleteMovieSchedulePeriodically() {
	db := connect()
	defer db.Close()

	result, errQuery := db.Exec("DELETE FROM movie_schedules WHERE ?-playing_time > 14", time.Now())
	num, _ := result.RowsAffected()

	if errQuery != nil {
		if num != 0 {
			fmt.Print("Movie Schedule ")
			return
		}
	}
}

//STREAMING
func UpdateStreamingMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	id := c.PostForm("id")
	updateMonth, _ := strconv.Atoi(c.PostForm("Update Month"))

	var currentLastPremier time.Time
	row1 := db.QueryRow("SELECT streaming_date_end FROM streaming_movies WHERE id=?",
		id,
	)
	if err := row1.Scan(&currentLastPremier); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	updatedMonth := currentLastPremier.AddDate(0, updateMonth, 0)

	_, errQuery := db.Exec("UPDATE streaming_movies SET streaming_date_end=? WHERE id = ?", updatedMonth, id)
	if errQuery != nil {
		ErrorMessage(c, http.StatusBadRequest, "Query Error")
	} else {
		SuccessMessage(c, http.StatusCreated, "Berhasil Update Streaming movie")
	}

}

func ShowStreamingList(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := ("SELECT DISTINCT(m.movie_name), m.thumbnail_path FROM streaming_movies sm JOIN movies m ON sm.movie_id = m.id")
	rows, err := db.Query(query)
	if err != nil {
		return
	}

	var streamingList m.StreamingList
	var streamingLists []m.StreamingList

	for rows.Next() {
		err = rows.Scan(&streamingList.MovieName, &streamingList.ThumbnailPath)
		if err != nil {
			ErrorMessage(c, http.StatusNotFound, "Data Not Found")
		}
		streamingLists = append(streamingLists, streamingList)
	}

	if len(streamingLists) != 0 {
		SuccessMessage(c, http.StatusOK, "Streaming List Data Found")
		c.IndentedJSON(http.StatusCreated, streamingLists)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func StreamingMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	idMovie := c.Query("ID_Stream")

	row := db.QueryRow("SELECT m.movie_name, m.synopsis, sm.movie_path FROM movies m JOIN streaming_movies sm ON m.id = sm.movie_id WHERE sm.id=?", idMovie)

	var movieStream m.StreamingMovie
	if err := row.Scan(&movieStream.MovieName, &movieStream.Synopsis, &movieStream.MoviePath); err != nil {
		ErrorMessage(c, http.StatusNotFound, "Scan not found")
	} else {
		SuccessMessage(c, http.StatusOK, "Streaming Movie Data Found")
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
			ErrorMessage(c, http.StatusNotFound, "Data Not Found")
		}
		theaters = append(theaters, theater)
	}

	if len(theaters) != 0 {
		c.IndentedJSON(http.StatusCreated, theaters)
		SuccessMessage(c, http.StatusOK, "Theaters Data Found")
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func ShowMovieDescription(c *gin.Context) {
	db := connect()
	defer db.Close()

	idMovie := c.Query("ID_Movie")

	//Get Movie Data
	rows := db.QueryRow("SELECT * FROM movies WHERE id = ?", idMovie)

	var movie m.Movie
	if err := rows.Scan(&movie.ID, &movie.Movie_name, &movie.Thumbnail_path, &movie.Synopsis, &movie.Last_premier, &movie.Streamable); err != nil {
		ErrorMessage(c, http.StatusNoContent, "Movie Not Found")
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
			ErrorMessage(c, http.StatusNotFound, "Data Not Found")
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

	idMovie := c.Query("ID_Movie")

	row := db.QueryRow("SELECT movie_name, thumbnail_path FROM movies WHERE id=?", idMovie)

	var movieTheatersInfo m.MovieTheaterInfo
	var allMovieTheatersInfo []m.MovieTheaterInfo
	var theatersCertainMovie m.TheatersCertainMovie

	if err1 := row.Scan(&theatersCertainMovie.MovieName, &theatersCertainMovie.ThumbnailPath); err1 != nil {
		ErrorMessage(c, http.StatusNotFound, "Data Not Found")
	} else {
		rows1, err2 := db.Query("SELECT DISTINCT(theaters.id), theaters.theater_name, theaters.price FROM movie_schedules JOIN studios ON movie_schedules.studio_id = studios.id JOIN theater_studio ON studios.id = theater_studio.studio_id JOIN theaters ON theater_studio.theater_id = theaters.id WHERE movie_schedules.movie_id =?", idMovie)
		if err2 != nil {
			ErrorMessage(c, http.StatusNoContent, "Query Error")
		}

		for rows1.Next() {
			var idTheater int //Id theater temporary variabel buat where di query selanjutnya
			err3 := rows1.Scan(&idTheater, &movieTheatersInfo.TheaterName, &movieTheatersInfo.Price)

			if err3 != nil {
				ErrorMessage(c, http.StatusNoContent, "Query Error")
			}

			rows2, err4 := db.Query("SELECT movie_schedules.playing_time FROM movie_schedules JOIN theater_studio ON movie_schedules.studio_id = theater_studio.studio_id WHERE movie_schedules.movie_id=? AND theater_studio.theater_id=?", idMovie, idTheater)
			if err4 != nil {
				ErrorMessage(c, http.StatusNoContent, "Query Error")
			}
			var timeArr []time.Time
			var time time.Time
			for rows2.Next() {
				err5 := rows2.Scan(&time)
				if err5 != nil {
					ErrorMessage(c, http.StatusNotFound, "Data Not Found")
				}
				timeArr = append(timeArr, time)
			}
			//Memasukkan array semua playing time movies_schedule dari 1 theater ke dalam variabel temporary
			movieTheatersInfo.DataPlayingTime = timeArr
			//Menggabungkan variabel temporary yang berisi informasi 1 theater suatu movie kedalam array temporary
			allMovieTheatersInfo = append(allMovieTheatersInfo, movieTheatersInfo)
		}
		//Menggabungkan variabel temporary yang berisi SEMUA informasi theater suatu movie kedalam kelas utama
		theatersCertainMovie.DataTheaters = allMovieTheatersInfo
		if len(allMovieTheatersInfo) != 0 {
			SuccessMessage(c, http.StatusOK, "Data Found Success")
			c.IndentedJSON(http.StatusOK, theatersCertainMovie)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	}
}

func AddMovie(c *gin.Context) {

	db := connect()
	defer db.Close()

	var movie m.Movie

	err := c.Bind(&movie)
	if err != nil {
		return
	}

	t := time.Now()
	t2 := t.AddDate(0, 3, 0)

	insert, err := db.Query("INSERT INTO movies(movie_name, thumbnail_path, synopsis, last_premier, streamable) VALUES(?,?,?,?,?)",
		movie.Movie_name, movie.Thumbnail_path, movie.Synopsis, t2, movie.Streamable,
	)

	if err != nil {
		ErrorMessage(c, http.StatusNoContent, "Query Error")
	} else {
		SuccessMessage(c, http.StatusOK, "Movie Added")
		c.IndentedJSON(http.StatusOK, insert)
	}

	defer insert.Close()
}

func UpdateMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	var movie m.Movie
	updateMonth, _ := strconv.Atoi(c.PostForm("Update Month"))
	err := c.Bind(&movie)
	if err != nil {
		return
	}

	var currentLastPremier time.Time
	row1 := db.QueryRow("SELECT last_premier FROM movies WHERE id=?",
		movie.ID,
	)
	if err := row1.Scan(&currentLastPremier); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	updatedMonth := currentLastPremier.AddDate(0, updateMonth, 0)

	update, err := db.Query("UPDATE movies SET movie_name=?, thumbnail_path=?, synopsis=?, last_premier=?, streamable=? WHERE id=?",
		movie.Movie_name, movie.Thumbnail_path, movie.Synopsis, updatedMonth, movie.Streamable, movie.ID,
	)

	if err != nil {
		ErrorMessage(c, http.StatusNoContent, "Query Error")
	} else {
		SuccessMessage(c, http.StatusOK, "Movie Updated")
		c.IndentedJSON(http.StatusOK, update)
	}
}
