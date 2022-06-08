package algo_test

import (
	"algo/algo"
	"testing"
)

func TestMatchesNum(t *testing.T) {
	arr := []int{1,1,3,3,5,7,2,2,4,4,6,4,8,7}
	k := 2
	expect := 6
	got := algo.MatchesNum(arr, k)
	if got != expect {
		t.Errorf("players:%v,k:%d,expect:%d,got:%d\n", arr, k, expect, got)
	}
}
