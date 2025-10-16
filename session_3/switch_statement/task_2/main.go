package main

import "fmt"

// Task 2: Type Inspector
func typeInspector(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("The type is int")
	case string:
		fmt.Println("The type is string")
	case bool:
		fmt.Println("The type is bool")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	fmt.Println("\n=== Type Inspector ===")
	typeInspector(42)
	typeInspector("Hello Go!")
	typeInspector(true)
	typeInspector(3.14)
}
