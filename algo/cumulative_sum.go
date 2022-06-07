package algo

// 计算给定数组 arr 中累加和是 s 的数量
// 限制：要求 arr 中的数据都是 >= 0 的整数（如果有负数，则全局加上最小负数的相反数就都变成 >= 0 的了）

func Accumulate(arr []int, s int) int {
	if s < 0 {
		return 0
	}

	if s == 0 {
		return 1
	}

	return processAcc(arr, 0, s)
}

func AccumulateDp(arr []int, s int) int {
	if s < 0 {
		return 0
	}

	if s == 0 {
		return 1
	}

	return dpAcc(arr, s)
}

// 递归函数
// 能够算出数组 arr[i,len - 1] 范围中累加和等于 rest 的组合数量
func processAcc(arr []int, i int, rest int) int {
	// base case
	if rest < 0 {
		// 因为我们限制 arr 中必须都是 >= 0 的，所以不可能累加得到负数
		return 0
	}

	if rest == 0 {
		// 处理完了
		return 1
	}

	if i == len(arr) - 1 {
		// 游标到达右边界，此时形成的子数组只有一个元素
		if rest == arr[i] {
			return 1
		}

		return 0
	}

	// 普通情况，有两种分支
	// 分支1：使用第 i 个元素
	p1 := processAcc(arr, i + 1, rest - arr[i])
	// 分支2：不使用第 i 个元素
	p2 := processAcc(arr, i + 1, rest)

	return p1 + p2
}

// 将暴力尝试函数 processAcc 改成动态规划
func dpAcc(arr []int, sum int) int {
	N := len(arr)
	M := sum + 1

	// 生成 N*M 的数组
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	// 初始值：第一列
	for i := 0; i < N; i++ {
		dp[i][0] = 1
	}

	// 初始值：最后一行
	lst := N - 1
	for j := 1; j < M; j++ {
		if arr[lst] == j {
			dp[lst][j] = 1
		} else {
			dp[lst][j] = 0
		}
	}

	// 普通情况
	for i := N - 2; i >= 0; i-- {
		for j := 1; j < M; j++ {
			// 分支1：使用第 i 个元素
			// 只有 arr[i] 的值小于等于剩余累加数 j 时才能用
			if j - arr[i] >= 0 {
				dp[i][j] += dp[i + 1][j - arr[i]]//processAcc(arr, i + 1, rest - arr[i])
			}
			// 分支2：不使用第 i 个元素
			dp[i][j] += dp[i + 1][j]//processAcc(arr, i + 1, rest)
		}
	}

	return dp[0][M - 1]
}
