package main

import (
	"fmt"
)

func main() {

	courses := []string{"Docker & Kubernetes: The Big Picture",
		"Getting Started with Docker",
		"Getting Started with Kubernetes"}

	fmt.Printf("Length of slice is %d and capacity is %d\n",
		len(courses), cap(courses))

	fmt.Println(courses)
}
