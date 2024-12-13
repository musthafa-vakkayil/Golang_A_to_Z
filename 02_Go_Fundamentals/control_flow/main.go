package main

import (
	"fmt"
	"strconv"
)

func main() {
	// if stataments -.  if test {....}

	type score struct {
		name  string
		score int
	}

	scores := []score{
		{"John", 90},
		{"Adam", 80},
		{"Berlin", 70},
		{"Alice", 100},
	}

	fmt.Println("Enter Student name")

	var stName string

	fmt.Scanln(&stName)

	if stName == "John" {
		fmt.Println(scores[0])
	} else if stName == "Adam" {
		fmt.Println(scores[1])
	} else if stName == "Berlin" {
		fmt.Println(scores[2])
	} else if stName == "Alice" {
		fmt.Println(scores[3])
	} else {
		fmt.Println("Not Found")
	}

	// Switch Statements

	fmt.Println("Enter Student name")

	fmt.Scanln(&stName)

	switch stName {
	case "John":
		fmt.Println(scores[0])
	case "Adam":
		fmt.Println(scores[1])
	case "Berlin":
		fmt.Println(scores[2])
	case "Alice":
		fmt.Println(scores[3])
	default:
		fmt.Println("Not Found")
	}

	// Loops
	// for {...} - infinite loop
	// for condition {...} - condition based loop
	// for initializer; test; post cluase {...} // counter-based loop

	scoresNew := []score{}

	// infinite loop
	for {
		fmt.Println("Enter a student name and score")

		var name, rawString string

		fmt.Scanln(&name, &rawString)

		s, _ := strconv.Atoi(rawString)

		scoresNew = append(scoresNew, score{name, s})

		fmt.Println(scoresNew)
	}

	// Looping With Collections

	// for key, value := range collection {...} -> works for array slice and map

	arr := [3]int{101, 102, 103}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("Done!")

}
