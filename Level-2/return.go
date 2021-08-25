package main

import "fmt"

func main(){
	fmt.Println("Go functions can return one or more values. Classical programming languages often only have zero or at most one return value.")
	fmt.Println("1. Lets return sum and multiply together")
	fmt.Println(SumAndMul(10,15))
	fmt.Println("2. Lets try for combination of strings and numbers")
	fmt.Println(NumAndStr("My lucky number is ", 0))
}

func NumAndStr(s string, x int) (string, int) {
	return s, 74
}

func SumAndMul(y int, x int) (int, int) {
	return 10+15, 10*15
}
