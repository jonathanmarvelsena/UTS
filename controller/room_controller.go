package controller

import (
	"log"
	"net/http"

	m "UTS/model"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	gameIDStr := r.URL.Query().Get("id_game")
	gameID, err := strconv.Atoi(gameIDStr)


	query := "SELECT * FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	SendSuccessResponse(w, 200, "succes")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	userID := mux.Vars(r)["id"]

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userID).Scan(&count)
	if err != nil {
		SendErrorResponse(w, 500, "error preparing SQL statement")
		return
	}
	if count == 0 {
		SendErrorResponse(w, 404, "user not found")
		return
	}