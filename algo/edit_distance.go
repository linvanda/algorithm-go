package algo

// 字符串 src 到 Dst 的最小编辑距离
// 字符串 src 的 i 位置字符要变成字符串 Dst 的 j 位置的字符，有 4 种可能的方式：
// 1. 保留，代价是 0（ src[i] == Dst[j]）
// 2. 删除，代价是 d
// 3. 插入，代价是 i
// 4. 替换，代价是 r
// 其中，在 src 的 i 字符处，如果做的决策是插入，则 i 原来的值会保留到后面，其他的决策都会消耗 i 原来的字符

// 操作代价
type EditCost struct{
	Retain int
	Del int
	Insert int
	Replace int
}

var defaultEdCost = EditCost{
	Retain: 0,
	Del: 1,
	Insert: 1,
	Replace: 1,
}

// 编辑代价
var EdCost = defaultEdCost

// 重置为默认值
func ResetEdCost() {
	EdCost = defaultEdCost
}

// 方法一：暴力尝试
func EditDistance1(src, dst string) int {
	return processEDist([]byte(src), []byte(dst), 0, 0)
}

// 方法二：动态规划
func EditDistance2(src, dst string) int {
	return dpEDist([]byte(src), []byte(dst))
}

// 暴力尝试
// 声明：该函数能够计算子串 src[i,len-1] -> Dst[j,len-1] 的最短距离
//（另一种方式是 i、j 表示处理 src[0,i]->Dst[0,j]）的最短距离（前缀法））
func processEDist(src, dst []byte, i, j int) int {
	// base case
	if i == len(src) && j == len(dst) {
		// src 和 Dst 同时空了，无需再做任何处理
		return 0
	}

	if j == len(dst) {
		// Dst 已经生成完了，此时 src 中还有需要处理的元素，则只能删除
		return (len(src) - i) * EdCost.Del
	}

	if i == len(src) {
		// src 已经用完了，此时 Dst 还没有完成，只能执行插入操作
		return (len(dst) - j) * EdCost.Insert
	}

	// 普通情况：src 和 Dst 都还有待处理字符

	// 情况1：当前两个字符相同
	if src[i] == dst[j] {
		// 两个字符相同，啥也不做，都往后移动一位
		return processEDist(src, dst, i + 1, j + 1)
	}

	// 情况2：两个字符不同，有三种选择

	// 选择1：删除 i 位置的字符，此时 i 往后移一位，j 不变
	// 本步骤产生删除代价
	c1 := EdCost.Del + processEDist(src, dst, i + 1, j)

	// 选择2：插入 j 位置需要的元素，此时 i 不动，j 往后移一位
	// 本步骤产生插入代价
	c2 := EdCost.Insert + processEDist(src, dst, i, j + 1)

	// 选择3：将 i 的字符替换成 j 的字符，此时 i、j 都往后移动一位
	// 本步骤产生替换代价
	c3 := EdCost.Replace + processEDist(src, dst, i + 1, j + 1)

	// 取三种情况代价最小的
	return minInt(c1, c2, c3)
}

// 将上面的暴力尝试改造成动态规划版本
// 样本对应模型
func dpEDist(src, dst []byte) int {
	// 生成一个 N*M 的数组，其 arr[i][j] 的含义是：从 src[i:]->Dst[j:] 的编辑距离
	// 初始值（递归的最里层）是：dp[N-1][M-1] = 0，它表示由空生成空，代价必然是 0（注意：N、M 都比原数组的长度大 1）
	// 最终要求解的值：dp[0][0]，它表示求解 src[0:]->Dst[0:] 的编辑距离
	// 求解方向是 右下 -> 左上
	ls,ld := len(src), len(dst)
	N, M := ls + 1, ld + 1
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	// 初始值
	dp[N-1][M-1] = 0

	// 最后一行：src 中的字符都用完了，只能采用插入策略
	for j := ld-1; j >= 0; j-- {
		dp[N-1][j] = (ld - j) * EdCost.Insert
	}

	// 最后一列：Dst 中要的字符都生成完了，src 中剩下的只能全删除
	for i := ls-1; i >= 0; i-- {
		dp[i][M-1] = (ls - i) * EdCost.Del
	}

	// 普通情况：
	for i := ls-1; i >= 0; i-- {
		for j := ld-1; j >= 0; j-- {
			// 情况1：当前两个字符相同
			if src[i] == dst[j] {
				// 两个字符相同，啥也不做，都往后移动一位
				dp[i][j] = dp[i+1][j+1]
				continue
			}

			// 情况2：两个字符不同，有三种选择

			// 选择1：删除 i 位置的字符，此时 i 往后移一位，j 不变
			// 本步骤产生删除代价
			c1 := EdCost.Del + dp[i+1][j]

			// 选择2：插入 j 位置需要的元素，此时 i 不动，j 往后移一位
			// 本步骤产生插入代价
			c2 := EdCost.Insert + dp[i][j+1]

			// 选择3：将 i 的字符替换成 j 的字符，此时 i、j 都往后移动一位
			// 本步骤产生替换代价
			c3 := EdCost.Replace + dp[i+1][j+1]

			dp[i][j] = minInt(c1, c2, c3)
		}
	}

	return dp[0][0]
}
