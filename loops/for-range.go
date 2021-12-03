package main

import (
	"fmt"
)

func main() {

	coursesInProg := []string{
		"Docker & Kubernetes: The Big Picture",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive"}
	coursesCompleted := []string{
		"Docker & Kubernetes: The Big Picture",
		"Docker Deep Dive"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nHey we found a clash.", i, "is in both lists.")
			}
		}
	}
}
