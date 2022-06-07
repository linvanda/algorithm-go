package algo

import "math"

// 最长回文子序列
// 注意：子序列可以不连续
// 如：ab1238f321add 的最长回文子序列是 1238321

// 暴力尝试法 1
func LongestPalindrome1(s []byte) int {
	if s == nil || len(s) == 0 {
		return 0
	}

	// 最长回文子序列就是字符串和它的反串的最长公共子序列
	srev := rev(s)

	return process1(s, srev, len(s) - 1, len(s) - 1)
}

// 暴力尝试法 2
// 和 暴力尝试法 1 本质是一样的
func LongestPalindrome2(s []byte) int {
	if s == nil || len(s) == 0 {
		return 0
	}

	return process2(s, 0, len(s) - 1)
}

// 暴力尝试法 3
func LongestPalindrome3(s []byte) int {
	if s == nil || len(s) == 0 {
		return 0
	}

	return process3(s, 0, len(s) - 1)
}

// 动态规划版本1
func LongestPalindromeDP1(s []byte) int {
	if s == nil || len(s) == 0 {
		return 0
	}

	// 最长回文子序列就是字符串和它的反串的最长公共子序列
	srev := rev(s)

	return dp1(s, srev)
}

// 动态规划版本3
func LongestPalindromeDP3(s []byte) int {
	if s == nil || len(s) == 0 {
		return 0
	}

	return dp3(s)
}

// 求 s1、s2 的最长公共子序列
func LongestCommonSubsequence(s1, s2 []byte) int {
	if s1 == nil || s2 == nil || len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	return process1(s1, s2, len(s1) - 1, len(s2) - 1)
}

// 求 s1、s2 的最长公共子序列：动态规划版本
func LongestCommonSubsequenceDP(s1, s2 []byte) int {
	if s1 == nil || s2 == nil || len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	return dp1(s1, s2)
}

func rev(b []byte) []byte {
	rb := make([]byte, len(b))
	l := len(b) - 1
	for i := l; i >= 0; i-- {
		rb[l - i] = b[i]
	}

	return rb
}

// 暴力尝试
// 该函数能求解两个字符序列 s1[0...i]、s2[0...j] 的最长公共子序列
// 这里用的是样本对应模型，s1、s2 两个样本，从右到左，每次（每步）有三种情况（三种选择）
func process1(s1, s2 []byte, i, j int) int {
	// 终止条件1
	if i == 0 && j == 0 {
		// 都到第一个了，比较是否相同
		if s1[0] == s2[0] {
			return 1
		}
		return 0
	}

	// 终止条件2:i 到了 0, j 没到
	if i == 0 {
		// 如果两者当前值相同，则直接返回 1，否则 j 继续行进
		if s1[0] == s2[j] {
			return 1
		}

		// 继续行进
		return process1(s1, s2, 0, j - 1)
	}

	// 终止条件3: j 到了 0，i 没到，跟条件2 一样
	if j == 0 {
		if s1[i] == s2[0] {
			return 1
		}

		return process1(s1, s2, i - 1, 0)
	}

	// i、j 都没到 0，此时有三种分支情况（三种选择）
	// 因此要暴力尝试（递归）所有三种情况

	// 选择1：用 i 不用 j，让 j 往前走
	p1 := process1(s1, s2, i, j - 1)
	// 选择2：用 j 不用 i，让 j 往前走
	p2 := process1(s1, s2, i - 1, j)
	// 选择3：用 i 和 j，然后同时往前走
	p3 := process1(s1, s2, i - 1, j - 1)
	if s1[i] == s2[j] {
		p3 += 1
	}

	return int(math.Max(float64(p1), math.Max(float64(p2), float64(p3))))
}

// 计算回文子序列的第二种暴力尝试模型
// 该递归函数能够得到 s 中 [l, len - 1] 和 [0, r] 两个子串的最长回文子序列
func process2(s []byte, l, r int) int {
	// base case
	// 条件1：l、r 同时到达边界，此时[l, len - 1] 和 [0, r] 两个子串都各持有一个字符
	if l == len(s) - 1 && r == 0 {
		if s[l] == s[r] {
			return 1
		}

		return 0
	}

	// 条件2：l 到达右边界，r 没有到达左边界
	if l == len(s) - 1 {
		if s[l] == s[r] {
			return 1
		}

		// 此步骤两者不相等，让 r 继续行进
		return process2(s, l, r - 1)
	}

	// 条件3：r 到达左边界，l 没有到达右边界
	if r == 0 {
		if s[l] == s[r] {
			return 1
		}

		// 此步骤两者不相等，l 继续行进
		return process2(s, l + 1, r)
	}

	// l 和 r 都没有到达边界，分三种情况
	// 情况一：此步骤不考虑 l，考虑 r（让 l 往右行进）
	p1 := process2(s, l + 1, r)
	// 情况二：此步骤考虑 l，不考虑 r（让 r 往左行进）
	p2 := process2(s, l, r - 1)
	// 情况三：此步骤考虑 l 和 r，比较后将 l、r 同时行进
	p3 := process2(s, l + 1, r - 1)
	if s[l] == s[r] {
		p3 += 1
	}

	return int(math.Max(float64(p1), math.Max(float64(p2), float64(p3))))
}

