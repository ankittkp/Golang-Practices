package main

import "fmt"

func con(fN string, lN string) string {
	var fullName = fN +" " + lN
	return fullName
}
func main()  {
	fmt.Println("1. Initialise a string variable: ")
	var str = "Hi! this is Ankit"
	fmt.Println(str)

	fmt.Println("2. Pass two strings to any function, concatenate it and return to main function")
	var firstName = "Ankit"
	var lastName = "Kumar"

	fmt.Println(con(firstName, lastName))
	fmt.Println("3. Access the strings by indexes")
	fmt.Printf("3rd character of: %s is: %c\n", firstName, firstName[2])
}
