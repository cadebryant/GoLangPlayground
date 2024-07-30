package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcd(30, 1000))
}

func gcd(a int, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
