package main

import "fmt"

func main(){
	fmt.Println("Maps in go : A Golang map is a unordered collection of key-value pairs. For every key, there is a unique value. If you have a key, you can lookup the value in a map.")
	fmt.Println("Also known as associative array or hash tables")
	fmt.Println("1. Sample maps ")
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2

	fmt.Printf("Result: %v\n", m)
	fmt.Println("2. Or you can write it like hashmaps ")
	hm := map[string]int{
		"A" : 1,
		"B" : 2,
		"C" : 3,
	}
	fmt.Println("This do exactly same but elegant notations")
	fmt.Printf("new Result : %v\n", hm)
	fmt.Println("3. Store Metadata, map of strings to map of strings to strings")
	website := map[string]map[string]string{
		"Gmail": map[string]string{
			"name": "Gmail",
			"type": "Mail",
		},
		"Youtube": map[string]string{
			"name": "Youtube",
			"type": "Videos",
		},
	}
	fmt.Println("Metadata looks like")
	fmt.Println(website)
	fmt.Println("4. Access keys in maps")
	fmt.Println(website["Gmail"]["type"])
	fmt.Println(website["Youtube"]["type"])




}
/*
	Map is not ordered in go.
 */