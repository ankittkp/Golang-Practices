package main

import "fmt"

func main(){
	fmt.Println("Defer is a special statement in Go. The defer statement schedules a function to be called after the current function has completed.")
	//Go will run all statements in the function, then do the function call specified by defer after.

	who()
	defer are()
	you()

	fmt.Println("Defer can also be called on simple function calls from packages.")
}

func you() {
	fmt.Println("You????")
}

func are() {
	fmt.Println("are ?????")
}

func who() {
	fmt.Println("Who ?????")
}