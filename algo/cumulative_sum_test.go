package algo_test

import (
	"algo/algo"
	"testing"
)

func TestAccumulate(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,18,19,20,30,24,12}
	sum := 50
	expect := 5440

	rst := algo.Accumulate(arr, sum)
	if rst != expect {
		t.Errorf("arr %v,sum:%v,expect:%d,get:%d\n", arr, sum, expect, rst)
	}
}

func TestAccumulateDp(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,18,19,20,30,24,12}
	sum := 50
	expect := 5440

	rst := algo.AccumulateDp(arr, sum)
	if rst != expect {
		t.Errorf("arr %v,sum:%v,expect:%d,get:%d\n", arr, sum, expect, rst)
	}
}

func BenchmarkAccumulate(b *testing.B) {
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,18,19,20,30,24,12}
	sum := 50
	for i := 0; i < b.N; i++ {
		algo.Accumulate(arr, sum)
	}
}

func BenchmarkAccumulateDp(b *testing.B) {
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,18,19,20,30,24,12}
	sum := 50
	for i := 0; i < b.N; i++ {
		algo.AccumulateDp(arr, sum)
	}
}