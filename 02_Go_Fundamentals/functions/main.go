package main

import "fmt"

func main() {
	// Functions
	// func functionName(params)(retrun values{}

	// variadic parameters

}

func great(names ...string) {
	for _, n := range names {
		fmt.Println(n)
	}
}

// named return types

func divide(a int, b int) (result int, ok bool) {
	if b == 0 {
		return
	}

	result = a / b

	ok = true

	return
}
