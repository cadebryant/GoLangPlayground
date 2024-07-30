package main

import (
	"fmt"
)

func main() {
	fmt.Println(FactorialArray(20))
}

func FactorialArray(n int) []int {
	result := make([]int, n+1)
	result[0] = 1
	for i := 1; i <= n; i++ {
		result[i] = result[i-1] * i
	}

	return result
}
