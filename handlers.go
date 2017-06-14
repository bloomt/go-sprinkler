package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/stianeikeland/go-rpio"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./public")).ServeHTTP(w, r)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

//Start of custom code
// GPIO return Struct
type GPIOReturn struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

//Open Relay 1
func Open1(w http.ResponseWriter, r *http.Request) {
	ControllValve(1, "h")
}

//Close Relay 1
func Close1(w http.ResponseWriter, r *http.Request) {
	ControllValve(1, "l")
}

//Open Relay 2
func Open2(w http.ResponseWriter, r *http.Request) {
	ControllValve(2, "h")
}

//Close Relay 2
func Close2(w http.ResponseWriter, r *http.Request) {
	ControllValve(2, "l")
}

//Open Relay 3
func Open3(w http.ResponseWriter, r *http.Request) {
	ControllValve(3, "h")
}

//Close Relay 3
func Close3(w http.ResponseWriter, r *http.Request) {
	ControllValve(3, "l")
}

//Open Relay 4
func Open4(w http.ResponseWriter, r *http.Request) {
	ControllValve(4, "h")
}

//Close Relay 4
func Close4(w http.ResponseWriter, r *http.Request) {
	ControllValve(4, "l")
}

//Close Relay 4
func Open1Timer(w http.ResponseWriter, r *http.Request) {
	Valve1Open1Minute()
}
