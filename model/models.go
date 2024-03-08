package models

type Account struct {
	id       int    `json:"id"`
	username string `json:"username"`
}

type AccountResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}
type AccountsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

type Game struct {
	id         int    `json:"id"`
	name       string `json:"name"`
	max_player int    `json:"max_player"`
}

type Room struct {
	id        int    `json:"id"`
	room_name string `json:"room_name"`
	id_game   int    `json:"id_game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type Participant struct {
	id         int `json:"id"`
	id_room    int `json:"id_room"`
	id_account int `json:"id_account"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseGorm struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	ID      int    `json:"id"`
}
