package main

import (
	"fmt"
	"math"
)

func main() {
	//lst := []int{4, 2, 1, 5, 6, 3, 8, 9}
	//algo.InsertSort(lst)
	//lst2 := []int{4, 3, 6, 5, 1, 8, 6, 9, 2, 23, 12, 31, 45, 23}
	//algo.MergeSort(lst2)
	//fmt.Printf("merge sort:%v\n", lst2)
	//lst3 := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	//sub := algo.MaxSubArray(lst3)
	//fmt.Printf("max sub array:%v\n", sub)
	//lst4 := []int{4,2,6,5,9,1,2, 13}
	//heap := algo.NewMaxHeap(lst4)
	//size := heap.Size()
	//for i := 0; i < size; i++ {
	//	v := heap.Pop()
	//	fmt.Printf("heap pop:%v\n", v)
	//}
	//heap.Add(3)
	//heap.Add(1)
	//
	//fmt.Printf("heap pop2:%v\n", heap.Pop())
	//fmt.Printf("heap pop2:%v\n", heap.Pop())
	//
	//lst5 := []int{5, 4, 2, 7, 6, 8, 9, 12, 11, 4, 45, 34}
	//algo.QuickSort(lst5)
	//fmt.Printf("quick sort:%v\n", lst5)
	//
	//lst6 := []int{5, 4, 2, 7, 6, 5, 8, 12}
	//lst6 = algo.CountingSort(lst6, -1)
	//fmt.Printf("counting sort:%v\n", lst6)
	//
	//lst7 := []string{
	//	"18609184738",
	//	"13028782213",
	//	"18609184732",
	//	"13028722289",
	//	"13128782213",
	//}
	//lst7 = algo.RadixSort(lst7, 11)
	//fmt.Printf("radix sort:%v\n", lst7)
	//
	//str := "abcadcabadaf"
	//s := algo.GetLongestUnrepeatedSubStr(str)
	//fmt.Printf("str:%v,longest unrepeat sub str:%v\n", str, s)
	//
	//lst8 := []float64{0.5, 0.1, 0.4, 0.9, 0.3, 0.6}
	//fmt.Printf("lst:%v,bucket sorted:%v\n", lst8, algo.BucketSort(lst8))
	//
	//lst9 := []int{5, 4, 2, 7, 6, 8, 12}
	//v := algo.OrderStatistic(lst9, 6)
	//fmt.Printf("the %v mini value:%v\n", 7, v)

	//stack := algo.NewStack(5)
	//stack.Push(1)
	//stack.Push(3)
	////stack.Push(1)
	//fmt.Printf("stack:%v,%v\n", stack.Pop(), stack.Pop())

	//queue := algo.NewQueue(4)
	//for i := 0; i < 3; i++ {
	//	queue.Enqueue(i)
	//	fmt.Printf("len:%v,is full:%v\n", queue.Length(), queue.IsFull())
	//}
	//
	//println()
	//
	//for i := 0; i < 3; i++ {
	//	fmt.Printf("dequeue:%v, len:%v\n", queue.Dequeue(), queue.Length())
	//}
	//
	//println()
	//queue.Enqueue(10)
	//queue.Enqueue(20)
	//fmt.Printf("dequeue:%v, len:%v\n", queue.Dequeue(), queue.Length())
	//fmt.Printf("dequeue:%v, len:%v\n", queue.Dequeue(), queue.Length())
	//
	//queue.Enqueue(30)
	//fmt.Printf("dequeue:%v, len:%v\n", queue.Dequeue(), queue.Length())

	//link := algo.NewLink()
	//link.Append(1)
	//link.Append(2)
	//link.Append(3)
	//fmt.Printf("link length:%d\n", link.Length())
	//fmt.Printf("pop:%v\n", link.Pop())
	//fmt.Printf("shift:%v\n", link.Shift())
	//fmt.Printf("len:%d\n", link.Length())
	//link.Delete(2)
	//fmt.Printf("len:%d\n", link.Length())

	//input := []byte("hello")
	//hash := sha1.Sum(input)
	//fmt.Printf("%v\n", hash[:5])

	//x := struct {
	//	Name string `json:"name"`
	//}{}

	//s := struct {
	//	b string
	//	a int8
	//	c int32
	//}{"abc",0x01, 0x03040506}
	//
	//t := reflect.TypeOf(s)
	//n := t.NumField()
	//fmt.Printf("s size:%d\n", t.Size())
	//for i := 0; i < n; i++ {
	//	f := t.Field(i)
	//	fmt.Printf("name:%s,offset:%d,size:%d,align:%d\n", f.Name, f.Offset, f.Type.Size(), f.Type.Align())
	//}
	//fmt.Printf("val:%#v\n", *(*[32]byte)(unsafe.Pointer(&s)))

	//err := errors.New("some error")
	//err1 := fmt.Errorf("get error:%v", err)
	//err2 := fmt.Errorf("get error:%w", err)
	//println(reflect.TypeOf(err1).String())
	//println(reflect.TypeOf(err2).String())
	//fmt.Printf("%T\n", err1)
	//fmt.Printf("%T\n", err2)
	//
	//perr := &os.PathError{"write", "/tmp/", os.ErrPermission}
	//wperr := fmt.Errorf("some path error:%w", perr)
	//var patherr *os.PathError
	//fmt.Printf("%v\n", errors.Is(wperr, os.ErrPermission))
	//if errors.As(wperr, &patherr) {
	//	fmt.Printf("path:%v\n", patherr.Path)
	//}
	//testDefer()

	//s2 := "a3i92o1f8gf9129h3ab"//
	//fmt.Printf("%d\n", algo.LongestPalindrome1([]byte(s2)))

	fmt.Printf("%b\n", math.MaxFloat64)
}
