package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("1. Input a string")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Output here: Hello ", name)
	fmt.Println("2. Input a number")
	var num int32
	fmt.Print("Get a number from terminal and check its between 1 and 100 :")
	fmt.Scanf("%d",&num)
	if num <= 100 && num >= 0 {
		fmt.Println("Number is valid")
	} else {
		fmt.Println("Number is not valid")
	}
	fmt.Println("3. Read input from a file.txt and print here")
	file, err:= os.Open("/Users/ankitkumar/personal/goPractice/Level-1/numbers.txt")
	if err != nil{
		fmt.Println("error while loading file")
		os.Exit(1)
	}
	count:= 0
	for {
		perLine:= 0
		_,err := fmt.Fscanf(file, "%d\n", &perLine)
		if err != nil{
			if err == io.EOF {
				break //stop reading file after end
			}
			fmt.Println("Error is : err")
			os.Exit(1)
		}
		count++
		fmt.Printf("%d. Number is : %d\n", count, perLine)
	}
}

/*
Blank Identifier : _ or Blank Identifier , Golang has a special feature to define and use the
unused variable using Blank Identifier. Unused variables are those variables which are defined by the user throughout the program
but he/she never make the use of these variables.
The real use of Blank Identifier comes when a function returns multiple values,
but we need only a few values and want to discard some values.
*/
