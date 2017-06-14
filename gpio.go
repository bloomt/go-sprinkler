package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	pin1     = rpio.Pin(25) // Sprinkler valve 1
	pin2     = rpio.Pin(17) // Sprinkler valve 2
	pin3     = rpio.Pin(27) // Sprinkler valve 3
	pin4     = rpio.Pin(22) // Sprinkler valve 4
	breakFor bool
	oneMin   = 10 * time.Second
)


// Function to open and close relays - takes two perameters,
// First number of GPIO pin to controll
// Second High or low - High is open, Low is closed
func ControllValve(gpio int, highlow string) {

	//Create a number channel
	n := make(chan int)

	go func() {
	var currentTime = time.Now()
	var timeOut = currentTime.Add(oneMin)

	for { // Get current time
		var t2 = time.Now()
		var duration = t2.Sub(currentTime) // Calculate running duration

		var timeCompare = currentTime.Add(duration)
		//fmt.Println(currentTime)
		//fmt.Println(timeCompare)
		//fmt.Println(duration)

		breakFor = timeCompare.After(timeOut)
		//fmt.Println(breakFor)

		//Read number channel and dump in variable
		n := <-n
		//fmt.Println(n)

		//Break for loop if time has expired
		if breakFor == true {
			//fmt.Println(duration) //Print end duration in console
			break
		} else if n == 0 {
			//fmt.Println(duration) //Print exit duration in console
			break
		}
		//fmt.Println(breakFor)

		// Get duration.
		//fmt.Println("Duration", duration)
	}
	}()

	if gpio == 1 && highlow == "h" {
		pin1.High()
		//fmt.Println("Pin 1 Open")
		n <- 1 //Set n to 1 to signal timer
	} else if gpio == 2 && highlow == "h" {
		pin2.High()
		n <- 1 //Set n to 1 to signal timer
	} else if gpio == 3 && highlow == "h" {
		pin3.High()
		n <- 1 //Set n to 1 to signal timer
	} else if gpio == 4 && highlow == "h" {
		pin4.High()
		n <- 1 //Set n to 1 to signal timer
	} else if gpio == 1 && highlow == "l" {
		pin1.Low()
		n <- 0 //Set n to 0 to signal timer off
		//fmt.Println("Pin 1 Close")
	} else if gpio == 2 && highlow == "l" {
		pin2.Low()
		n <- 0 //Set n to 0 to signal timer off
	} else if gpio == 3 && highlow == "l" {
		pin3.Low()
		n <- 0 //Set n to 0 to signal timer off
	} else if gpio == 4 && highlow == "l" {
		pin4.Low()
		n <- 0 //Set n to 0 to signal timer off
	}
	//fmt.Println(pin)

	// Set pin to output mode
	//in.High()
	// pinStatus := gpio.Read()
	// fmt.Println(pinStatus)

	// Toggle pin 20 times
	// pin1.Toggle()
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