// 第三种模式：范围尝试
// 该函数能够计算出 s[l...r] 子串的最长回文子序列长度
func process3(s []byte, l, r int) int {
	if l > r {
		return 0
	}

	// 只有一个元素
	if l == r {
		return 1
	}

	// 普通情况，有三种可能

	// 情况1：本步骤可能使用 l，不使用 r（让 r 继续行进）
	p1 := process3(s, l, r - 1)
	// 情况2：本步骤不使用 l，可能使用 r（让 l 继续行进）
	p2 := process3(s, l + 1, r)
	// 情况3：本步骤使用 l 和 r，然后两者同时行进
	p3 := process3(s, l + 1, r - 1)
	if s[l] == s[r] {
		p3 += 2// 注意回文长度要加 2
	}

	return int(math.Max(float64(p1), math.Max(float64(p2), float64(p3))))
}

// 将 process3 改写成动态规划：l 作为行，r 作为列
func dp3(s []byte) int {
	arr := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = make([]int, len(s))
	}

	L := len(s) - 1

	// 注意：行要倒着走，这样才能依据初始行的值推导出新行的值
	// 注意：这里可以优化：将特殊条件的值（初始化值）先单独 for 循环填了（对角线上的值），然后此处单独循环处理一半情况，这样此处的
	// 双循环就不用做 i == j 的判断了
	for i := L; i >= 0; i-- {
		// 只需要跑一半就行了（二维数组对角线左下角全部是 0，无需处理）
		for j := i; j <= L; j++ {
			if i == j {
				// 对角线上是 1
				arr[i][j] = 1
			} else {
				// 一般情况
				// 情况1：本步骤可能使用 l，不使用 r（让 r 继续行进）
				p1 := arr[i][j - 1]
				// 情况2：本步骤不使用 l，可能使用 r（让 l 继续行进）
				p2 := arr[i + 1][j]
				// 情况3：本步骤使用 l 和 r，然后两者同时行进
				p3 := arr[i + 1][j - 1]
				if s[i] == s[j] {
					p3 += 2// 注意回文长度要加 2
				}

				// 可以采用严格位置优化上面的逻辑，p3 不是必须计算的，仅在 s[i] == s[j] 才需要
				//p := maxInt(arr[i][j - 1], arr[i + 1][j])
				//if s[i] == s[j] {
				//	p = maxInt(2 + arr[i + 1][j - 1], p)
				//}

				arr[i][j] = maxInt(p1, maxInt(p2, p3))
			}
		}
	}

	return arr[0][L]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// 将 process1 改造成动态规划
// 改造步骤：
// 	1. 从暴力尝试方法 process1 的入参中找出变量：i、j，有两个变量，说明动态规划表是一个二维数组
//  2. 确定变量 i、j 的范围，即动态规划表二维数组的范围（大小）：i：[0, len(s1))，j：[0, len(s2))
//  3. 根据步骤二，用 s1 作为行，s2 作为列构成动态规划表
//  4. 根据暴力尝试的最终调用参数确定最终要返回表中哪个记录：arr[N][M]（process1(s, srev, len(s) - 1, len(s) - 1)）
//	5. 根据暴力尝试的 base case 确定表中初始值
//	6. 根据暴力尝试的一般情况逻辑翻译出动态规划的其他表项
func dp1(s1, s2 []byte) int {
	// 动态规划表
	arr := make([][]int, len(s1))
	// 初始化第二维
	for i := 0; i < len(s1); i++ {
		arr[i] = make([]int, len(s2))
	}

	// 初始值
	// 对终止条件1的翻译
	if s1[0] == s2[0] {
		arr[0][0] = 1
	} else {
		arr[0][0] = 0
	}

	N := len(s1) - 1
	M := len(s2) - 1

	// 对终止条件2的翻译（第 0 行的数据情况）
	// 第 0 列上面已经处理了，从第一列开始
	for j := 1; j <= M; j++ {
		if s1[0] == s2[j] {
			arr[0][j] = 1
		} else {
			arr[0][j] = arr[0][j - 1]
		}
	}

	// 对终止条件3的翻译：第 0 列的数据
	// 第 0 行已经处理过了，从第一行开始处理
	for i := 1; i <= N; i++ {
		if s1[i] == s2[0] {
			arr[i][0] = 1
		} else {
			arr[i][0] = arr[i - 1][0]
		}
	}

	// i、j 都没到 0，此时有三种分支情况（三种选择）
	// 即从第 1 行 第 1 列开始的情况
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// 选择1：用 i 不用 j，让 j 往前走
			p1 := arr[i][j - 1]
			// 选择2：用 j 不用 i，让 j 往前走
			p2 := arr[i - 1][j]
			// 选择3：用 i 和 j，然后同时往前走
			p3 := arr[i - 1][j - 1]
			if s1[i] == s2[j] {
				p3 += 1
			}

			arr[i][j] = int(math.Max(float64(p1), math.Max(float64(p2), float64(p3))))
		}
	}

	return arr[N][M]
}
