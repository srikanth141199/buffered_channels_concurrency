package main

import "time"

type BarberShop struct {
	SeatingCapacity int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	open            bool
}
