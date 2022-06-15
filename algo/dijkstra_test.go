package algo_test

import (
	"algo/algo"
	"strconv"
	"testing"
)

var pathVals = [][3]int {
	{1,2,3},
	{1,3,1},
	{1,4,2},
	{2,3,3},
	{2,5,4},
	{3,4,2},
	{3,5,2},
	{3,6,4},
	{4,6,4},
	{4,9,8},
	{5,6,3},
	{5,7,1},
	{5,8,5},
	{6,7,4},
	{6,9,5},
	{7,8,3},
	{8,9,2},
}

func TestGraph_ShortestPath(t *testing.T) {
	g := algo.NewGraph(9)
	for _,v := range pathVals {
		if err := g.Add(v[0], v[1], v[2]); err != nil {
			t.Errorf("get error:%v\n", err)
		}
	}

	p := g.ShortestPath(1, 9)

	if p == nil || p.Ids == nil {
		t.Errorf("get nil")
		return
	}

	s := ""
	for _,v := range p.Ids {
		s += strconv.Itoa(v)
	}
	expect := "135789"
	if s != expect || p.Dst != 9 || p.Val != 9 {
		t.Errorf("expect:%v,got:%v\n", expect, s)
	}
}
