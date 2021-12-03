package main

import (
	"fmt"
)

func main() {
	name := "Nigel Poulton"
	course := "Getting started with Kubernetes"

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)

	fmt.Println("You’re currently watching", course)
}

func updateCourse(course *string) string {
	*course = "Getting Started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
