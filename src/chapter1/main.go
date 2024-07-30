package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
)

func main() {
	resp, err := http.Get("https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	text := string(bytes[:])
	fmt.Println(MinimumSkew(text))
	// k := 9
	// L := 500
	// t := 3
	// clumps := FindClumps(text, k, L, t)
	// fmt.Println(len(clumps))
}

func MaxArray(arr []int) int {
	len := len(arr)
	max := math.MinInt
	for i := 0; i < len; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func FindClumps(text string, k int, L int, t int) []string {
	patterns := map[string]bool{}
	var result []string
	n := len(text)
	for i := 0; i <= n-L; i++ {
		window := text[i : i+L]
		freqMap := FrequencyTable(window, k)
		for s := range freqMap {
			if freqMap[s] >= t {
				patterns[s] = true
			}
		}
	}
	for p := range patterns {
		result = append(result, p)
	}
	return result
}

func PatternCount(pattern string, text string) int {
	n := len(text)
	k := len(pattern)
	patCount := 0
	for i := 0; i <= n-k; i++ {
		if text[i:i+k] == pattern {
			patCount++
		}
	}
	return patCount
}

func Contains(strings []string, pattern string) bool {
	for _, s := range strings {
		if pattern == s {
			return true
		}
	}
	return false
}

func FrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)
	freqMap := FrequencyTable(text, k)
	max := MaxMap(freqMap)
	for str := range freqMap {
		if freqMap[str] == max {
			freqPatterns = append(freqPatterns, str)
		}
	}
	return freqPatterns
}

func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)
	var pattern string
	for i := 0; i <= n-k; i++ {
		pattern = text[i : i+k]
		_, ok := freqMap[pattern]
		if ok {
			freqMap[pattern]++
		} else {
			freqMap[pattern] = 1
		}
	}
	return freqMap
}

func MaxMap(dict map[string]int) int {
	max := math.MinInt
	for _, count := range dict {
		if count > max {
			max = count
		}
	}
	return max
}

func Reverse(str string) string {
	len := len(str)
	rev := make([]byte, len)
	for i := range str {
		rev[i] = str[len-i-1]
	}
	return string(rev)
}

func Complement(str string) string {
	cmp := make([]byte, len(str))
	for i, symbol := range str {
		switch symbol {
		case 'A':
			cmp[i] = 'T'
		case 'C':
			cmp[i] = 'G'
		case 'G':
			cmp[i] = 'C'
		case 'T':
			cmp[i] = 'A'
		default:
			panic("Invalid symbol in string given to Complement().")
		}
	}
	return string(cmp)
	// var cmp strings.Builder
	// var cmpMap = map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	// len := len(str)
	// for i := 0; i < len; i++ {
	// 	cmp.WriteByte(cmpMap[str[i]])
	// }
	// return cmp.String()
}

func ReverseComplement(str string) string {
	return Reverse(Complement(str))
}

func StartingIndices(pattern string, text string) []int {
	positions := make([]int, 0)
	n := len(text)
	k := len(pattern)
	for i := 0; i <= n-k; i++ {
		if text[i:i+k] == pattern {
			positions = append(positions, i)
		}
	}
	return positions
}

func MinimumSkew(genome string) []int {
	var indices []int
	arr := SkewArray(genome)
	min := MinIntegerArray(arr)
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == min {
			indices = append(indices, i)
		}
	}
	return indices
}

func SkewArray(genome string) []int {
	n := len(genome)
	arr := make([]int, n+1)
	arr[0] = 0
	chars := []byte(genome)
	for i := 1; i <= n; i++ {
		arr[i] += arr[i-1] + Skew(chars[i-1])
	}
	return arr
}

func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	}
	return 0
}

func MinIntegerArray(arr []int) int {
	n := len(arr)
	min := math.MaxInt
	for i := 0; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
