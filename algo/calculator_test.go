package algo_test

import (
	"algo/algo"
	"testing"
)

var calcQsn = map[string]float64 {
	"1+2+3": 6.0,
	"1-2*3": -5,
	"4*(2+3)": 20,
	"1*(2+3)-4*(6/(3-4))":29.0,
	"-4*(2-34)-46+33 - 12 / 4 * (22 - 4 * (223-86/(3+3-4*(9-9+(5-4)))))+3":2212,
}

func TestCalculate(t *testing.T) {
	for k,v := range calcQsn {
		got := algo.Calculate(k)
		if got != v {
			t.Errorf("exp:%s,expect:%v,got:%v\n", k, v, got)
		}
	}
}
