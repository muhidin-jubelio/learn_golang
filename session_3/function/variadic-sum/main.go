package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("sumAll(1, 2, 3) =", sumAll(1, 2, 3))
	fmt.Println("sumAll(4, 5, 6, 7, 8) =", sumAll(4, 5, 6, 7, 8))
	fmt.Println("sumAll() =", sumAll())
}
