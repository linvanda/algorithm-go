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
	lst3 := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	rst := algo.MaxSubArray(lst3)
	fmt.Printf("max sub array:%v\n", rst)
}