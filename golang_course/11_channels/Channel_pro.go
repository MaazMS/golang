package main

import "fmt"

func channelPro() {

	// creating channel by using var
	var ourchannel chan int

	fmt.Println("value of channel", ourchannel)
	fmt.Printf("channel %T", ourchannel)

	// creating channel by using make()

	mychannel := make(chan int)
	fmt.Println("value of channel", mychannel)
	fmt.Printf("channel %T", mychannel)
}
func manin() {

	channelPro()
}
// Today I am available after 6.00 AM