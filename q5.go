package main

import "fmt"

// find the smallest possible number that is evenely divisible by all integers from 1 to 20

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
func Q5() int {
	ans := 1
	for i := 1; i <= 20; i++ {
		ans = lcm(ans, i)
	}
	fmt.Printf("Answer to Q5: %d\n", ans)
	return ans
}
