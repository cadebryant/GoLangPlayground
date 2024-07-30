package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(NextPerfectNumber(5))
}

func NextPerfectNumber(n int) int {
	perfectNumbers := ListPerfectNumbers(20)
	found := false
	result := n + 1
	for !found {
		_, found = perfectNumbers[result]
		if !found {
			result++
		}
	}
	return result
}

func IsPerfect(n int) bool {
	perfectNums := ListPerfectNumbers(20)
	return perfectNums[n]
}

func ListPerfectNumbers(n int) map[int]bool {
	result := make(map[int]bool)
	primes := ListPrimes(n)
	for _, p := range primes {
		m := int(math.Pow(2, float64(p))) - 1
		if IsPrime(m) {
			result[int(math.Pow(2, float64(p-1)))*m] = true
		}
	}
	return result
}

func ListPrimes(n int) []int {
	primeArray := SieveOfEratosthenes(n)
	var result []int
	for i := 0; i <= n; i++ {
		if primeArray[i] {
			result = append(result, i)
		}
	}
	return result
}

func IsPrime(n int) bool {
	primeArray := SieveOfEratosthenes(n)
	return primeArray[n]
}

func SieveOfEratosthenes(n int) []bool {
	primeArray := make([]bool, n+1)
	sqrtN := math.Sqrt(float64(n))
	for i := 2; i <= n; i++ {
		primeArray[i] = true
	}
	for i := 2; float64(i) <= sqrtN; i++ {
		if primeArray[i] {
			primeArray = CrossOffMultiples(primeArray, i)
		}
	}
	return primeArray
}

func CrossOffMultiples(primeArray []bool, n int) []bool {
	len := len(primeArray)
	for i := 2 * n; i < len; i += n {
		primeArray[i] = false
	}
	return primeArray
}
