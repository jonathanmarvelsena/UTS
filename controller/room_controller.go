package controller

import (
	m "UTS/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	gameIDStr := r.URL.Query().Get("id_game")
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid game ID")
		return
	}

	query := "SELECT * FROM rooms WHERE id_game = ?"

	rows, err := db.Query(query, gameID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}
	defer rows.Close()

	var rooms []m.Room

	for rows.Next() {
		var room m.Room
		err := rows.Scan(&room.id, &room.room_name, &room.id_game)
		if err != nil {
			log.Println(err)
			SendErrorResponse(w, 500, "Internal Server Error")
			return
		}

		// Append the room to the list
		rooms = append(rooms, room)
	}

	// Create the response
	response := m.RoomsResponse{
		Status:  200,
		Message: "Success",
		Data:    rooms,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
