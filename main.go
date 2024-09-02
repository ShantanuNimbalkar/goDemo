package main

import "fmt"

// Add function adds two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract function subtracts the second integer from the first.
func Subtract(a, b int) int {
	return a - b
}

// IsPositive checks if a number is positive.
func IsPositive(a int) bool {
	if a > 0 {
		return true
	} else if a < 0 {
		return false
	} else {
		return false
	}
}

func main() {
	fmt.Println("Add:", Add(3, 4))
	fmt.Println("Subtract:", Subtract(10, 4))
	fmt.Println("IsPositive:", IsPositive(5))
}
