package simple

import "fmt"

func Plus(x int, y int) int {
	result := x + y
	fmt.Printf("%d + %d = %d", x, y, result)
	return result
}

func Minus(x int, y int) int {
	result := x - y
	fmt.Printf("%d - %d = %d", x, y, result)
	return result
}
