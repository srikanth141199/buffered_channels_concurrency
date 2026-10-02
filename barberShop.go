package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	SeatingCapacity int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to waiting room to check for clients.", barber)
		for {
			//if there are no clients barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("%s sees there are no clients so barber takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes up and starts cutting hair of %s.", barber, client)
					isSleeping = false
				}
				//cut hair
				shop.cutHair(barber, client)
			} else {
				//shop is closed so send barber home and close go routine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber string, client string) {
	// Implementation for cutting hair
	color.Green("%s is cutting hair of %s.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done cutting hair of %s.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Red("Barber shop is closing for the day.")
	shop.open = false
	close(shop.ClientsChan)

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	color.Red("-----------------------------------")
	color.Red("Barber shop is closed for the day.")
}

func (shop *BarberShop) addClient(client string) {
	//print out msg
	color.Green("***A %s has arrived at the barber shop.***", client)

	if shop.open {
		shop.ClientsChan <- client
		select {
		case shop.ClientsChan <- client:
			color.Blue("%s is waiting for a haircut.", client)
		default:
			color.Red("%s leaves the barber shop because there are no available seats.", client)
		}
	} else {
		color.Red("Barber shop is closed for the day. No more clients will be accepted.")
	}
}
