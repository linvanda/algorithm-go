package algo

// 求数组中最大累加和子数组
// 只有同时包含正数和负数才有意义

// 方法一：
// 诸如"子串"、"子数组"之类的问题，都可以采用如下的通用思想：
// 从左到右，依次求子串以 0、1、2、3...n 元素结尾的情况下的结果是多少
// 即问："以 0 结尾的是什么？以 1 结尾的是什么？..."
// 它基于这样的思考：如果能求得以 i 结尾的子串 S，我们可以根据 S 求得 i + 1 结尾的子串 S'
func MaxCumulativeSum1(lst []int) ([]int, int) {
	if len(lst) == 0  {
		return nil, 0
	}

	if len(lst) == 1 {
		return lst, lst[0]
	}

	// 最大累加和子数组左右边界，以及当前正在计算的累加和左边界起始处
	l, r, s := 0, 0, 0
	// 最大累加和
	max := lst[0]
	// 当前位置前面的最大累加和
	preSum := lst[0]

	// 从左到右，依次求出以 i 元素作为右边界的最大累加和子数组是什么
	// 第 i 个位置的最大累加和取决于 i - 1 的情况：
	// 1. 如果 i - 1 的最大累加和>=0，则 i 的直接将 lst[i] + preSum（一个数加上非负数总是不小于它自己）
	// 2. 如果 i - 1 的最大累加和<0，则 i 的最大累加和是 lst[i]（一个数加上负数会变小，所以不能加）
	// 由于 i 位置的最大累加和只依赖 i - 1的，所以只需要一个 int 变量记录即可（无需记录所有位置的最大累加和）
	i := 1
	for ; i < len(lst); i++ {
		if preSum >= 0 {
			preSum += lst[i]
		} else {
			// 前面的累加和是负数
			if max < preSum {
				// 记录下新的最大累加和数组信息
				l = s
				r = i - 1
				max = preSum
			}

			// 开启新的子数组
			s = i
			preSum = lst[i]
		}
	}

	// 再算一下
	if max < preSum {
		// 记录下新的最大累加和数组信息
		l = s
		r = i - 1
		max = preSum
	}

	return lst[l:r+1], max
}

const MinInt = ^int(^uint(0) >> 1)

// 方法二：分治法
// 此处使用分治算法求解，其时间复杂度是 O(nlgn)
// 分治算法思路：将数组 lst[i...j]分成两个相等的子数组 left[i...m]和 right[m+1...j]，然后分别求两个子数组的最大子数组
// 数组 lst[i...j]的最大子数组要么在 left 中，要么在 right 中，要么跨越 left 和 right。前两种情况可用分治算法递归求解，第三种
// 情况要单独求解。三种情况求解出来后，取最大值即是 lst[i...j]的最大子数组
func MaxCumulativeSum2(lst []int) ([]int, int) {
	if len(lst) == 0  {
		return nil, 0
	}

	if len(lst) == 1 {
		return lst, lst[0]
	}

	rst := findMaxSubArray(lst, 0, len(lst) - 1)

	return lst[rst[0]:rst[1]+1], rst[2]
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

