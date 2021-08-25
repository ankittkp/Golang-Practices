package main

import "fmt"

func main(){
	fmt.Println("1. Simple way to define functions")
	//Go will run all statements in the function, then do the function call specified by defer after.

	Who()
	Are()
	fmt.Println("2. Pass arguments")
	You("hello")
	fmt.Println("3. Return Sum")
	fmt.Println(sum(10, 15))
}
func sum(x int, y int) int {
	return x + y
}
func You(str string) {
	fmt.Println(str)
}

func Are() {
	fmt.Println("are ?????")
}

func Who() {
	fmt.Println("Who ?????")
}
