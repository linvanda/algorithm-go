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

	lst5 := []int{5, 4, 2, 7, 6, 8, 9, 12, 11, 4, 45, 34}
	algo.QuickSort(lst5)
	fmt.Printf("quick sort:%v\n", lst5)

	lst6 := []int{5, 4, 2, 7, 6, 5, 8, 12}
	lst6 = algo.CountingSort(lst6, -1)
	fmt.Printf("counting sort:%v\n", lst6)

	lst7 := []string{
		"18609184738",
		"13028782213",
		"18609184732",
		"13028722289",
		"13128782213",
	}
	lst7 = algo.RadixSort(lst7, 11)
	fmt.Printf("radix sort:%v\n", lst7)

	str := "aadfaaadfdffewqrffg"
	s := algo.GetLongestUnrepeatedSubStr(str)
	fmt.Printf("str:%v,longest unrepeat sub str:%v\n", str, s)

	lst8 := []float64{0.5, 0.1, 0.4, 0.9, 0.3, 0.6}
	fmt.Printf("lst:%v,bucket sorted:%v\n", lst8, algo.BucketSort(lst8))
}