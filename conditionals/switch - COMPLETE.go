package main

import (
	"fmt"
)

func main() {

	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Case 1. kubernetes with lower-case \"k\".")
	case "Kubernetes":
		fmt.Println("Case 2. Kubernetes with a cpital \"K\".")
	case "K8s":
		fmt.Println("Case 3. Kubernetes short name \"Kates\".")
	case "Istio":
		fmt.Println("Case 4. Service mesh is the future.")
	default:
		fmt.Println("Hit the default")
	}

}
