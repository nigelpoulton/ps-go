package main

import (
	"fmt"
)

func main() {

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2HeadWins := map[string]int{
		"Sunderland": 6,
		"Newcastle":  0,
	}

	fmt.Printf("League titles: %v\nRecent Head to heads: %v\n",
		leagueTitles, recentHead2HeadWins)

}
