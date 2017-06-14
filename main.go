package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin11 = rpio.Pin(25)
	pin22 = rpio.Pin(17)
	pin33 = rpio.Pin(27)
	pin44 = rpio.Pin(22)
)

func exitCloseValves() {
	pin11.Low()
	pin22.Low()
	pin33.Low()
	pin44.Low()
}

func main() {
	defer os.Exit(0)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		exitCloseValves()
		os.Exit(1)
	}

	//Run code before exit
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exitCloseValves()
		os.Exit(1)
	}()

	// Unmap gpio memory when done
	defer rpio.Close()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
	runtime.Goexit()
}
