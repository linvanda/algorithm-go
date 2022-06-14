package algo

// 求 a 的 n 次方
// n 是正数

// 方法：将指数拆解成二进制形式，如：7^75 = 7^(64+8+2+1) = 7^64 * 7^8 * 7^2 * 7^1
// 令 t = 7^1，然后不停地堆 t*t，然后看刚才的拆分中是否包含，包含则乘入
// 所谓包含，就是对应的二进制位是 1
func Power(a,n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return a
	}

	r := 1
	t := a
	for n != 0 {
		if n & 1 != 0 {
			r *= t
		}
		t *= t
		n = n >> 1
	}

	return r
}
