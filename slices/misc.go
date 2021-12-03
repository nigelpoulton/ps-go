package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	for _, i := range mySlice {
		fmt.Println(i)
	}

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Printf("mySlice NOW contains %d and has a new length of %d and capacity of %d\n",
		mySlice, len(mySlice), cap(mySlice))

}
