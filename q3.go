package main

import "fmt"

func primeFactors(n int) []int {
	factors := []int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

// find the largest prime factor of a given number
func Q3() int {
	n := 600851475143
	primeFactors := primeFactors(n)
	ans := primeFactors[len(primeFactors)-1]
	fmt.Printf("Answer to Q3: %v\n", ans)
	return ans
}
