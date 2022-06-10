package algo_test

import (
	"algo/algo"
	"testing"
)

var jumpArr = [][]int {
	{4,1,1,7,1,1,1,1,1,1,1},
}

var jumpAns = []int {4}

func TestJumpGame(t *testing.T) {
	for i,v := range jumpArr {
		got := algo.JumpGame(v)
		if got != jumpAns[i] {
			t.Errorf("arr:%v,expect:%v,got:%v\n", v, jumpAns[i], got)
		}
	}
}
