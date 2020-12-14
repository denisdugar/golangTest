package main

import (
	"fmt"
	"io"
	"os"
	"projects/day1/day3/numtoword"
	"strings"
)

func main() {
	file, err := os.Open("nums.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var s string
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		s = s + string(data[:n])
	}
	num := strings.Fields(s)

	file1, err1 := os.Create("text.txt")

	if err1 != nil {
		fmt.Println("Unable to create file:", err1)
		os.Exit(1)
	}
	defer file1.Close()

	for i := 0; i < len(num); i++ {
		file1.WriteString(numtoword.NumToWord(num[i]))
		file1.WriteString("\n")
	}
}
