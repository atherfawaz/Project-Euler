package main

import "fmt"

// Find the sum of all multiples of 3 and 5 below 1000
func Q1() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Answer to Q1: %d\n", sum)
	return sum
}
