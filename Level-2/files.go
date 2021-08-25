package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
var filePath = "/Users/ankitkumar/personal/goPractice/Level-2/sample_file.txt"
func main() {
	fmt.Println("1. Lets check error before checking file")
	if _,err := os.Stat(filePath)
		os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
	} else{
		fmt.Println("So no error, good to go")
	}
	fmt.Println("2. File exist in go? (enter correct path for file Read)")
	if _,err := os.Stat(filePath)
	err == nil {
		fmt.Println("File exist though")
	} else{
		fmt.Println("It doesn't exist")
	}

	fmt.Println("3. Read file in go at once")
	fx, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error is : %v\n",err)
	}

	//convert bytes to string
	str := string(fx)

	fmt.Printf("File Data: \n%v\n", str)

	/*
		Use Case: If you have a very large file, line by line is better because otherwise it won’t fit into memory.
	*/
	fmt.Println("4. Read file in go line by line")
	lines, err := readLines(filePath)
	if err != nil {
		log.Fatalf("readLines : %s\n", err) //The fatal function is just like the print function the difference is that it calls os.Exit(1) at the end.
	}

	for i, line := range lines {
		fmt.Println(i, line)
	}


	fmt.Println("5. Write file in go")
	file, err := os.Create("Sample_file.md")

	if err != nil {
		return
	}
	defer file.Close()
	var a = []string{"Rio", "London", "Berlin"}
	file.WriteString("# New File Created -- Hi! there.\n")
	for i := 0; i<len(a);i++ {
		file.WriteString(a[i])
		file.WriteString("\n")
	}


	fmt.Println("6. Rename file , change file name according to you")
	source := "Sample_file.md"
	destination := "sample_file.md"

	os.Rename(source, destination)


}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() //defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.

	var lines []string
	scanner := bufio.NewScanner(file) //Package bufio implements buffered I/O.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}




