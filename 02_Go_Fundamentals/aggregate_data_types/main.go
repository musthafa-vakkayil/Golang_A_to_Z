package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// Aggregate Data Types -> Array, Slices, Maps and Structs

	// Arrays - Fixed length

	var arr [2]int

	// prints [0, 0, 0]
	fmt.Println(arr)

	arrLiteral := [3]int{1, 23, 3}

	fmt.Println(arrLiteral)

	arrLiteral[1] = 99

	// Arrays are copied by value
	arr2 := arr
	fmt.Println(arr2)

	// arr == arr2 - true - arrays are comparable

	// Slices - are reference types

	var s []int    // slices of ints
	fmt.Println(s) // prints - [](nil)

	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	s = append(s1, 5, 10, 15)

	fmt.Println(s)

	s = slices.Delete(s, 1, 3) // remove indices 1,2 from slice
	fmt.Println(s)

	// slices are copied by reference
	// use slices.Clone to clone - copy by value

	// s == s2 - not comparable, compile time error

	fmt.Println("Demo")

	fmt.Println("Student Scores")

	fmt.Println(strings.Repeat("-", 14))

	students := []string{"Berlin",
		"John",
	}

	fmt.Println(students)

	fmt.Println(strings.Repeat("-", 14))

	scores := []int{99, 90}

	fmt.Println(scores)

	// Maps - Reference type

	var m map[string]int

	fmt.Println(m)

	m = map[string]int{"foo": 1, "bar": 2}

	fmt.Println(m["foo"])

	delete(m, "foo")

	m["baz"] = 418

	fmt.Println(m)

	// 0 - queries always return results
	fmt.Println(m["foo"])

	v, ok := m["foo"] // ok will be tru if the value is present

	fmt.Println(v, ok)

	// maps are copied by reference
	// use maps.Clone to clone - copy by value

	studensMap := map[string]int{"John": 90, "Adam": 99}

	fmt.Println(studensMap)

	// Structs - fixed size but can have multiple types
	// Struct is a value type
	var st struct {
		name string
		id   int
	}

	fmt.Println(st)

	st.name = "Arthur"

	fmt.Println(st.name)

	type myStrcut struct {
		name string
		id   int
	}

	st1 := myStrcut{
		name: "Musthafa",
		id:   4,
	}

	fmt.Println(st1)

	// structs are comparable
}
