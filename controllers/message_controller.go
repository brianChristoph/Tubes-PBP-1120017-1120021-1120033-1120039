package controllers

import (
	"encoding/json"
	"net/http"
)

func SuccessMessage(status int, message string, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("...")
}

func ErrorMessage(status int, message string, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("...")
}
