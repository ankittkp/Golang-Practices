package main

import "fmt"

var y = 200
func main()  {
	fmt.Println("Scope is where a variable can be used, often inside a function called a local variable")
	fmt.Println("Initialise a local variable")
	x := 100
	fmt.Printf("Value of x is : %v ",x)
	fmt.Println("If you move the variable x outside the function, It will be global variable , Create a new one then ")
	fmt.Println("FYI: Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.")
	fmt.Printf("Value of y is : %v ",y)
	fmt.Println("Lets talk about Variable priority now: Local variables within the functions are priotized")
	fmt.Println("Rename the variable y to x and check the value")

}
