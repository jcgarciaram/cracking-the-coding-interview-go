package main

import (
	"fmt"
)

func main() {
	a := 9
	b := -7
	fmt.Println(numberMax(a,b))
}

func numberMax(a, b int) int {
	slice := []int{a, b}
	
	c := a - b
	
	pos := (uint64(c) | 0) >> 63
	
	return slice[pos]
}
	
	
