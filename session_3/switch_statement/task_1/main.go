package main

import "fmt"

// Task 1: Day of the Week
func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func main() {
	// Test Task 1
	fmt.Println("=== Day of the Week ===")
	for i := 0; i <= 8; i++ {
		fmt.Printf("%d -> %s\n", i, dayOfWeek(i))
	}
}
