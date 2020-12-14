package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"projects/day1/day3/numtoword"
)

func CloseFileError(f io.Closer, name string) {
	err := f.Close()
	if err != nil {
		log.Println("Error open file %s (err: %v", name, err)
	}
}

func main() {
	file, err := os.Open("nums.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer CloseFileError(file, "nums.txt")

	var s []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file1, err := os.Create("text.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer CloseFileError(file1, "text.txt")

	for i := 0; i < len(s); i++ {
		word, _ := numtoword.NumToWord(s[i])
		file1.WriteString(word)
		file1.WriteString("\n")
	}
}
