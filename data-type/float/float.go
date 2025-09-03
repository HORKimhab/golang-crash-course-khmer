package main

import (
	"fmt"
	"math"
)

func main() {
	// Cos float32 is max 3.4e+38
	var kh float32 = 3.4e+38
	var amountNumber float64 = 3.e+22

	fmt.Printf("Type: %T, value: %v\n", kh, kh)
	fmt.Printf("Type: %T, value: %v\n", amountNumber, amountNumber)

	// Exercise 1: Basic arithmetic
	fmt.Println("=== Exercise 1: Arithmetic ===")
	a := 10.5
	b := 2.3
	fmt.Println("Sum:", a+b)
	fmt.Println("Difference:", a-b)
	fmt.Println("Product:", a*b)
	fmt.Println("Quotient:", a/b)

	// Exercise 2: Rounding
	fmt.Println("\n=== Exercise 2: Rounding ===")
	num := 7.89
	fmt.Println("Original:", num)
	fmt.Println("Round:", math.Round(num))
	fmt.Println("Ceil:", math.Ceil(num))
	fmt.Println("Floor:", math.Floor(num))

	// Exercise 3: User input
	fmt.Println("\n=== Exercise 3: User Input ===")
	var x, y float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)
	fmt.Println("Average:", (x+y)/2)
}
