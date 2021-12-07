package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "nigel poulton"
	course := "getting started with kubernetes"

	fmt.Println(converter(author, course))
}

//It's not recommended to change the names of arguments like this (author > s1 > str1) as it causes confusion
//This is only shown here as an example of what "can" be done 
func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.ToUpper(s1)
	s2 = strings.Title(s2)

	return s1, s2
}
