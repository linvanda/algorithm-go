package algo_test

import (
	"algo/algo"
	"testing"
)

var bstreeVals = [][]int{
	{50, 30, 80, 20, 40, 70, 90, 10, 25, 85, 75, 65},// 没有重复的
	{50, 50, 30, 30, 80, 20, 80, 40, 70, 90, 40, 40, 40, 10, 25, 85, 75, 65},// 有很多重复的
	{1,2,3,4,5,6,7,8,9,10,11},// 一条线下来的
	{11,10,9,8,7,6,5,4,3,2,1},// 一条线下来的
	{11,10,9,8,7,6,5,4,3,2,100,110},
	{11,10,9,8,7,6,5,4,3,2,100,90},
	{11,10,9,8,7,6,10,5,4,3,2,100,90},
}

var preorderResults = [][]int {
	{50,30,20,10,25,40,80,70,65,75,90,85},
}

var middleorderResults = [][]int{
	{10,20,25,30,40,50,65,70,75,80,85,90},
}

var postorderResults = [][]int{
	{10,25,20,40,30,65,75,70,85,90,80,50},
}

func TestBSTree_Add(t *testing.T) {
	for _, bv := range bstreeVals {
		bstree := algo.NewBSTree()

		if bstree.Count() != 0 {
			t.Errorf("count of empty bstree must 0,got %d\n", bstree.Count())
		}

		for _, v := range bv {
			bstree.Add(v)
		}

		if bstree.Count() != len(bv) {
			t.Errorf("bstree count error:expect:%d,got:%d\n", len(bv), bstree.Count())
		}

		// 查找元素
		for _, v := range bv {
			node := bstree.Find(v)
			if node == nil || node.Val() != v {
				t.Errorf("node error,expect val:%d,got node:%v\n", v, node)
			}
		}
	}
}

func TestBSTree_Del(t *testing.T) {
	for _, bv := range bstreeVals {
		bstree := algo.NewBSTree()

		static := make(map[int]int)
		for _, v := range bv {
			bstree.Add(v)

			static[v] += 1
		}

		for _, v := range bv {
			bstree.Del(v)
			node := bstree.Find(v)
			if static[v] == 0 && node != nil {
				t.Errorf("del node error,got:%v\n", node.Val())
			}
		}

		if bstree.Count() != 0 {
			t.Errorf("bstree count error after delete,expect 0,got:%d\n", bstree.Count())
		}

		// 重新加入进去
		for _, v := range bv {
			bstree.Add(v)
		}

		if bstree.Count() != len(bv) {
			t.Errorf("bstree val %v readd error.expect:%d,got:%d\n", bv, len(bv), bstree.Count())
		}
	}
}

func TestBSTree_Preorder(t *testing.T) {
	// 暂时只检测第一个
	for bi, bv := range bstreeVals {
		bstree := algo.NewBSTree()

		for _, v := range bv {
			bstree.Add(v)
		}

		prevs := bstree.Preorder()
		prevsVal := make([]int, len(prevs))
		for i, v := range prevs {
			prevsVal[i] = v.Val()
		}

		for i, v := range preorderResults[bi] {
			if v != prevsVal[i] {
				t.Errorf("arr:%v, preorder fail.expect:%v, got:%v\n", bv, preorderResults[bi], prevsVal)
			}
		}

		break
	}
}

func TestBSTree_MiddleOrder(t *testing.T) {
	// 暂时只检测第一个
	for bi, bv := range bstreeVals {
		bstree := algo.NewBSTree()

		for _, v := range bv {
			bstree.Add(v)
		}

		middles := bstree.MiddleOrder()
		middlesVal := make([]int, len(middles))
		for i, v := range middles {
			middlesVal[i] = v.Val()
		}

		for i, v := range middleorderResults[bi] {
			if v != middlesVal[i] {
				t.Errorf("arr:%v, middleorder fail.expect:%v, got:%v\n", bv, middleorderResults[bi], middlesVal)
			}
		}

		break
	}
}

func TestBSTree_Postorder(t *testing.T) {
	// 暂时只检测第一个
	for bi, bv := range bstreeVals {
		bstree := algo.NewBSTree()

		for _, v := range bv {
			bstree.Add(v)
		}

		posts := bstree.Postorder()
		postsVal := make([]int, len(posts))
		for i, v := range posts {
			postsVal[i] = v.Val()
		}

		for i, v := range postorderResults[bi] {
			if v != postsVal[i] {
				t.Errorf("arr:%v, postorder fail.expect:%v, got:%v\n", bv, postorderResults[bi], postsVal)
			}
		}

		break
	}
}