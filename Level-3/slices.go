package main

import "fmt"

func main(){
	fmt.Println("Slices in go: Slices are subsets,  slice can be a subset of an array, list or string. You can take multiple slices out of a string, each as a new variable.")
	fmt.Println("A slice is never longer than then the original variable. This makes sense, if you take a slice of a pizza you don’t suddenly have two pizzas. In programming it’s similar.")
	fmt.Println("1. Lets slice out of a list of numbers")
	nums := []int{1,2,3,4,5,6}

	//take slice
	s := nums[0:4]

	fmt.Printf("Original list is : %v\nand After Slicing new List is: %v\n", nums, s)
	fmt.Println("2. Can you take slice of slices ? Yes, Why not")
	sos := s[1:3]

	fmt.Printf("Result : %v", sos)

}
