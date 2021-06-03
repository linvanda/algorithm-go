package algo

const MaxInt = int(^uint(0) >> 1)

// 归并排序
// 分治算法
// 非原址排序
// 分为归和并两部分：
//     归：用分治算法，将大问题化成多个子问题，递归解决
//	   并：对子问题解合并，形成大问题的解
func MergeSort(lst []int) {
	l := len(lst)
	if l <= 1 {
		return
	}

	sort(lst, 0, l - 1)
}

// 对 lst[start...end]的子切片执行归并排序
func sort(lst []int, start int, end int) {
	if start >= end {
		return
	}

	// 一分为二，对左右子切片分别排序
	middle := start + (end - start) / 2
	// 左排序
	sort(lst, start, middle)
	// 右排序
	sort(lst, middle + 1, end)

	// 左右子切片排序完后，对左右子切片合并
	merge(lst, start, middle, end)
}

// 将 lst[p...q]和 lst[q+1...k]两个子数组合并，要求合并后的子数组 lst[p...k]是已经排好序的
// q 两边的子数组已经排好序
func merge(lst []int, p int, q int, k int) {
	if p >= q && q >= k {
		return
	}

	// 创建两个临时slice 存放左右两个子数组。注意此处分配的 slice 大小已经包含哨兵的位置
	lLen, rLen := q - p + 2, k - q + 1
	left := make([]int, lLen)
	right := make([]int, rLen)

	// 分别复制到临时切片
	for i := p; i <= q; i++ {
		left[i - p] = lst[i]
	}

	for i := q + 1; i <= k; i++ {
		right[i - q - 1] = lst[i]
	}

	// 为防止每次循环都要去判断切片是否越界，我们在两个数组最后加上哨兵，其值为 int 最大值（保证切片中的任何值不会大于该值)
	left[lLen - 1] = MaxInt
	right[rLen - 1] = MaxInt

	// 将两个临时切片的数据由小到大写回原切片
	// 每次循环写回一个元素，一共需要执行 k - p + 1 次循环
	l, r := 0, 0
	for i := p; i <= k; i++ {
		// 从两个临时切片取小的那个写入 lst
		if left[l] <= right[r] {
			lst[i] = left[l]
			l++
		} else {
			lst[i] = right[r]
			r++
		}
	}
}