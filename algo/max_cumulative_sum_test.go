package algo_test

import (
	"algo/algo"
	"testing"
)

var arr = [][]int{
	{1,2,3,4,5},
	{-1,-2,-3,-4,-5},
	{},
	{4},
	{-4},
	{5,3,-2,7,-10,7,-1,9},
}
var ans = []int {15,-1,0,4,-4,18}
var benchArr = []int{5,3,-2,7,-10,7,-1,9,5,3,-2,7,-10,7,-1,9,5,3,-2,7,-10,7,-1,9,8,3,6,1,-9,7,-5,34,-4,-6,43,-14}

func TestMaxCumulativeSum1(t *testing.T) {
	for k, val := range arr {
		rarr, sum := algo.MaxCumulativeSum1(val)
		if sum != ans[k] {
			t.Errorf("arr:%v,expect:%v,got:%v\n", val, ans[k], sum)
		} else {
			t.Logf("get sub array:%v,sum:%v\n", rarr, sum)
		}
	}
}

func TestMaxCumulativeSum2(t *testing.T) {
	for k, val := range arr {
		rarr, sum := algo.MaxCumulativeSum2(val)
		if sum != ans[k] {
			t.Errorf("arr:%v,expect:%v,got:%v\n", val, ans[k], sum)
		} else {
			t.Logf("get sub array:%v,sum:%v\n", rarr, sum)
		}
	}
}

func BenchmarkMaxCumulativeSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.MaxCumulativeSum1(benchArr)
	}
}

func BenchmarkMaxCumulativeSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.MaxCumulativeSum2(benchArr)
	}
}