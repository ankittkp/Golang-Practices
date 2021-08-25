package main

import "fmt"

type Vertex struct { //structure in next Level
	X int
	Y int
}

func main(){
	fmt.Println("Not all programming languages support pointers, typically they are found in low level languages like C or C++. Go also supports them.")
	fmt.Println("Every variable is is stored in the memory and has a unique address. This can be accessed with the ampersand (&) operator.")
	var x int = 10
	fmt.Printf("Value and Address of x is %d, %d\n",x,&x)
	fmt.Println("The variables are stored in memory and a pointer is a variable whose address is the " +
		"direct memory address of another variable")
	fmt.Println("2. Pointers examples")
	b := 200
	var a *int = &b
	fmt.Printf("Deferencing value of a : %d\n",*a)
	fmt.Println("3. Pointers to structs")
	v := Vertex{1, 2}
	p := &v //
	p.X = 1e9
	fmt.Println(v)
}
/*
1. Do not pass a pointer to an array as an argument to a function. Use slice instead.
2. You can Return pointer from a function.
3. You can pass pointer to a function.
4. You can Create pointers using the new function
5. Go does not support pointer arithmetic

*/
