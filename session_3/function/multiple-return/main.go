package main

import (
	"fmt"
)

// Task 1: Multiple Return Values
func calculate(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	sum, product := calculate(4, 5)
	fmt.Println("Task 1:")
	fmt.Printf("Sum: %d, Product: %d\n\n", sum, product)
}
