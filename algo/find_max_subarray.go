package algo

const MinInt = ^int(^uint(0) >> 1)

// 求解最大子数组：该数组中连续的一段子数组之和最大化。只有同时包含正数和负数才有意义
// 可以用暴力遍历求解，其时间复杂度是 O(n^2)
// 此处使用分治算法求解，其时间复杂度是 O(nlgn)
// 分治算法思路：将数组 lst[i...j]分成两个相等的子数组 left[i...m]和 right[m+1...j]，然后分别求两个子数组的最大子数组
// 数组 lst[i...j]的最大子数组要么在 left 中，要么在 right 中，要么跨越 left 和 right。前两种情况可用分治算法递归求解，第三种
// 情况要单独求解。三种情况求解出来后，取最大值即是 lst[i...j]的最大子数组
func MaxSubArray(lst []int) [3]int {
	return findMaxSubArray(lst, 0, len(lst) - 1)
}

// 求解子切片的最大子数组
// 使用分治算法，将原切片分成左右两部分，分别求两部分的最大子数组
// return array [左边界, 右边界, 最大值]
func findMaxSubArray(lst []int, low int, high int) [3]int {
	if low == high {
		return [3]int{low, high, lst[low]}
	}

	mid := low + (high - low) / 2
	// 左切片最大子数组
	leftInfo := findMaxSubArray(lst, low, mid)
	// 右切片最大子数组
	rightInfo := findMaxSubArray(lst, mid + 1, high)
	// 中间（越过 mid）最大子数组
	midInfo := findMaxCrossSubArray(lst, low, mid, high)

	// 最大子数组必然是上面的之一
	lSum, mSum, rSum := leftInfo[2], midInfo[2], rightInfo[2]
	if lSum >= mSum && lSum >= rSum {
		return leftInfo
	} else if mSum >= lSum && mSum >= rSum {
		return midInfo
	} else {
		return rightInfo
	}
}

// 查找跨越中点 mid 的最大子数组
// 该最大子数组 = left[i...mid] + right[mid+1...j]，即左右贴着 mid 的最大子数组拼接成的（左右子数组最少元素是 1 个）
// 复杂度：O(n)
// return array [左边界, 右边界, 最大值]
func findMaxCrossSubArray(lst []int, low int, mid int, high int) [3]int {
	if low == mid && mid == high {
		return [3]int{low, high, lst[low]}
	}

	leftMax, rightMax := MinInt, MinInt

	// 左边求值
	leftIndex, leftSum := low, 0
	for i := mid; i >= low; i-- {
		leftSum += lst[i]
		if leftMax < leftSum {
			leftMax = leftSum
			leftIndex = i
		}
	}

	// 右边求值
	rightIndex, rightSum := high, 0
	for i := mid + 1; i <= high; i++ {
		rightSum += lst[i]
		if rightMax < rightSum {
			rightMax = rightSum
			rightIndex = i
		}
	}

	return [3]int{leftIndex, rightIndex, leftMax + rightMax}
}