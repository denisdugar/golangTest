package RewriteLine

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CloseFileError(f io.Closer, name string) {
	err := f.Close()
	if err != nil {
		log.Println("Error open file %s (err: %v)", name, err)
	}
}

func RewriteLine(fileName string, search string, newLine string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer CloseFileError(file, fileName)

	var s []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == search {
			s = append(s, newLine)
		} else {
			s = append(s, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file1, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer CloseFileError(file1, fileName)

	for i := 0; i < len(s); i++ {
		file1.WriteString(s[i])
		file1.WriteString("\n")
	}
	return err
}
