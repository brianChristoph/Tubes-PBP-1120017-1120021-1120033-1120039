package main

import (
	"log"
	"net/http"

	api_tools "github.com/Tubes-PBP/api-tools"
	"github.com/rs/cors"

	// controllers
	c "github.com/Tubes-PBP/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	// Background Function
	go api_tools.RunBackgroundFunc()

	//ADMIN
	router.GET("/admin/users", c.GetAllUser)               //Show ALl User
	router.DELETE("/admin/user", c.DeleteUser)             //Delete User
	router.POST("/admin/movies", c.Register)               //Add Movie
	router.PUT("/admin/movies", c.UpdateMovie)             //Update Movie
	router.GET("/movie_streaming", c.UpdateStreamingMovie) //Update Streaming

	//USER
	router.POST("/user/login", c.Login)              //Login
	router.POST("/user/register", c.Register)        //Register
	router.PUT("/user/update", c.UpdateUser)         //Update User
	router.POST("/user/logout", c.Logout)            //Logout
	router.GET("/user/profile", c.UserProfile)       //User Profile
	router.PUT("/user/transaction/buyVIP", c.BuyVIP) //Buy VIP

	//STREAMING
	router.GET("/streaming_movies/list", c.ShowStreamingList) //Show Streaming List
	router.GET("/streaming_movies/stream", c.StreamingMovie)  //Streaming Movie

	//MOVIES
	router.GET("/theaters/list", c.TheaterList)                     //Theater List
	router.GET("/movies/description", c.ShowMovieDescription)       //Show Movie Description
	router.GET("/movies/list", c.ShowMovieList)                     //Show Movie List
	router.GET("/theaters/available", c.ShowTheaterForCertainMovie) //Show Available Theater for Certain Movie

	//TRANSACTION
	router.POST("/transaction/buyTicket", c.TransactionBuyTicket) //Transaction Buy Ticket
	router.PUT("/theaters/studios/seats", c.BookingSeats)         //Booking Seats

	// CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	router.Run("localhost:" + c.LoadEnv("PORT"))
	log.Fatal(http.ListenAndServe(":8080", handler))
}
