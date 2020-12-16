package CalcStr

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

func CalcStr(fileName string, s string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer CloseFileError(file, fileName)

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == os.Args[2] {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count, err
}
