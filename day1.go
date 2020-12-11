package main

import (
	"fmt"
	"projects/sort"
)

func main() {
	arr := []int{5, 6, 7, 0}
	sort.SortT(arr)
	fmt.Println(arr)
}
