package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ListPrimes((1000)))
}

func IsPrime(n int) bool {
	primeBooleans := SieveOfEratosthenes(n)
	return primeBooleans[n]
}

func ListPrimes(n int) []int {
	primeList := make([]int, 0)
	primeBooleans := SieveOfEratosthenes(n)
	for p := 0; p < n+1; p++ {
		if primeBooleans[p] {
			primeList = append(primeList, p)
		}
	}
	return primeList
}

func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)
	sqrtN := math.Sqrt(float64(n))
	for k := 2; k <= n; k++ {
		primeBooleans[k] = true
	}
	for p := 2; float64(p) <= sqrtN; p++ {
		if primeBooleans[p] {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}
	}
	return primeBooleans
}

func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1
	for k := 2 * p; k <= n; k += p {
		primeBooleans[k] = false
	}
	return primeBooleans
}
