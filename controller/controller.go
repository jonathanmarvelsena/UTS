package controller

import (
	m "UTS/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid form data")
		return
	}

	gameIDStr := r.URL.Query().Get("id")
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid game ID")
		return
	}

	query := "SELECT id, room_name FROM rooms WHERE id_game = ?"

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
		err := rows.Scan(&room.ID, &room.Room_name)
		if err != nil {
			log.Println(err)
			SendErrorResponse(w, 500, "Internal Server Error")
			return
		}

		rooms = append(rooms, room)
	}

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

func GetRoomDetails(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		SendErrorResponse(w, 500, "error parsing form data")
		return
	}

	roomIDStr := r.Form.Get("id_room")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid room ID")
		return
	}

	var room m.Room
	err = db.QueryRow("SELECT id, room_name FROM rooms WHERE id = ?", roomID).
		Scan(&room.ID, &room.Room_name)
	if err == sql.ErrNoRows {
		SendErrorResponse(w, 404, "Room not found")
		return
	} else if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	rows, err := db.Query("SELECT p.id, p.id_account, a.username FROM participants p JOIN accounts a ON p.id_account = a.id WHERE p.id_room = ?", roomID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}
	defer rows.Close()

	var participants []m.ParticipantDetail

	for rows.Next() {
		var participant m.ParticipantDetail
		err := rows.Scan(&participant.ID, &participant.ID_account, &participant.Username)
		if err != nil {
			log.Println(err)
			SendErrorResponse(w, 500, "Internal Server Error")
			return
		}

		participants = append(participants, participant)
	}

	response := m.RoomDetailResponse{
		Status:       200,
		Message:      "Success",
		Data:         room,
		Participants: participants,
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

func EnterRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		SendErrorResponse(w, 500, "error parsing form data")
		return
	}

	roomIDStr := r.Form.Get("id_room")
	accountIDStr := r.Form.Get("id_account")

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid room ID")
		return
	}

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid account ID")
		return
	}

	var room m.Room
	err = db.QueryRow("SELECT r.id, r.room_name, g.id, g.name, g.max_player FROM rooms r JOIN games g ON r.id_game = g.id WHERE r.id = ?", roomID).
		Scan(&room.ID, &room.Room_name, &room.ID_game.ID, &room.ID_game.Name, &room.ID_game.Max_player)
	if err == sql.ErrNoRows {
		SendErrorResponse(w, 404, "Room not found")
		return
	} else if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	var participantCount int
	err = db.QueryRow("SELECT COUNT(id) FROM participants WHERE id_room = ?", roomID).Scan(&participantCount)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	if participantCount >= room.ID_game.Max_player {
		SendErrorResponse(w, 400, "Room is at its maximum limit")
		return
	}

	_, err = db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?, ?)", roomID, accountID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	SendSuccessResponse(w, 200, "Entered the room successfully")
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		SendErrorResponse(w, 500, "error parsing form data")
		return
	}

	roomIDStr := r.Form.Get("id_room")
	accountIDStr := r.Form.Get("id_account")

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid room ID")
		return
	}

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid account ID")
		return
	}

	var participantID int
	err = db.QueryRow("SELECT id FROM participants WHERE id_room = ? AND id_account = ?", roomID, accountID).
		Scan(&participantID)
	if err == sql.ErrNoRows {
		SendErrorResponse(w, 404, "Participant not found in the room")
		return
	} else if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	_, err = db.Exec("DELETE FROM participants WHERE id = ?", participantID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	SendSuccessResponse(w, 200, "Left the room successfully")
}
