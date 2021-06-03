package main

import (
	"algo/algo"
	"fmt"
)

func main() {
	lst := []int{4, 2, 1, 5, 6, 3, 8, 9}
	algo.InsertSort(lst)
	lst2 := []int{4, 3, 6, 5, 1, 8, 6, 9, 2, 23, 12, 31, 45, 23}
	algo.MergeSort(lst2)
	fmt.Printf("merge sort:%v\n", lst2)
}