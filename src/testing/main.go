package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	inputs := []int{2, 4, 5}

	for i := range inputs {
		fmt.Println(solve(int32(inputs[i])))
	}
}

func solve(n int32) []int32 {
	one := big.NewInt(1)
	tmpResult := []*big.Int{one}
	result := []int32{1}
	mod := int64(math.Pow(10, 9))
	bigMod := big.NewInt(mod)
	bigN := big.NewInt(int64(n))
	i := int32(0)
	for ; i < n; i++ {
		tmpItem := new(big.Int).Set(tmpResult[i])
		tmpN := new(big.Int).Set(bigN)
		num := tmpN.Sub(tmpN, big.NewInt(int64(i))).Mul(tmpN, tmpItem)
		denom := big.NewInt(int64(i))
		denom.Add(denom, one)
		val := num.Div(num, denom)
		tmpResult = append(tmpResult, val)
	}

	j := int32(1)
	for ; j <= n; j++ {
		tmpVal := new(big.Int).Set(tmpResult[j])
		modVal := tmpVal.Mod(tmpVal, bigMod)
		result = append(result, int32(modVal.Int64()))
	}
	return result
}
