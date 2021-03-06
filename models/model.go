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

type UpdateRegister struct {
	Name            string `form:"name" json:"name"`
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
}

type Movie struct {
	ID             int       `json:"id"`
	Movie_name     string    `json:"movie_name"`
	Thumbnail_path string    `json:"thumbnail_path"`
	Synopsis       string    `json:"synopsis"`
	Last_premier   time.Time `json:"last_premier"`
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

type BookingSeats struct {
	Theater_ID int `json:"theater_id"`
	Studio_ID  int `json:"studio_id"`
	Seat_ID    int `json:"seat_id"`
	Ms_ID      int `json:"ms_id"`
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

type TheatersCertainMovie struct {
	MovieName     string             `json:"movie_name"`
	ThumbnailPath string             `json:"thumbnail_path"`
	DataTheaters  []MovieTheaterInfo `json:"data_theaters"`
}

type MovieTheaterInfo struct {
	TheaterName     string      `json:"theater_name"`
	Price           int         `json:"price"`
	DataPlayingTime []time.Time `json:"data_playing_time"`
}

type StreamingList struct {
	MovieName     string `json:"movie_name"`
	ThumbnailPath string `json:"thumbnail_path"`
}

type UpdateStreamingMovie struct {
	StreamingDateEnd time.Time `json:"streaming_date_end"`
}
