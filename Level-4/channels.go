package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Channels are a way for goroutines to communicate with each other. ")
	fmt.Println("Because goroutines run concurrently, they can’t simply pass data from one goroutine to another. Channels are needed.")
	fmt.Println("So How Channels work?")


	fmt.Println("1. Send Message to Channels: c <- message, message gets sends to the channel")
	fmt.Println("2. Receive the message and store in variable , msg:= <-c")
	var c chan string = make(chan string)

	go f1(c) //these go routines communicate via a channel
	go f2(c)
	time.Sleep(time.Second)

}
func f1(c chan string) {
	c <- "Insert message to channel"
}
func f2(c chan string)  {
	msg := <- c
	fmt.Printf("%v\n", msg)
}
