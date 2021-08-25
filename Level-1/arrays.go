package main

import "fmt"

func main()  {
	fmt.Println("An array in the Go language is a type of value (not a pointer to the first element in C/C++), so it can be created by new():")
	fmt.Println("If each element is an integer value, all elements are automatically initialized to the default value of 0 when the array is declared.")
	fmt.Println("1. Create a new array of size 10")
	var arr = new([10]int)
	fmt.Println("Check the 0th and 1st index value")
	fmt.Printf("Oth Value: %d, 1st Value: %d", arr[0],arr[1])
	fmt.Println("2. Several assignment modes in array")
	var arrNum = [6]int{10,20,30,40,50,60}
	var arrNoSize = [...]int{1,2,3}
	arrKey := [3]string{1: "ankit", 2: "kumar"}

	fmt.Println("Different modes are :")
	fmt.Printf("arrNum are : %v\n",arrNum)
	fmt.Printf("arrNoSize are : %v\n",arrNoSize)
	fmt.Printf("arrKey are : %v\n",arrKey)

}
