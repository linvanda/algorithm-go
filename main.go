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
	sub := algo.MaxSubArray(lst3)
	fmt.Printf("max sub array:%v\n", sub)
	lst4 := []int{4,2,6,5,9,1,2, 13}
	heap := algo.NewMaxHeap(lst4)
	size := heap.Size()
	for i := 0; i < size; i++ {
		v := heap.Pop()
		fmt.Printf("heap pop:%v\n", v)
	}
	heap.Add(3)
	heap.Add(1)

	fmt.Printf("heap pop2:%v\n", heap.Pop())
	fmt.Printf("heap pop2:%v\n", heap.Pop())
}