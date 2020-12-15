package main

import (
	"fmt"
	"os"
	"projects/day1/day4/calculatesuare"
)

func test(i int) (int, error){
	return i, nil
}

func main() {
	a, err := calculatesuare.CalculateSuare(os.Args)
	fmt.Println(a)
	fmt.Println(err)
}
