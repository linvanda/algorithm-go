package algo

//给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//数组中的每个元素 n 代表你在该位置可以跳跃长度：你要么跳 1，要么跳 n。
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//假设你总是可以到达数组的最后一个位置。
// 注意：这里和 LeeCode 的有点区别：LeeCode 的是说 n 是你能跳的最大长度，比如 n 表示你能跳 1 - n

func JumpGame(arr []int) int {
	l := len(arr)
	if l <= 1 {
		return 0
	}

	if l == 2 {
		return 1
	}

	// 该数组记录从 i 位置跳到 arr 的末尾最少需要几步
	dp := make([]int, l)
	dp[l - 1] = 0
	dp[l - 2] = 1

	for i := l - 3; i >= 0; i-- {
		st := arr[i]
		if i + st == l - 1 {
			// 一步就能到
			dp[i] = 1
		} else if i + st > l - 1 {
			// 越界了，只能跳一步，则 i 位置跳到末尾的最少步数是 i+1 位置的步数加 1
			dp[i] = dp[i+1] + 1
		} else {
			// 有两种选择：只跳一步，或者跳 st 步，去总步数最小的
			dp[i] = minInt(dp[i+1], dp[i+st]) + 1
		}
	}

	return dp[0]
}
