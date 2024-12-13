package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	result := calculateAverage(numbers)
	fmt.Printf("The average is: %d\n", result)
}

func calculateAverage(nums []int) int {
	total := 0
	for i := 0; i <= len(nums); i++ {
		total += nums[i]
	}
	average := total / len(nums)
	return average
}
