package main

import (
	"fmt"
	"os"
	"projects/day1/day5/CalcStr"
	"projects/day1/day5/RewriteLine"
)

func main() {
	if len(os.Args) == 3 {
		fmt.Println(CalcStr.CalcStr(os.Args[1], os.Args[2]))
	}
	if len(os.Args) == 4 {
		RewriteLine.RewriteLine(os.Args[1], os.Args[2], os.Args[3])
	}
}
