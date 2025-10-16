package main

import "fmt"

func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	nums := []int{5, 6, 5, 4, 5}
	result := sumSlice(nums)
	fmt.Println("Total:", result)
}
