package complex

import (
	"fmt"

	simple "github.com/thebigkoch/go-lib/libs/simple"
)

func Times(x int, y int) int {
	result := 0
	for i := 0; i < x; i++ {
		result = simple.Plus(result, y)
	}
	fmt.Printf("%d * %d = %d", x, y, result)
	return result
}
