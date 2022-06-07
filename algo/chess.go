package algo

type Pos struct {
	X int
	Y int
}

// 将象棋棋盘放在第一象限（左下角在 0,0 位置），在 x 轴 上有 9 条竖线（包括 y 轴），在 y 轴上有 10 条线（包括 x 轴）
// 对应在 x 轴范围是 [0, 8]，y 轴范围是 [0, 9]
// 现在将马放在原点（0,0）位置，给定一个点 p(x,y)，要求经过 s 步骤走到点 p，问一共有多少种方法

func HorseJump(pos Pos, step int) int {
	return chessProcess1(pos, Pos{0, 0}, step)
}

// 动态规划版本
func HorseJumpDp(pos Pos, step int) int {
	return chessDp(pos, step)
}

// 暴力尝试函数
// 该函数能够算出在点 a(x,y) 经过 z 步到达点 p 的方法数量
func chessProcess1(p, a Pos, z int) int {
	// 先判断位置 a 的合法性
	if a.X < 0 || a.X > 8 || a.Y < 0 || a.Y > 9 {
		return 0
	}

	if z < 0 {
		// 先处理掉异常情况
		return 0
	}

	if z == 0 {
		// 不能走动了
		if p == a {
			return 1
		}

		return 0
	}

	// 一般情况
	// 有 8 种可能
	next := z - 1
	p1 := chessProcess1(p, Pos{a.X + 2, a.Y + 1}, next)
	p2 := chessProcess1(p, Pos{a.X + 1, a.Y + 2}, next)
	p3 := chessProcess1(p, Pos{a.X - 1, a.Y + 2}, next)
	p4 := chessProcess1(p, Pos{a.X - 2, a.Y + 1}, next)
	p5 := chessProcess1(p, Pos{a.X - 2, a.Y - 1}, next)
	p6 := chessProcess1(p, Pos{a.X - 1, a.Y - 2}, next)
	p7 := chessProcess1(p, Pos{a.X + 1, a.Y - 2}, next)
	p8 := chessProcess1(p, Pos{a.X + 2, a.Y - 1}, next)

	return p1 + p2 + p3 + p4 + p5 + p6 + p7 + p8
}

// 将 chessProcess1 转换成动态规划
func chessDp(p Pos, step int) int {
	// 根据暴力尝试中三个变量（x,y,z）设计三维动态规划数组
	// 注意：step 维的数组长度是 step + 1，最后一步是用来做汇总判断的（参见暴力尝试的代码可知）
	Z, X, Y := step + 1, 9, 10
	arr := make([][][]int, Z)
	// 初始化
	for i := 0; i < Z; i++ {
		arr[i] = make([][]int, X)
		// 初始化第三维度
		for j := 0; j < X; j++ {
			arr[i][j] = make([]int, Y)
		}
	}

	// 初始值
	arr[0][p.X][p.Y] = 1

	// 注意：最外层要写 z（step），因为x、y方向上有 +、有 -，两个方向都存在依赖，只有 z 方向上是单向的
	// 所以，如果以 z 方向分层看待的话，就是本层仅依赖上层的，所以要先算完上层的数据再计算下一层
	for z := 1; z < Z; z++ {
		for x := X - 1; x >= 0; x-- {
			for y := Y - 1; y >= 0; y-- {
					next := z - 1
					p1 := pick(arr, x + 2, y + 1, next)//arr[x + 2][y + 1][next]//chessProcess1(p, Pos{a.X + 2, a.Y + 1}, next)
					p2 := pick(arr, x + 1, y + 2, next)//arr[x + 1][y + 2][next]//chessProcess1(p, Pos{a.X + 1, a.Y + 2}, next)
					p3 := pick(arr, x - 1, y + 2, next)//arr[x - 1][y + 2][next]//chessProcess1(p, Pos{a.X - 1, a.Y + 2}, next)
					p4 := pick(arr, x - 2, y + 1, next)//arr[x - 2][y + 1][next]//chessProcess1(p, Pos{a.X - 2, a.Y + 1}, next)
					p5 := pick(arr, x - 2, y - 1, next)//arr[x - 2][y - 1][next]//chessProcess1(p, Pos{a.X - 2, a.Y - 1}, next)
					p6 := pick(arr, x - 1, y - 2, next)//arr[x - 1][y - 2][next]//chessProcess1(p, Pos{a.X - 1, a.Y - 2}, next)
					p7 := pick(arr, x + 1, y - 2, next)//arr[x + 1][y - 2][next]//chessProcess1(p, Pos{a.X + 1, a.Y - 2}, next)
					p8 := pick(arr, x + 2, y - 1, next)//arr[x + 2][y - 1][next]//chessProcess1(p, Pos{a.X + 2, a.Y - 1}, next)

					arr[z][x][y] = p1 + p2 + p3 + p4 + p5 + p6 + p7 + p8
			}
		}
	}

	return arr[Z - 1][0][0]
}

func pick(arr [][][]int, x, y, z int) int {
	// 先判断位置 a 的合法性
	if x < 0 || x > 8 || y < 0 || y > 9 {
		return 0
	}

	return arr[z][x][y]
}