package main

import (
	"algo/algo"
	"fmt"
)

func main() {
	lst := []int{4, 2, 1, 5, 6, 3, 8, 9}
	algo.InsertSort(lst)
	fmt.Printf("sorted:%v\n", lst)
}