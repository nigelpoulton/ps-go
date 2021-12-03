package main

import (
	"fmt"
)

func main() {

	if dddLengthMins, cawLengthMins := 275, 30; dddLengthMins > cawLengthMins {
		fmt.Println("Docker Deep Dive is longer than Containers on AWS Wavelength")
		if dddLengthMins > 240 {
			fmt.Println("But it's so long it'll put viewers to sleep!")
		}
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Both courses are the same length")
	} else {

		fmt.Println("Containers on AWS Wavelength must be longer than Docker Deep Dive")
	}

}
