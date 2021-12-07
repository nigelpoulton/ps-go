package main

import (
	"fmt"
)

func main() {

	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}


	testMap["A"] = 100
	testMap["J"] = 1973
	fmt.Println(testMap)

	delete(testMap, "J")
	fmt.Println(testMap)
}
