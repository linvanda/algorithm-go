package algo

import (
	"fmt"
	"strconv"
	"unicode"
)

// 字符串表达式如：3+4*(5+6*(4/2))+12*(-3)/(4-7)
// 实现一个计算器，计算其结果
// 实现思路：先思考最简单的形式：完全没有括号的情况，如：3+4*5+2-5/6，这种情况可以用栈实现计算
// 加上括号：括号其实就是分组（整个表达式最外层可以看做有一个隐式的括号），每组是个独立的表达式，最里层的需要最先计算，就是个递归

var opSymbols = map[byte]bool {
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

func Calculate(exp string) float64 {
	return innerCalc([]byte(exp))
}

func innerCalc(exp []byte) float64 {
	// stk：用来做当前平面的四则运算的栈
	// curr：当前正在读取的操作信息（操作数或者操作符）
	stk, curr := make([]string, 0, 20), make([]byte, 0, 5)
	// 指向栈 stk 中最早出现 *、/ 的位置，如果没有*和/则为-1
	pt := -1

	if exp[0] == '-' {
		// - 开头，说明是负数
		curr = append(curr, '-')
		exp = exp[1:]
	}

	for i := 0; i < len(exp); i++ {
		c := exp[i]
		if c == ' ' {
			// 过滤掉空格
			continue
		}

		if unicode.IsNumber(rune(c)) || c == '.' {
			// 普通的操作数：放入 curr 中
			curr = append(curr, c)
		} else if opSymbols[c] == true {
			// 操作符
			// 对于 +、-，需要先看栈顶是不是 *、/，如果是，则要先计算栈里面的属于 *、/ 的部分
			if c == '+' || c == '-' {
				stk, curr = tempCalc(stk, pt, curr)
			}

			// 当前操作数和操作符入栈
			stk = append(stk, string(curr), string(c))
			curr = curr[0:0]

			if c == '*' || c == '/' {
				if pt == -1 {
					// 指向操作符
					pt = len(stk) - 1
				}
			} else {
				pt = -1
			}
		} else if c == '(' {
			// 左括号，则要找到对应的有括号，将括号之间的表达式进一步求值
			k, j := 1, i+1
			for {
				// 每遇到左括号 k 加 1，遇到有括号减 1
				if exp[j] == '(' {
					k++
				} else if exp[j] == ')' {
					k--
				}

				if k == 0 {
					break
				}

				j++
			}

			// 此时 i 指向"("，j 指向对应的")"，取它们之间的表达式
			curr = []byte(strconv.FormatFloat(innerCalc(exp[i+1:j]), 'E', -1, 64))
			// 跳过这一坨表达式的区域
			i = j
		}
	}

	// 将 curr 入栈
	// 入栈前需要处理栈顶是 *、/ 的情况
	stk, curr = tempCalc(stk, pt, curr)
	stk = append(stk, string(curr))

	// 出栈计算
	// 此时栈里面只有 +、-
	// 注意要从左到右计算
	rst := baseCalc("0", stk[0], "+")
	for i := 1; i < len(stk); i += 2 {
		rst = baseCalc(float2str(rst), stk[i+1], stk[i])
	}

	fmt.Printf("====exp:%v, rst:%v\n", string(exp), rst)

	return rst
}

// 处理栈顶的 *、/
func tempCalc(stk []string, pt int, curr []byte) ([]string, []byte) {
	if pt == -1 {
		return stk, curr
	}

	t := stk[pt-1]
	// 注意该循环只到栈顶的下面一个元素
	// 3+4/5*6*
	tpt := pt
	for tpt < len(stk)-1 {
		t = float2str(baseCalc(t, stk[tpt+1], stk[tpt]))
		tpt += 2
	}
	curr = []byte(float2str(baseCalc(t, string(curr), stk[len(stk)-1])))
	// 出栈
	if pt > 0 {
		stk = stk[:pt-1]
	} else {
		stk = stk[0:0]
	}
	return stk, curr
}

func float2str(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

func baseCalc(s1, s2, op string) float64 {
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		panic("parse float fail")
	}
	f2, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		panic("parse float fail")
	}

	var r float64
	switch op {
	case "+":
		r = f1 + f2
	case "-":
		r = f1 - f2
	case "*":
		r = f1 * f2
	case "/":
		r = f1 / f2
	default:
		panic("invalid op:" + op)
	}

	return r
}


