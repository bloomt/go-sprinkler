package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	pin1        = rpio.Pin(25) // Sprinkler valve 1
	pin2        = rpio.Pin(17) // Sprinkler valve 2
	pin3        = rpio.Pin(27) // Sprinkler valve 3
	pin4        = rpio.Pin(22) // Sprinkler valve 4
	breakFor    bool           //Break for loop if time has expired bool
	oneMin      = 10 * time.Second
	timerStatus = 0 //initialize timer status to 0 off
)

//Function to start timer - will be called by a goroutine
func startTimer() {

	var currentTime = time.Now()
	var timeOut = currentTime.Add(oneMin)
	timerStatus = 1 //change timer status to 1 showing that it is active

	for { // Get current time
		var t2 = time.Now()
		var duration = t2.Sub(currentTime) // Calculate running duration
		var timeCompare = currentTime.Add(duration)

		fmt.Println(duration)

		breakFor = timeCompare.After(timeOut)

		//Break for loop if time has expired
		if breakFor == true {
			break
		} else if timerStatus == 0 {
			break
		}
	}
	currentTime = time.Now()
}

func stopTimer() {
	timerStatus = 0
}

// Function to open and close relays - takes two perameters,
// First number of GPIO pin to controll
// Second High or low - High is open, Low is closed
func ControllValve(gpio int, highlow string) {

	if gpio == 1 && highlow == "h" {
		pin1.High()
		go startTimer()
		//fmt.Println("Pin 1 Open")

	} else if gpio == 2 && highlow == "h" {
		pin2.High()
		go startTimer()

	} else if gpio == 3 && highlow == "h" {
		pin3.High()
		go startTimer()
	} else if gpio == 4 && highlow == "h" {
		pin4.High()
		go startTimer()

	} else if gpio == 1 && highlow == "l" {
		pin1.Low()
		stopTimer() //Stops running timer
		//fmt.Println("Pin 1 Close")
	} else if gpio == 2 && highlow == "l" {
		pin2.Low()
		stopTimer() //Stops running timer
	} else if gpio == 3 && highlow == "l" {
		pin3.Low()
		stopTimer() //Stops running timer
	} else if gpio == 4 && highlow == "l" {
		pin4.Low()
		stopTimer() //Stops running timer
	}
}

//Set Timer
func Valve1Open1Minute() {
	// Set pin to output mode
	pin1.Low()
	time.Sleep(time.Second * 5)
	pin1.High()
	//pin1Status := pin1.Read()
	//fmt.Println(pin1Status)

	//fmt.Fprintln(w, "Welcome!")

	// Toggle pin 20 times
	// pin1.Toggle()
}
