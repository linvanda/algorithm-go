package algo_test

import (
	"algo/algo"
	"testing"
)

func TestMaxIncome(t *testing.T) {
	drivers := [][2]int{{1, 2}, {3, 1}, {4, 5}, {8,5}, {1, 2}, {3, 1}, {7, 5}, {2,9}}
	expect := 39

	got := algo.MaxIncome(drivers)
	if got != expect {
		t.Errorf("drivers:%v,expecte:%d,got:%d\n", drivers, expect, got)
	}
}

func TestMaxIncomeDp(t *testing.T) {
	drivers := [][2]int{{1, 2}, {3, 1}, {4, 5}, {8,5}, {1, 2}, {3, 1}, {7, 5}, {2,9}}
	expect := 39

	got := algo.MaxIncomeDp(drivers)
	if got != expect {
		t.Errorf("drivers:%v,expecte:%d,got:%d\n", drivers, expect, got)
	}
}

func BenchmarkMaxIncome(b *testing.B) {
	drivers := [][2]int{{1, 2}, {3, 1}, {4, 5}, {8,5},{1, 2}, {3, 1}, {4, 5}, {8,5},{1, 2}, {3, 1}, {4, 5}, {8,5}, {3, 1}, {4, 5}, {8,5},{1, 2}, {3, 1}, {4, 5}, {8,5}, {8,5}}
	for i := 0; i < b.N; i++ {
		algo.MaxIncome(drivers)
	}
}

func BenchmarkMaxIncomeDp(b *testing.B) {
	drivers := [][2]int{{1, 2}, {3, 1}, {4, 5}, {8,5},{1, 2}, {3, 1}, {4, 5}, {8,5},{1, 2}, {3, 1}, {4, 5}, {8,5}, {3, 1}, {4, 5}, {8,5},{1, 2}, {3, 1}, {4, 5}, {8,5}, {8,5}}
	for i := 0; i < b.N; i++ {
		algo.MaxIncomeDp(drivers)
	}
}