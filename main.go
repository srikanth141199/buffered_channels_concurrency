package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

//variables

var seatingCapacity int = 10
var arrivalRate int = 100 // milliseconds
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome msg
	color.Yellow("The sleeping Barber problem.")
	color.Yellow("------------------------------")

	//create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create some data structure that represents the barber shop
	shop := BarberShop{
		SeatingCapacity: seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		open:            true,
	}

	color.Green("Barber shop is open for business!")

	// add barbers

	// start a barber shop

	// start a barber shop as a go routine

	// add customers

	// block until barber shop is closed

}
