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
	shop.addBarber("Barber 1")

	// start a barber shop as a go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add customers
	i := 1
	go func() {
		for {
			//get a random number with average arrival rate

			randomMilliSecond := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				color.Red("Barber shop is closed for the day. No more clients will be accepted.")
				return
			case <-time.After(time.Duration(randomMilliSecond) * time.Millisecond):
				client := "Client " + string(i)
				i++
				shop.addClient(client)
			}
		}
	}()

	// block until barber shop is closed
	<-closed

	//time.Sleep(5 * time.Second)
}
