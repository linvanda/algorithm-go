package algo_test

import (
	"algo/algo"
	"testing"
)

func TestHorseJump(t *testing.T) {
	p := algo.Pos{7, 7}
	step := 10
	m := algo.HorseJump(p, step)
	if m != 515813 {
		t.Errorf("step %d to pos:%v,real:%d,get:%d\n", step, p, 515813, m)
	}
}

func TestHorseJumpDp(t *testing.T) {
	p := algo.Pos{7, 7}
	step := 10
	m := algo.HorseJumpDp(p, step)
	if m != 515813 {
		t.Errorf("step %d to pos:%v,real:%d,get:%d\n", step, p, 515813, m)
	}
}

func BenchmarkHorseJump(b *testing.B) {
	p := algo.Pos{7, 7}
	step := 10
	for i := 0; i < b.N; i++ {
		algo.HorseJump(p, step)
	}
}

func BenchmarkHorseJumpDp(b *testing.B) {
	p := algo.Pos{7, 7}
	step := 10
	for i := 0; i < b.N; i++ {
		algo.HorseJumpDp(p, step)
	}
}