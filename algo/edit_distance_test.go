package algo_test

import (
	"algo/algo"
	"testing"
)

var dstArr = [][2]string {
	{"abc", "ac"},
	{"", ""},
	{"abc", "abc"},
	{"abc", ""},
	{"", "abc"},
	{"abcdef", "abdh"},
}

// 使用默认编辑代价的期望结果
var defaultDstAns = []int {1, 0, 0, 3, 3, 3}

var benchDstArr = [2]string {"abcdef87dava5rfasa", "abdh087fwe0oihvah"}

func TestEditDistance1WithDefaultCost(t *testing.T) {
	algo.ResetEdCost()

	for i, v := range dstArr {
		got := algo.EditDistance1(v[0], v[1])
		if got != defaultDstAns[i] {
			t.Errorf("src:%s,dst:%s,expect:%d,got:%d\n", v[0], v[1], defaultDstAns[i], got)
		}
	}
}

func TestEditDistance2WithDefaultCost(t *testing.T) {
	algo.ResetEdCost()

	for i, v := range dstArr {
		got := algo.EditDistance2(v[0], v[1])
		if got != defaultDstAns[i] {
			t.Errorf("src:%s,dst:%s,expect:%d,got:%d\n", v[0], v[1], defaultDstAns[i], got)
		}
	}
}

func BenchmarkEditDistance1(b *testing.B) {
	algo.ResetEdCost()

	for i := 0; i < b.N; i++ {
		algo.EditDistance1(benchDstArr[0], benchDstArr[1])
	}
}

func BenchmarkEditDistance2(b *testing.B) {
	algo.ResetEdCost()

	for i := 0; i < b.N; i++ {
		algo.EditDistance2(benchDstArr[0], benchDstArr[1])
	}
}