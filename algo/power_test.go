package algo_test

import (
	"algo/algo"
	"testing"
)

var powerVals = [][3]int {
	{10, 0, 1},
	{9, 1, 9},
	{5, 3, 125},
	{6, 12, 2176782336},
}

func TestPower(t *testing.T) {
	for _, v := range powerVals {
		r := algo.Power(v[0], v[1])
		if r != v[2] {
			t.Errorf("%d^%d,expect:%d,got:%d\n", v[0], v[1], v[2], r)
		}
	}
}
