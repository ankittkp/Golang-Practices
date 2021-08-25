package main

import "fmt"

func main() {
	fmt.Println("Variadic functions can be called with any number of arguments.")
	fmt.Println("A function is said to be variadic, if the number of arguments are not explictly defined.")
	fmt.Println("1. Adding multiple numbers")
	add(2,3,4)
	fmt.Println("2. Printing multiple name of strings")
	a := []string{"country", "state","charley"}
	show(a...)
}

func show(a ...string) {
	for _, a := range a {
		fmt.Println(a)
	}
}
func add(numbers ...int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println(total)
}