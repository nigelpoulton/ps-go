package main

import (
	"fmt"
)

func main() {
	s := make(intSet)
	for i := 0; i < 4; i++ {
		s.put(i)
	}
	counts := make(map[int]int)
	for i := 0; i < 1024*1024; i++ {
		v, ok := s.get()
		if !ok {
			return
		}
		counts[v]++
	}
	for k, v := range counts {
		fmt.Printf("Value: %v, Count: %v\n", k, v)
	}
}

type intSet map[int]struct{}

func (s intSet) put(v int) {
	s[v] = struct{}{}
}
func (s intSet) get() (int, bool) {
	for k := range s {
		return k, true
	}
	return 0, false
}
