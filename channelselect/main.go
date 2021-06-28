package main

import "fmt"

func main() {
	orderChannel := make(chan string)
	paymentChannel := make(chan string)
	locationChannel := make(chan string)
	//send data
	go func() {
		orderChannel <- "Order Placed"
	}()
	go func() {
		paymentChannel <- "Payment Success"
	}()
	go func() {
		locationChannel <- "Location Found"
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

}
