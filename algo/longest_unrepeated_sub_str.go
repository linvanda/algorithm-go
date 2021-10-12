package algo

// 使用滑动窗口求解最大非重复子串问题
// 滑动窗口：通过left、right 两个指针，各自在一定条件下往右边滑动，直到 right 指针到达末端，此时left 和 right 之间的数据就是我们要的解
// 算法：每次将 right 右移时，先看看将要移动的那个位置的元素 a 有没有曾经出现过，没有出现过，则正常右移，否则将 left 向后移动一位
// 注意：此处假设 str 中都是 asc 码
func GetLongestUnrepeatedSubStr(str string) string {
	if len(str) < 2 {
		return str
	}

	left, right, l := 0, -1, len(str)
	// 记录每个字符第一次出现的位置
	positions := make(map[uint8]int)

	max := [2]int{0, -1}

	// 注意：right 整体向左偏移了一位，为了 for 里面计算 right + 1
	for right < l - 1 {
		// right 的下一个元素
		ele := str[right + 1]
		pos, ok := positions[ele]
		if !ok || pos < 1 {
			// 之前没有出现过该元素，将 right 指向该元素，并将位置记录到 positions 中
			right++
			positions[ele] = 1
		} else {
			// 之前出现过该元素，将 left 右移一位，移动前先把位置记录下来，记录目前为止最大子串
			if max[1] - max[0] < right - left {
				max[0] = left
				max[1] = right
			}
			positions[str[left]] -= 1
			left++
		}
	}

	// 再做一次赋值
	if max[1] - max[0] < right - left {
		max[0] = left
		max[1] = right
	}

	return str[max[0]:max[1] + 1]
}
