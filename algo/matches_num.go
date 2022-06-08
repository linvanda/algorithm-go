package algo

import sort2 "sort"

// 给定一个数组 arr[2,3,4,6,9] 代表每个人的能力值
// 要求：能力值差值等于非负数 k 的才能在一起比赛
// 一局比赛有两人
// 求：最多可以组成多少场比赛
// 本题采用贪心算法
// 贪心算法：每步都选择最优解，根据局部最优解可以得到全局最优解
// 贪心算法一般都需要根据某策略排序，因为其每步都需要取最优解

// 采用滑动窗口+贪心算法实现
func MatchesNum(players []int, k int) int {
	if len(players) < 2 {
		return 0
	}

	// 先对 players 根据能力值升序排序
	sort2.Ints(players)

	N := len(players)
	// 窗口指针，初始化在 0 位置
	l, r := 0, 0
	// 记录每个运动员是否配对过了
	paired := make([]bool, N)
	// 能进行几场比赛
	cnt := 0

	for l < N && r < N {
		if l == r {
			// 自己和自己无法比赛
			r++
		} else if paired[l] {
			// l 指向的人已经配对过了，跳过
			l++
		} else {
			// 可以比较
			cmp := players[r] - players[l]
			if cmp == k {
				// 可以比赛，采用贪心策略，立即比赛
				cnt++
				// 这两人配对过了，同时向右滑动
				// 先要标记下 r 人（因为该人后面还会被 l 指到）
				paired[r] = true
				l++
				r++
			} else if cmp < k {
				// 差值不够，r 右移
				r++
			} else {
				// 差值超了，由于能力值是拍过序的，说明 l 指向的人没有合适的配对，跳过 l
				l++
			}
		}
	}

	return cnt
}

