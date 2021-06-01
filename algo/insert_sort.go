package algo

// 插入排序
// 从小到大排序
// 原址排序
func InsertSort(lst []int) {
	// 从第二个元素开始遍历到最后一个元素
	for i, l := 1, len(lst); i < l; i++ {
		// 将第 i 个元素插入到前面子数组的适当位置
		// 先将该元素存到临时变量中，空出当前位置让子数组的元素可以往后移动
		curr := lst[i]
		j := i - 1
		for ; j >= 0; j-- {
			// 如果当前位置元素比 curr 大，则往后面移，否则跳出循环
			if lst[j] <= curr {
				break
			}

			// 往后移
			lst[j + 1] = lst[j]
		}

		// 插入到 j 的后面
		lst[j + 1] = curr
	}
}
