package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//created channel
	orderChannel := make(chan string)
	paymentChannel := make(chan string)
	locationChannel := make(chan string)
	//send data
	go func() {
		for {
			time.Sleep(5 * time.Second)
			orderChannel <- "Order Placed" + "," + strconv.Itoa(int(rand.Int31n(1000)))
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			paymentChannel <- "Payment Success" + "," + strconv.Itoa(int(rand.Int31n(1000)))
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			locationChannel <- "Location Found" + "," + strconv.Itoa(int(rand.Int31n(1000)))
		}
	}()

	//receiver go routine
	go func() {
		for {
			select {
			case msg1 := <-orderChannel:
				fmt.Println(msg1)
			case msg2 := <-paymentChannel:
				fmt.Println(msg2)
			case msg3 := <-locationChannel:
				fmt.Println(msg3)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
