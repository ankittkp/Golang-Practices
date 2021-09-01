package main

import (
	"fmt"
	"time"
)

func message(msg string) {
	fmt.Printf("----- the message is: %v\n", msg)
}
func f(msg string) {
	fmt.Println(msg)
}
func main() {
	fmt.Println("1. Go Routines: A function that run concurrently, see it as a lightweight thread")
	fmt.Println("It helps in working on more than one task simultaneously, To use a go routine always use go before function call")
	fmt.Println("Lets see this function by calling it through goroutines")
	fmt.Println("Our two function calls are running asynchronously in separate goroutines now.")
	v:= "\"Hi there! The message is we are learning GO Routines today.\""
	go message(v + ". And this message is printed by using go routines")
	message(v)
	go secMessage("This is the second message passed by another go routines")
	secMessage("Second Message without go Routines")
	time.Sleep(time.Second)
	fmt.Println("When we run this program, we see the output of the blocking call first, then the output of the goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.")
	fmt.Println("Done")
}

func secMessage(s string) {
	fmt.Printf("+++++++++ %v\n", s)
}

/*
Goroutines are light on memory, a program can easily have hundreds or thousands of goroutines.
 */