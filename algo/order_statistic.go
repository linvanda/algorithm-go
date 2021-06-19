package algo

// 求互异数组中第 i 个顺序统计量，即第 i 小的元素
// 最小值和最大值（第一个和最后一个）可以通过 n-1次比较获得
// 可以通过随机选择算法在线性时间内求得第 i 个顺序统计量（求法类似快排）
// 注意：为了方便，这里 i 从 0 开始
func OrderStatistic(lst []int, i int) int {
	return randomizedSelect(lst, 0, len(lst) - 1, i)
}

// 从 lst 数组中找出 l...r 之间的第 i 小的元素
// 思想：使用和快排一样的分治算法，对 l...r 按照标的 q 进行partition后，q左边的都会比 q 小，如果 q 的位置是 j，则说明 q 是第 j 小的元素，
//		否则比较 j 和 i 的大小以决定是继续对 q 左边的元素 partition 还是对右边 partition。
// 和快排不同，此处只需要对一半的元素进行 partition，其时间复杂度可以做到 O(n)
func randomizedSelect(lst []int, l int, r int, i int) int {
	// 如果 l == r，说明已经到了最后，无法继续 partition 了，此时一定就是 i，直接返回
	if l == r {
		return lst[l]
	}

	// 执行 partition，返回标的位置 j，这个就是第 j 小的值
	j := partition(lst, l, r)
	// 如果 j == i，说明找到了
	if j == i {
		return lst[j]
	} else if j > i {
		// 继续对左边进行 select
		return randomizedSelect(lst, l, j - 1, i)
	} else {
		// 继续对右边 select
		return randomizedSelect(lst, j + 1, r, i)
	}
}
