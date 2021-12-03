package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	gettingStartedWithK8s := courseMeta{
		author: "Nigel Poulton",
		level:  "Intermediate",
		rating: 5,
	}

	fmt.Println("Author of Getting Started with Kubernetes is", gettingStartedWithK8s.author)
	gettingStartedWithK8s.rating = 1.2
	fmt.Println("Course rating is currently", gettingStartedWithK8s.rating)
}
