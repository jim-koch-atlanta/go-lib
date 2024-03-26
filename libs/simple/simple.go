package simple

import "fmt"

func plus(x int, y int) int {
	result := x + y
	fmt.Printf("%d + %d = %d", x, y, result)
	return result
}

func minus(x int, y int) int {
	result := x - y
	fmt.Printf("%d - %d = %d", x, y, result)
	return result
}
