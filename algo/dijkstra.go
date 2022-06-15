package algo

import (
	"container/heap"
	"errors"
)

// dijkstra 最短路径算法
// 单源最短路径算法
// 要求所有距离必须 >= 0
// dijkstra 是贪心算法，每次从已经走过的路径中选择最小路径所到达的那个目的地（顶点），从此处出发前往和它相邻的目标顶点（广度优先遍历）；周而复始
// 涉及到三个点：1. 贪心算法；2. 堆（实现贪心）；3. 广度优先遍历；

// 边
// 单向链表
type edgeNode struct {
	val int // 边权重（距离）
	dst int // 边的目标顶点id
	next *edgeNode
}

// 通过邻接表实现图
type Graph struct {
	// 数组的下标表示节点编号，值表示以该顶点作为源点形成的所有的边的链表
	nodes []*edgeNode
}

// 创建图
// cnt：指定节点数
// 节点编号必须是 1 ~ n 的数
func NewGraph(cnt int) *Graph {
	if cnt <= 0 {
		return nil
	}

	return &Graph{nodes: make([]*edgeNode, cnt+1, cnt+1)}
}

// 往图中添加边
// src：源顶点编号，n~n-1的正整数
// Dst：目标顶点编号
// distance：距离，正数（dijkstra 不支持负数权值。此处我们限制为正整数）
func (g *Graph) Add(src, dst, distance int) error {
	if distance < 0 {
		return errors.New("dist 必须是正数")
	}

	if src < 1 || src > len(g.nodes) || dst < 1 || dst > len(g.nodes) {
		return errors.New("id 不合法")
	}

	// 不能存在环路
	if g.searchEdge(dst, src) != nil {
		return errors.New("存在环路")
	}

	// 是否已经加入过了
	edge := g.searchEdge(src, dst)
	if edge != nil {
		edge.val = distance
		return nil
	}

	// 放到链表头部
	edge = &edgeNode{
		val: distance,
		dst: dst,
		next: g.nodes[src],
	}

	g.nodes[src] = edge

	return nil
}

func (g *Graph) searchEdge(src, dst int) *edgeNode {
	edge := g.nodes[src]

	for edge != nil {
		if edge.dst == dst {
			return edge
		}

		edge = edge.next
	}

	return nil
}

// 通过 dijkstra 算法计算 src 到 Dst 的最短路径
func (g *Graph) ShortestPath(src, dst int) *nPath {
	if src < 1 || src > len(g.nodes) || dst < 1 || dst > len(g.nodes) {
		return nil
	}

	// 小顶堆
	ph := &pathHeap{}
	heap.Init(ph)

	// 将 src 加入到堆中
	heap.Push(ph, &nPath{Ids: []int{src}, Val: 0, Dst: src})

	// 从堆中取出堆顶元素处理（取当前最短距离）
	cPath := heap.Pop(ph).(*nPath)
	for cPath != nil {
		// 找到了
		if cPath.Dst == dst {
			return cPath
		}

		// 计算该顶点到相邻顶点的路径并放入堆中
		node := g.nodes[cPath.Dst]
		for node != nil {
			// 注意：这里要创建新的切片，不能直接在 cPath.Ids 上用 append，否则它们底层共用同一个数组，就会导致数据混乱
			ids := make([]int, 0, len(cPath.Ids)+1)
			ids = append(ids, cPath.Ids...)
			heap.Push(ph, &nPath{Ids: append(ids, node.dst), Val: cPath.Val + node.val, Dst: node.dst})
			node = node.next
		}

		cPath = heap.Pop(ph).(*nPath)
	}

	return nil
}

// 路径
type nPath struct {
	Ids []int // 路径经过的顶点
	Val int   // 路径的权值
	Dst int   // 路径终点的顶点id，也就是 Ids 中最后一个元素
}

// 小顶堆
type pathHeap struct {
	pathes []*nPath
}

func (h pathHeap) Len() int {
	return len(h.pathes)
}

func (h pathHeap) Less(i, j int) bool {
	return h.pathes[i].Val < h.pathes[j].Val
}

func (h *pathHeap) Swap(i, j int) {
	h.pathes[i], h.pathes[j] = h.pathes[j], h.pathes[i]
}

func (h *pathHeap) Push(x interface{}) {
	h.pathes = append(h.pathes, x.(*nPath))
}

func (h *pathHeap) Pop() interface{} {
	l := len(h.pathes)
	v := h.pathes[l-1]
	h.pathes = h.pathes[:l-1]

	return v
}
