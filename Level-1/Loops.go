package main

import "fmt"

func main()  {
	fmt.Println("The Definition of for loop in go")
	fmt.Println("for initialization statement; conditional statement; modified statement {}")
	fmt.Println("1. Simple For Loop: Store 2 table to an array")
	arr := new([11]int)
	for i:= 1;i<10;i++ {
		arr[i] = 2*i
	}
	fmt.Println("There's nothing like while keyword in go, A for loop has a predefined amount of iterations. A while loop doesn’t necessarily.")
	fmt.Println("2. Now using for(while) loop")
	j := 1
	max := 11
	for j < max {
		arr[j] = 2*j
		j += 1
	}
	fmt.Println("3. Ranges in go: Range iterates over elements. That can be elements of an array, elements of a dictionary or other data structures.")
	fmt.Println("When using range, you can name the current element and current index:")
	fmt.Println(" you are free to ignore the index: for _, num := range nums {")
	nums := []int{1,2,3,4,5}
	for index, num := range nums {
		fmt.Print(index, ":",num,"\n")
	}
}
/*
	What is the purpose of range ?
		To iterate over all the elements of a data structure, array, slice and others.
	What the difference between the line for index, element := range a and the line for _, element := range a ?
		index holds the position of the element in the data structure. Sometimes you may not need it, then you can use underscore.
		underscore is blank identifier check input.go file.
 */