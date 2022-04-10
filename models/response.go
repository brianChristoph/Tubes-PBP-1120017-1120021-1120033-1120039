package models

// response
type UsersResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

type MoviesResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

type TransactionsResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
