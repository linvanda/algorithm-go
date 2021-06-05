package algo

import "strconv"

// 基数排序
// 时间复杂度：O(n)
// 非原址排序
// 稳定排序
// 思想：给定 n 个 d 位数（如手机号码）,从低位到高位（从右到左）一位一位排序
// 这和通常思维是相反的，通常思维是从高位到低位排序，但对于计算机来说，如果采用从高到低的排序方式，在处理低位数排序的时候还要兼顾高位数，会比较复杂
// 该排序的循环不变量：在对第 i 位进行排序前，i + 1位以及后面的（低位）已经排好序了。当 i == 0 排序完毕后，整个就排好序了
// 循环不变式成立的关键是对单个位置进行排序的算法是稳定算法，这样保证在第 i 位的两个值相等时，排序后两个串的位置不变，保证了 i + 1 位的顺序正确
// 此处假定字符串中都是数字（非负整数）
func RadixSort(lst []string, size int) []string {
	for i := size - 1; i >= 0; i-- {
		// 对第 i 位排序
		lst = countingSort(lst, 9, i)
	}

	return lst
}

func countingSort(lst []string, max int, index int) []string {
	// 创建一个 max + 1 大小的切片，其值初始化为 0
	c := make([]int, max + 1)

	// 遍历 lst，计算每个元素出现的次数
	for i := 0; i < len(lst); i++ {
		item, _ := strconv.ParseInt(lst[i][index:index + 1], 10, 64)
		c[int(item)] += 1
	}

	// 计算元素位置：后面元素 x 的左位置等于前面元素位置+1，右位置等于前面元素位置+m，m 为 x 出现的次数
	for i := 1; i <= max; i++ {
		c[i] += c[i - 1]
	}

	rst := make([]string, len(lst))
	// 从右往左遍历 lst 元素，将其放入正确的位置
	// 注意必须从右往左，因为 c 中记录的是元素最右边的位置（如果有多个元素的话），这样才能保证排序的稳定性
	for i := len(lst) - 1; i >= 0; i-- {
		item, _ := strconv.ParseInt(lst[i][index:index + 1], 10, 64)
		rst[c[item] - 1] = lst[i]
		// 完成一个了，需要从里面减去，接下来的（左边的）该元素会紧贴着此位置左边的位置放
		c[item] -= 1
	}

	return rst
}