package main

import "fmt"

// find the sum of even-valued terms in Fibonacci

var fibMap = make(map[int]int64)

func fib(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	} else {
		if val, ok := fibMap[n]; ok {
			return val
		}
		ans := fib(n-1) + fib(n-2)
		fibMap[n] = ans
		return ans
	}
}

func Q2() int64 {
	var ans int64
	for i := 1; i < 100; i++ {
		val := fib(i)
		if val%2 == 0 {
			ans += val
		}
		if val > 4000000 {
			break
		}
	}
	fmt.Printf("Answer to Q2: %d\n", ans)
	return ans
}
