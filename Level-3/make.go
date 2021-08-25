package main

import "fmt"

func main() {
	fmt.Println("Make in go")
	fmt.Println("1. Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.")
	a := make([]int, 5)
	b := make([]int, 0, 5) //third argument denotes capacity
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("2. New v/s Make in Go")
	fmt.Println("You need make() to create channels and maps (and slices, but those can be created from arrays too)." +
		"There's no alternative way to make those, so you can't remove make() from your lexicon.")
	fmt.Println("The builtin new(T) function allocates “zeroed” storage for a new item of type T and returns its address," +
		" a value of type *T. In Go terminology, it returns a pointer to a newly allocated zero value of type T")
}
