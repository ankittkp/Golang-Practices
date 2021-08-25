package main

import "fmt"

type Employee struct {
	name string
	role string
	salary int
}

func main()  {
	fmt.Println("A struct can bundle attributes together. If you create a struct, you can set a couple of variables for that struct. Those variables can be of any datatype")
	fmt.Println("1. Lets implement One")
	var emp Employee
	emp.name = "Mark"
	emp.role = "Sales"
	emp.salary = 10000

	fmt.Println("So the structure you created have these data")
	fmt.Printf("Name: %s\n",emp.name)
	fmt.Printf("Role: %s\n",emp.role)
	fmt.Printf("Salary: %d\n",emp.salary)
}
/*
	Structs instead of Class?
	Go does not provide classes, but it does provide structs. Methods can be added on structs.
	This provides the behaviour of bundling the data and methods that operate on the data together skin to a class.
 */
