package controllers

import (
	"net/http"
	"time"

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
func UpdateStreamingMovie(c *gin.Context) {
	db := connect()
	defer db.Close()
}

func ShowStreamingList(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := ("SELECT  FROM streaming_movies sm")

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	var streamingList m.StreamingList
	var streamingLists []m.StreamingList

	for rows.Next() {
		err = rows.Scan(&streamingList.ID, &streamingList.MovieId, &streamingList.MoviePath)
		if err != nil {
			panic(err.Error())
		}
		streamingLists = append(streamingLists, streamingList)
	}

	if len(streamingLists) != 0 {
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

	idMovie := c.Query("ID_Movie")

	row := db.QueryRow("SELECT movie_name, thumbnail_path FROM movies WHERE id=?", idMovie)

	var movieTheatersInfo m.MovieTheaterInfo
	var allMovieTheatersInfo []m.MovieTheaterInfo
	var theatersCertainMovie m.TheatersCertainMovie

	if err := row.Scan(&theatersCertainMovie.MovieName, &theatersCertainMovie.ThumbnailPath); err != nil {
		panic(err.Error())
	} else {
		rows1, err := db.Query("SELECT DISTINCT(theaters.id), theaters.theater_name, theaters.price FROM movie_schedules JOIN studios ON movie_schedules.studio_id = studios.id JOIN theater_studio ON studios.id = theater_studio.studio_id JOIN theaters ON theater_studio.theater_id = theaters.id WHERE movie_schedules.movie_id =?", idMovie)
		if err != nil {
			panic(err.Error())
		}

		for rows1.Next() {
			var idTheater int //Id theater temporary variabel buat where di query selanjutnya
			err := rows1.Scan(&idTheater, &movieTheatersInfo.TheaterName, &movieTheatersInfo.Price)

			if err != nil {
				panic(err.Error())
			}

			rows2, err := db.Query("SELECT movie_schedules.playing_time FROM movie_schedules JOIN studios ON movie_schedules.studio_id = studios.id JOIN theater_studio ON studios.id = theater_studio.studio_id WHERE movie_schedules.movie_id=? AND theater_studio.theater_id=?", idMovie, idTheater)
			if err != nil {
				panic(err.Error())
			}
			var timeArr []time.Time
			var time time.Time
			for rows2.Next() {
				err := rows2.Scan(&time)
				if err != nil {
					panic(err.Error())
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
		c.IndentedJSON(http.StatusOK, theatersCertainMovie)
	}
}

func ChangePrice(c *gin.Context) {
	db := connect()
	defer db.Close()
}
