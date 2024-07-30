package main

import (
	"fmt"
)

func main() {
	fmt.Println(ListMersennePrimes(60))
}

func ListMersennePrimes(n int) []int {
	var result []int
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			m := Power(2, i) - 1
			if IsPrime(m) {
				result = append(result, m)
			}
		}
	}
	return result
}

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func Power(num int, pow int) int {
	if pow == 0 {
		return 1
	}
	var result = 1
	for pow > 0 {
		if pow%2 != 0 {
			result *= num
		}
		num *= num
		pow >>= 1
	}
	return result
}
