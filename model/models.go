package models

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
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
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	ID_game   Game   `json:"id_game"`
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

type RoomDetailResponse struct {
	Status       int                 `json:"status"`
	Message      string              `json:"message"`
	Data         Room                `json:"data"`
	Participants []ParticipantDetail `json:"participants"`
}

type ParticipantDetail struct {
	ID         int    `json:"id"`
	ID_account int    `json:"id_account"`
	Username   string `json:"username"`
}

type Participant struct {
	ID         int `json:"id"`
	ID_room    int `json:"id_room"`
	ID_account int `json:"id_account"`
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
