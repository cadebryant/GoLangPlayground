package main

import (
	"fmt"
)

func main() {
	sample1 := map[string]int{
		"lions":  2,
		"tigers": 4,
	}
	sample2 := map[string]int{
		"bears":  2,
		"lions":  4,
		"tigers": 3,
	}
	fmt.Println(SumOfMaxima(sample1, sample2))
}

func SimpsonsIndex(sample map[string]int) float64 {
	sum := float64(SumOfValues(sample))
	var result float64 = 0
	for _, count := range sample {
		result += (float64(count) / sum) * (float64(count) / sum)
	}
	return result
}

func Richness(sample map[string]int) int {
	var richness = 0
	for _, count := range sample {
		if count > 0 {
			richness++
		}
	}
	return richness
}

func SumOfValues(sample map[string]int) int {
	var sum = 0
	for _, count := range sample {
		sum += count
	}
	return sum
}

func SumOfMinima(sample1, sample2 map[string]int) int {
	var sum = 0
	allKeys := make(map[string]bool)
	for key := range sample1 {
		allKeys[key] = true
	}
	for key := range sample2 {
		allKeys[key] = true
	}
	for key := range allKeys {
		sum += Min2(GetMapValue(sample1, key), GetMapValue(sample2, key))
	}
	return sum
}

func SumOfMinima2(sample1, sample2 map[string]int) int {
	var sum = 0
	for key := range sample1 {
		sum += Min2(GetMapValue(sample1, key), GetMapValue(sample2, key))
	}
	return sum
}

func SumOfMaxima(sample1, sample2 map[string]int) int {
	var sum = 0
	allKeys := make(map[string]bool)
	for key := range sample1 {
		allKeys[key] = true
	}
	for key := range sample2 {
		allKeys[key] = true
	}
	for key := range allKeys {
		sum += Max2(GetMapValue(sample1, key), GetMapValue(sample2, key))
	}
	return sum
}

func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func GetMapValue(m map[string]int, key string) int {
	val, ok := m[key]
	if ok {
		return val
	}
	return 0
}

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	minimaSum := float64(SumOfMinima(sample1, sample2))
	avgSumValues := (float64(SumOfValues(sample1)) + float64(SumOfValues(sample2))) / 2.0
	return 1.0 - (minimaSum / avgSumValues)
}

func JaccardDistance(sample1, sample2 map[string]int) float64 {
	minimaSum := float64(SumOfMinima(sample1, sample2))
	maximaSum := float64(SumOfMaxima(sample1, sample2))
	return 1 - (minimaSum / maximaSum)
}
