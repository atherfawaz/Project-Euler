package main

import (
	"fmt"
	"strconv"
)

// largest palindrome that is the product of two three-digit integers

func isPalindrome(strNum string) bool {
	j := len(strNum) - 1
	for i := 0; i < j; i++ {
		if strNum[i] != strNum[j] {
			return false
		}
		j--
	}
	return true
}

func Q4() int {
	maxSoFar := -1
	for i := 999; i >= 99; i-- {
		for j := 999; j >= 99; j-- {
			prod := i * j
			p := isPalindrome(strconv.Itoa(prod))
			if p && prod > maxSoFar {
				maxSoFar = prod
			}
		}
	}
	fmt.Printf("Answer to Q4: %d\n", maxSoFar)
	return maxSoFar
}
