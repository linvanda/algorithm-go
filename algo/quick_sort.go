package algo

import (
	"math/rand"
	"time"
)

// 快速排序
// 原址排序
// 时间复杂度：最坏复杂度:O(n^2)，期望复杂度：O(nlgn)。因其常数因子很小，而且可以通过随机选取种子来规避最坏情况，因而在实际应用中往往是最好的排序算法
// 使用分治思想，将数组 lst 分成 left 和 right 两部分（partition），left 中所有元素都小于等于 v，right 中元素都大于 v
// 关键算法在 partition
func QuickSort(lst []int) {
	quickSort(lst, 0, len(lst) - 1)
}

// 对lst[l...r]范围的子数组执行快排
func quickSort(lst []int, l int, r int) {
	// 终止条件：l >= r
	if l >= r {
		return
	}

	// partition
	p := partition(lst, l, r)
	// 递归式：左子数组继续排序
	quickSort(lst, l, p - 1)
	// 递归式：右子数组继续排序
	quickSort(lst, p + 1, r)
}

// 对 [l...r] 范围的子数组执行 partition，返回选取的值 v 被插入的位置 p
// 思路：从左到右设立两个游标：p、q，使得 p 以及其左边的值都小于等于标的 v，p + 1 和 q 之间的值都大于 v。然后将 q 循环往后移动，每遇到一个
// 	    小于等于 v 的，便和 p 后面的那个元素交换，并将 p 往后移动一个位置。
func partition(lst []int, l int, r int) int {
	size := r - l + 1
	// 先随机获取 v
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	flag := int(rd.Float64() * 1000) % size + l
	// 将位置 flag 和最后一个元素互换，让标的在最后面
	lst[flag], lst[r] = lst[r], lst[flag]

	// p 游标的初始化位置是 l - 1，表示没有哪个元素切确地比标的小
	p := l - 1
	v := lst[r]
	for q := l; q < r; q++ {
		if lst[q] <= v {
			// 小于等于标的，则将 p 游标往后移动一个，然后交换，使得 p 所指向的位置是最后一个确切地小于等于 v 的元素
			p++
			// 交换
			if p < q {
				// 当 p < q 时才需要交换（此时 p 指向的元素一定比 v 大）
				// 当 p == q 时，说明 q 此时指向的就是小于等于 v 的元素
				// p 不可能大于 q（因为每次循环，p 最多加 1）
				lst[p], lst[q] = lst[q], lst[p]
			}
		}
	}

	// q 循环结束（此时 q 指向 r - 1 的位置），将 v（最后一个元素）和 p + 1 交换（此时如果 p == q，说明所有元素都小于等于 v）
	p++
	if p < r {
		lst[p], lst[r] = lst[r], lst[p]
	}

	return p
}
