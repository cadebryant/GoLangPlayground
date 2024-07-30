package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 1000
	k := 998
	fmt.Println(CombinationBig(n, k))
}

func FactorialBig(n int) *big.Int {
	return FactorialRangeBig(n, 1)
}

func PermutationBig(n int, k int) *big.Int {
	if n == k {
		return FactorialBig(n)
	}
	if k == 0 {
		return big.NewInt(1)
	}
	if k == 1 {
		return big.NewInt(int64(n))
	}
	return FactorialRangeBig(n, n-k)
}

func CombinationBig(n int, k int) *big.Int {
	if k == 0 || k == n {
		return big.NewInt(1)
	}
	if k == 1 {
		return big.NewInt(int64(n))
	}
	facN := FactorialRangeBig(n, n-k)
	return facN.Div(facN, FactorialBig(k))
}

func FactorialRangeBig(n int, m int) *big.Int {
	result, one := big.NewInt(1), big.NewInt(1)
	bigN := big.NewInt(int64(n))
	bigM := big.NewInt(int64(m))
	for bigN.Cmp(bigM) == 1 {
		result.Mul(result, bigN)
		bigN.Sub(bigN, one)
	}
	return result
}
