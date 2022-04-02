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
	Password string    `json:"password"`
	Email    string    `json:"email"`
	UserType int       `json:"user_type"`
	Balance  int       `json:"balance"`
	LastSeen time.Time `json:"last_seen"`
}
type Users struct {
}

// movie related
type Movie struct {
}
type Movies struct {
}

type MovieSchedule struct {
	ID          int       `json:"id"`
	MovieID     int       `json:"movie_id"`
	StudioID    int       `json:"studio_id"`
	PlayingTime time.Time `json:"playing_time"`
}

type Studio struct {
}

type Seat struct {
}

// transaction

// response

// nanti dipertimbangkan dengan klmpk
