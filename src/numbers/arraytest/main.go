package main

import "fmt"

func main() {
	arr := make([]int, 10)
	fmt.Println(arr)
	ModifyArray(&arr)
	result := &arr
	fmt.Println(result)
}

func ModifyArray(arr *[]int) {
	*arr = append(*arr, 999)
}
