package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Unit Testing")
}

func CalculateArea(width, length int) (int, error) {
	fmt.Printf("CalculateArea called with parameters %v and %v\n", width, length)

	if width <= 0 {
		return 0, errors.New("width is too low")
	}

	if length <= 0 {
		return 0, errors.New("length is too low")
	}

	return width * length, nil
}

type Calculator struct{}

func (c Calculator) Add(a, b float64) float64 {
	fmt.Printf("Add called with parameters %v and %v\n", a, b)
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	fmt.Printf("Subtract called with parameters %v and %v\n", a, b)
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	fmt.Printf("Multiply called with parameters %v and %v\n", a, b)
	return a * b
}

func (c Calculator) Divide(a, b float64) (float64, error) {
	fmt.Printf("Divide called with parameters %v and %v\n", a, b)
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
