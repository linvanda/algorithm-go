package algo

// 求解最大非重复子串问题

// 使用滑动窗口
// 滑动窗口：通过left、right 两个指针，各自在一定条件下往右边滑动，直到 right 指针到达末端，此时left 和 right 之间的数据就是我们要的解
// 算法：每次将 right 右移时，先看看将要移动的那个位置的元素 a 有没有曾经出现过，没有出现过，则正常右移，否则将 left 向后移动一位
func LongestUnrepeatedSubStr1(str []byte) []byte {
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
		if !ok || pos == -1 {
			// 之前没有出现过该元素，将 right 指向该元素，并将位置记录到 positions 中
			right++
			positions[ele] = right
		} else {
			// 之前出现过该元素，将 left 右移到 pos 右边，移动前先把位置记录下来，记录目前为止最大子串
			if max[1] - max[0] < right - left {
				max[0] = left
				max[1] = right
			}

			// 要将途中经过的字符个数都剔除掉
			n := pos - left + 1// left 要到的新位置（偏移量）
			for i := 0; i < n; i++ {
				positions[str[left + i]] = -1
			}

			left += n
		}
	}

	// 再做一次赋值
	if max[1] - max[0] < right - left {
		max[0] = left
		max[1] = right
	}

	return str[max[0]:max[1] + 1]
}

// 方法二
// 诸如"子串"、"子数组"之类的问题，都可以采用如下的通用思想：
// 从左到右，依次求子串以 0、1、2、3...n 元素结尾的情况下的结果是多少
// 即问："以 0 结尾的是什么？以 1 结尾的是什么？..."
// 它基于这样的思考：如果能求得以 i 结尾的子串 S，我们可以根据 S 求得 i + 1 结尾的子串 S'
// 这里有动态规划的影子：根据前面的结果滚动求出后面的结果
// 针对本题：依次求以 0、1、2、3...n 元素结尾的最长不重复子串是什么
// 第 i 位置的结果依赖于（且仅依赖于）第 i - 1 位置的结果，如果第 i - 1 位置最长不重复子串长度是 4，则以 i 结尾的最长不重复子串最长是 4 + 1 = 5（第 i 元素
// 没有在前者子串中出现过），最短则取决于第 i 元素在前者子串中出现的位置
func LongestUnrepeatedSubStr2(str []byte) []byte {
	if len(str) == 0 {
		return nil
	}

	// 记录最长子串的起止位置
	l, r := 0, 0
	// 当前考察的元素（i）前面（i - 1）位置的最长子串长度
	preLong := 0
	// 字符最后一次出现的位置
	var pos [256]int
	// 初始化为 -1
	for i := 0; i < 256; i++ {
		pos[i] = -1
	}

	i := 0
	for ; i < len(str); i++ {
		p := pos[str[i]]
		if p == -1 || p < i - preLong {
			// 该字符从来没有出现过，或者出现在 preLong 范围之外
			preLong++
		} else {
			// 该字符出现在 preLong 范围之内
			// 先记录 l、r
			if r - l + 1 < preLong {
				l = i - preLong
				r = i - 1
			}

			// 修剪 preLong
			preLong = i - p
		}

		pos[str[i]] = i
	}

	// 再处理一次 l、r
	if r - l + 1 < preLong {
		l = i - preLong
		r = i - 1
	}

	return str[l:r+1]
}
