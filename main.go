package main

import (
	api_tools "github.com/Tubes-PBP/api-tools"
	// controllers
	c "github.com/Tubes-PBP/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	go api_tools.RunBackgroundFunc()

	//ADMIN
	router.GET("/user", c.GetAllUser)//View ALl User
	router.DELETE("/user", c.DeleteUser)//Delete User
	router.POST("/movies", c.Register)//Add Movie
	router.PUT("/movies", c.Register)//Update Movie
	router.GET("/streaming", c.UpdateStreaming)//Update Streaming

	//USER
	router.GET("/user/login", c.Login)//Login
	router.POST("/user", c.Register)//Register
	router.PUT("/user", c.UpdateUser)//Update User
	router.GET("/user/logout", c.Logout)//Logout
	router.GET("/user/profile", c.UserProfile)//User Profile
	router.GET("/user/transaction/buyVIP", c.BuyVIP)//Buy VIP

	//STREAMING
	router.GET("/streaming_movies/list", c.ShowStreamingList) //Show Streaming List

	// Pakai query params
	router.GET("/streaming_movies/stream", c.StreamingMovie) //Streaming Movie

	//MOVIES
	router.GET("/theaters/list", c.TheaterList)                            //Theater List
	router.GET("/movies/description", c.ViewMovieDescription)              //View Movie Description
	router.GET("/movies/list", c.ShowMovieList)                            //Show Movie List
	router.GET("/theaters/available", c.ShowTheaterForCertainMovie)        //Show Available Theater for Certain Movie
	router.GET("/movies/changePrice", c.ChangePrice)                       //Change Price
	router.GET("/movies/changeMovieDescription", c.UpdateMovieDescription) //Update movie description

	//TRANSACTION
	router.GET("/transaction/buyTicket", c.TransactionBuyTicket) //Transaction Buy Ticket
	router.GET("/theaters/studios/seats", c.BookingSeat)         //Pemesanan Kursi

	router.Run("localhost:8080")
}
