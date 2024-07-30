package main

import (
	"fmt"
)

func main() {
	fmt.Println(MapTest())
}

func MapTest() bool {
	dict := make(map[int]bool)
	dict[1] = true
	ok := false
	_, ok = dict[1]
	return ok
}
