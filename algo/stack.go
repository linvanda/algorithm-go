package algo

// 栈
// 可以用数组和链表实现栈
// 此处我们用数组实现大小固定的栈，栈元素为 int

type Stack struct {
	stk []int
	p int
	size int
}

func NewStack(size int) *Stack {
	return &Stack{make([]int, size), -1, size}
}

func(s *Stack) Push(i int)  {
	if s.p == s.size - 1 {
		panic("overflow")
	}

	s.p++
	s.stk[s.p] = i
}

func(s *Stack) Pop() int  {
	if s.p == -1 {
		panic("stack is empty")
	}

	v := s.stk[s.p]
	s.p--
	return v
}

func(s *Stack) isEmpty() bool {
	return s.p < 0
}

func(s *Stack) Length() int {
	return s.p + 1
}