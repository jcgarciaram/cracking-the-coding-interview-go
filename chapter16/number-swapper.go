package main

import (
	"fmt"
)

func main() {
	a := 9
	b := -1
	a, b = numSwap(a, b)
	fmt.Println("a:", a, "b:", b)
	
}

func numSwap(a int, b int) (int, int) {
	a = a - b
	b = a + b
	a = b - a
	
	return a, b
}
	
