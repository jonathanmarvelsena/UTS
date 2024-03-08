package controller

import (
	"encoding/json"
	"net/http"

	m "UTS/model"
)

func SendErrorResponse(w http.ResponseWriter, kode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = kode
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(w http.ResponseWriter, kode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.SuccessResponse
	response.Status = kode
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
