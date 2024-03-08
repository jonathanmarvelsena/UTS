package main

import (
	"UTS/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/no1", controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/no2", controller.GetRoomDetails).Methods("GET")
	router.HandleFunc("/no3", controller.EnterRoom).Methods("POST")
	router.HandleFunc("/no4", controller.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
