package recursion

import (
	"fmt"
)

func main() {
	fact := factorial(3)
	fmt.Println("Factorial of 3 is:", fact)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial((n - 1))
}
