package algo

// 交错字符串
// https://leetcode.cn/problems/interleaving-string/
// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的

// 求 s 是否可以通过 s1、s2 交错组成
// 使用动态规划中的样本对应模型求解
func IsInterleavingString(s1 []byte, s2 []byte, s []byte) bool {
	if len(s) != len(s1) + len(s2) {
		return false
	}

	if len(s) == 0 && len(s1) == 0 && len(s2) == 0 {
		return true
	}

	L1, L2 := len(s1), len(s2)

	// 动态规划表
	// L1+1 * L2+1 的表，dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符能否交替组织出 s 的前 i+j 子串
	dp := make([][]bool, L1 + 1)
	for i := 0; i < L1 + 1; i++ {
		dp[i] = make([]bool, L2 + 1)
	}

	dp[0][0] = true

	// 初始值
	// 第 0 行：表示不使用 s1 中的字符，仅使用 s2 中前 j 个字符（子串）能否得出 s 中前 j 子串
	for j := 1; j <= L2; j++ {
		if s[j - 1] != s2[j - 1] {
			// 不相等，则从此处开始往后都是 false
			break
		}

		dp[0][j] = true
	}

	// 第 0 列：表示不使用 s2 中的字符，仅使用 s1 中前 i 个字符（子串）能否得出 s 中前 i 子串
	for i := 1; i <= L1; i++ {
		if s[i - 1] != s1[i - 1] {
			break
		}

		dp[i][0] = true
	}

	// 普通情况
	// s 中前 i + j 子串的组织情况
	// 从最后一个字符考虑：
	// 情况1：最后一个字符取 s1 的（s1的第 i 个字符，即 s1[i-1]），此时如果 s1 的前 i-1 子串和 s2 的前 j 子串能搞定 s 的前 j + i - 1 子串，则问题搞定
	// 情况2：最后一个字符取 s2 的，此时如果 s2 的前 j-1 子串和 s1 的前 i 子串能够搞定 s 的前 i + j - 1 子串，则问题搞定
	// 只要情况1 和情况2 有一个能成功，则成功
	for i := 1; i <= L1; i++ {
		for j := 1; j <= L2; j++ {
			ok := false
			// 注意：i、j 的含义是前多少个字符构成的子串，所以对应到 s1、s2 的下标要减 1
			if s[i+j-1] == s1[i-1] && dp[i-1][j] {
				ok = true
			} else if s[i+j-1] == s2[j-1] && dp[i][j-1] {
				ok = true
			}

			dp[i][j] = ok
		}
	}

	return dp[L1][L2]
}
