package main

import (
	"fmt"
	"Calculator/model"
)


var calculator = new(model.Calculator)

func main() {
	a, b := 10.0, 5.0
	sum(a, b)
	subtract(a, b)
	multiply(a, b)
	divide(a, b)
}

func sum(a, b float64) {
	sum := calculator.Add(a, b)
	fmt.Printf("Sum is: %f\n", sum)
}

func subtract(a, b float64) {
	sus := calculator.Subtract(a, b)
	fmt.Printf("Subtract is: %f\n", sus)
}

func multiply(a, b float64) {
	mul := calculator.Multiply(a, b)
	fmt.Printf("Multiply is: %f\n", mul)
}

func divide(a, b float64) {
	div := calculator.Divide(a, b)
	fmt.Printf("Divide is: %f\n", div)
}
