package algo

// 司机总体收入最大化问题
// 现有 N*2 个司机，调度中心将他们平均分配到 A、B 两个区域
// 现有如下数据：arr = [[2,3], [1,4]]，表示第 1 个司机到 A 区域可以挣 2 块钱，到 B 区域可以挣 3 块钱；第 2 个司机到 A 区域可以挣 1 块钱，
// 到 B 区域可以挣 4 块钱
// 问如何调度能让所有司机整体收入最大，最大收入是多少

// 计算整体最大收入
// 返回最大收入是多少
func MaxIncome(drivers [][2]int) int {
	if len(drivers) & 1 != 0 {
		panic("司机数量必须是偶数")
	}

	half := len(drivers) >> 1

	return processMaxIncome(drivers, 0, half, half)
}

// 动态规划版本
func MaxIncomeDp(drivers [][2]int) int {
	if len(drivers) & 1 != 0 {
		panic("司机数量必须是偶数")
	}

	return dpMaxIncome(drivers)
}

// 暴力尝试
// 该函数能够计算 drivers[i,len - 1] 子数组中，放 xa 个到 A 区域，xb 个到 B 区域的分配策略下的最大整体收入
func processMaxIncome(drivers [][2]int, i int, xa int, xb int) int {
	// base case
	if xa == 0 && xb == 0 {
		// 都无法放入了
		return 0
	}

	curr := drivers[i]

	if i == len(drivers) - 1 {
		// 最后一个，要么放入 A 区域，要么放入 B 区域，然后结束
		if xa == 0 {
			// A 区域满了，只能放到 B 区域
			return curr[1]
		} else if xb == 0 {
			// B 区域满了，只能放到 A 区域
			return curr[0]
		}
	}

	// 普通情况，有两种选择（分支）
	sa, sb := 0, 0

	if xa > 0 {
		// 选择1：放到 A 区域
		sa = curr[0] + processMaxIncome(drivers, i + 1, xa - 1, xb)
	}

	if xb > 0 {
		// 选择2：放到 B 区域
		sb = curr[1] + processMaxIncome(drivers, i + 1, xa, xb - 1)
	}

	// 返回最大的
	if sa > sb {
		return sa
	} else {
		return sb
	}
}

// 将 processMaxIncome 改成动态规划
func dpMaxIncome(drivers [][2]int) int {
	l := len(drivers)
	X, Y, Z := l, (l >> 1) + 1, (l >> 1) + 1

	dp := make([][][]int, X)
	for i := 0; i < X; i++ {
		dp[i] = make([][]int, Y)
		for j := 0; j < Y; j++ {
			dp[i][j] = make([]int, Z)
		}
	}

	// 初始化值：最后一层
	lst := l - 1
	for z := 0; z < Z; z++ {
		// A 区域满了，只能放到 B 区域
		dp[lst][0][z] = drivers[lst][1]
	}
	for y := 0; y < Y; y++ {
		// B 区域满了，只能放到 A 区域
		dp[lst][y][0] = drivers[lst][0]
	}

	// 普通情况，有两种选择（分支）
	// x 方向要倒序遍历。最后一层（x == lst）上面已经处理了，下面从 lst - 1 层开始处理
	for x := lst - 1; x >= 0; x-- {
		curr := drivers[x]
		for y := 0; y < Y; y++ {
			for z := 0; z < Z; z++ {
				sa, sb := 0, 0

				if y > 0 {
					// 选择1：放到 A 区域
					sa = curr[0] + dp[x + 1][y - 1][z]//curr[0] + processMaxIncome(drivers, i + 1, xa - 1, xb)
				}

				if z > 0 {
					// 选择2：放到 B 区域
					sb = curr[1] + dp[x + 1][y][z - 1]//curr[1] + processMaxIncome(drivers, i + 1, xa, xb - 1)
				}

				// 取最大的
				if sa > sb {
					dp[x][y][z] = sa
				} else {
					dp[x][y][z] = sb
				}
			}
		}
	}

	return dp[0][Y - 1][Z - 1]
}
