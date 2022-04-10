package models

import "time"

/*
	user : userType
	0 = Guest
	1 = Member
	2 = Admin
*/
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `form:"password" json:"password"`
	Email    string    `form:"email" json:"email"`
	UserType string    `json:"user_type"`
	Balance  int       `json:"balance"`
	LastSeen time.Time `json:"last_seen"`
}

type Movie struct {
	ID             int       `json:"id"`
	Movie_name     string    `json:"movie name"`
	Thumbnail_path string    `json:"thumbnail_path"`
	Synopsis       string    `json:"synopsis"`
	Last_premier   time.Time `json:"last premier"`
	Streamable     int       `json:"streamable"` // 1= can be stream, 0 = cant be stream
}

type Movie_Schedule struct {
	ID          int       `json:"id"`
	MovieID     int       `json:"movie_id"`
	StudioID    int       `json:"studio_id"`
	PlayingTime time.Time `json:"playing_time"`
}

type Studio struct {
	ID          int    `json:"id"`
	Studio_name string `json:"studio_id"`
	Theatre_id  int    `json:"theatre_id"`
	Status      int    `json:"status"`
	Quantity    int    `json:"quantity"` //seat quantity in studio
}

type Seat struct {
	ID         int    `json:"id"`
	Seats_name string `json:"seats_name"`
}

type Transaction struct {
	ID               int       `json:"id"`
	Person_id        int       `json:"person_id"`
	Transaction_date time.Time `json:"transaction_date"`
}

type Detail_Transaction struct {
	ID             int `json:"id"`
	Transaction_id int `json:"transaction_id"`
	Seat_id        int `json:"seat_id"`
	Ms_id          int `json:"ms_id"`
}

type Theater struct {
	ID          int    `json:"id"`
	TheaterName string `json:"theatre_name"`
	LocationID  int    `json:"location_id"`
	Price       int    `json:"price"`
}

type StreamingMovie struct {
	MovieName string `json:"movie_name"`
	Synopsis  string `json:"synopsis"`
	MoviePath string `json:"movie_path"`
}
