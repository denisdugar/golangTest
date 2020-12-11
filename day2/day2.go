package main

import "projects/day1/day2/unique"

func main() {
	a := []int{1, 1, 1, 1, 2, 2, 2}
	var b = unique.UniqueArr(a)
	for i := 0; i < len(b); i++ {
		print(b[i])
	}
}
