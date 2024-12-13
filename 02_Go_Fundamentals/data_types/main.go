package main

import (
	"fmt"
	"strings"
)

func main() {
	// Simple Data Types -> Strings, Numbers, Boolean and Errors

	// Strings
	// "this is string" - interpreted string - \n will be new line
	// `this also string` -  raw string - use enter to make new line

	// Numbers
	// int, uint, float32, float64, complex64 and complex128

	// Booleans
	// true aand false - no truthy and falsy (0,1)

	// Errors
	// built in interface type for error, if no error = nill
	// type error interfacte{ Error() string }

	// Declaring Variables
	// var Name string - assigns zero value
	// var myName string = "Adam John"
	// var myName string = "Adam John"  - initialize with inferred type
	// myName := "John Adam" -  short declaration syntax

	// Constants
	// const a = 42 - implicit typed
	// const b string = "hello world" - explicitly typed
	// const ( d = true  e = 11)

	// const e = someFunc() - this won't work, can't evaluate at compile meting

	fmt.Println("Demo")

	fmt.Println("Student Scores")

	var name string = "Adam John"
	// var score = 90
	score := 90
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(name, score)

	nameNew, scoreNew := "Alice", 70

	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(nameNew, scoreNew)

	// Pointers
	a := 43
	b := &a

	fmt.Println(*b)

	c := new(int)

	fmt.Println(c) // create a pointer to anonymous variables
}
