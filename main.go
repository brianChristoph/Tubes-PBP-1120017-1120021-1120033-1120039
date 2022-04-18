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

var userTypeForAdmin = []string{c.LoadEnv("ADMIN")}
var userTypeForMemberOnly = []string{c.LoadEnv("MEMBER")}
var userTypeForVIPOnly = []string{c.LoadEnv("VIP")}
var userTypeForVIPMember = []string{c.LoadEnv("VIP"), c.LoadEnv("MEMBER")}

func main() {
	router := SetupRouter()

	// Background Function
	go api_tools.RunBackgroundFunc()

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

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	// ADMIN
	router.GET("/admin/users", c.Authentication(c.GetAllUser, userTypeForAdmin))               //Show All User
	router.DELETE("/admin/user", c.Authentication(c.DeleteUser, userTypeForAdmin))             //Delete User
	router.POST("/admin/movies", c.Authentication(c.AddMovie, userTypeForAdmin))               //Add Movie
	router.PUT("/admin/movies", c.Authentication(c.UpdateMovie, userTypeForAdmin))             //Update Movie
	router.PUT("/movie_streaming", c.Authentication(c.UpdateStreamingMovie, userTypeForAdmin)) //Update Streaming

	//USER
	router.POST("/user/login", c.Login)                                                       //Login
	router.POST("/user/register", c.Register)                                                 //Register
	router.PUT("/user/update", c.Authentication(c.UpdateUser, userTypeForVIPMember))          //Update User
	router.POST("/user/logout", c.Logout)                                                     //Logout
	router.GET("/user/profile", c.Authentication(c.UserProfile, userTypeForVIPMember))        //User Profile
	router.PUT("/user/transaction/buyVIP", c.Authentication(c.BuyVIP, userTypeForMemberOnly)) //Buy VIP
	router.PUT("/user/topUp", c.Authentication(c.TopUp, userTypeForVIPMember))                //Top Up

	//STREAMING
	router.GET("/streaming_movies/list", c.Authentication(c.ShowStreamingList, userTypeForVIPOnly)) //Show Streaming List
	router.GET("/streaming_movies/stream", c.Authentication(c.StreamingMovie, userTypeForVIPOnly))  //Streaming Movie

	//MOVIES
	router.GET("/theaters/list", c.TheaterList)                     //Theater List
	router.GET("/movies/description", c.ShowMovieDescription)       //Show Movie Description
	router.GET("/movies/list", c.ShowMovieList)                     //Show Movie List
	router.GET("/theaters/available", c.ShowTheaterForCertainMovie) //Show Available Theater for Certain Movie

	//TRANSACTION
	router.PUT("/theaters/studios/seats", c.Authentication(c.BookingSeats, userTypeForVIPMember)) //Booking Seats

	return router
}
