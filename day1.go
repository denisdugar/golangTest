package main

import "fmt"
import "projects/sort"

func main() {
    //arr := []int{5,6,7,0}
    var arr []int
    var arr1 []int
    arr1 = sort.SortT(arr)	
    fmt.Println(arr1)
}