package main

import "fmt"

func fizzbuzz() {
	// ?
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
				continue
			}
			fmt.Println("Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
