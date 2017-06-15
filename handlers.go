package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// "github.com/stianeikeland/go-rpio"

func Index(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./public")).ServeHTTP(w, r)
}

// Start of custom code
// GPIO return Struct
type GPIOReturn struct {
	Time time.Time `json:"title"`
}

type GPIOReturns []GPIOReturn

//Open Relay 1
func Open1(w http.ResponseWriter, r *http.Request) {
	ControllValve(1, "h")

	//First attempt at returning time for clock
	gpioreturns := GPIOReturns{
		GPIOReturn{Time: duration},
	}

	if err := json.NewEncoder(w).Encode(gpioreturns); err != nil {
		panic(err)
	}
	//End attempt at returning time for clock
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
