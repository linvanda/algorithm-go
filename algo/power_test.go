package algo_test

import (
	"algo/algo"
	"math"
	"testing"
)

var powerVals = [][3]int {
	{10, 0, 1},
	{9, 1, 9},
	{5, 3, 125},
	{6, 12, 2176782336},
	{-3, 4, 81},
	{-3, 3, -27},
}

var benchPowerB = 4
var benchPowerP = 20

func TestPower(t *testing.T) {
	for _, v := range powerVals {
		r := algo.Power(v[0], v[1])
		if r != v[2] {
			t.Errorf("%d^%d,expect:%d,got:%d\n", v[0], v[1], v[2], r)
		}
	}

	p1 := algo.Power(benchPowerB, benchPowerP)
	p2 := math.Pow(float64(benchPowerB), float64(benchPowerP))
	if float64(p1) != p2 {
		t.Errorf("%d^%d,expect:%f,got:%d\n", benchPowerB, benchPowerP, p2, p1)
	}
}

func BenchmarkQuickPower(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		algo.Power(benchPowerB, benchPowerP)
	}
}

func BenchmarkInnerPower(b *testing.B)  {
	x := float64(benchPowerB)
	y := float64(benchPowerP)
	for i := 0; i < b.N; i++ {
		math.Pow(x, y)
	}
}