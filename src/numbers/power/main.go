package main

import "fmt"

func main() {
	fmt.Println(Power(2, 2))
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
