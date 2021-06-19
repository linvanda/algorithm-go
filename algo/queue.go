package algo

// 队列
// 可以使用数组或者链表实现队列
// 此处用数组实现环形队列，元素为 int
// 环形队列的实现：
//		enqueue 时将 right+1，dequeue 时将 left+1；
//		当 right 或 left 超过右边界时，则回到开头（即取模）
//		left 指向的位置：第一个元素位置；right 指向：最后一个元素的后面（空位置）
//		队列为空的状态：left == right；队列为满的状态：right + 1 == left
//		right 不会主动走到 left 位置，left 会走到 right 位置（队列空）

type Queue struct {
	c []int
	size int
	left int
	right int
	len int
}

func NewQueue(size int) *Queue {
	return &Queue{make([]int, size), size, 0, 0, 0}
}

func(q *Queue) Enqueue(item int) {
	// 先计算 right 的新位置
	r := q.right + 1
	if r == q.size {
		r = 0
	}

	// r 的位置是否被 left 占了
	if r == q.left {
		panic("queue is full")
	}

	// 注意这里：先将 item 放到 right 的位置，然后将 right + 1,即 right 指向的位置永远是空的
	q.c[q.right] = item
	q.right = r
	q.len++
}

func(q *Queue) Dequeue() int {
	// 当 left == right 的时候表示队列空了
	if q.left == q.right {
		panic("queue is empty")
	}

	// 先计算 left 的新位置
	l := q.left + 1
	if l == q.size {
		l = 0
	}

	// 弹出 left 位置的元素，并将 left 往后移动
	v := q.c[q.left]
	q.left = l
	q.len--

	return v
}

func(q *Queue) IsEmpty() bool {
	return q.len == 0
}

func(q *Queue) IsFull() bool {
	// 注意：循环队列中有一个空位不能放元素
	return q.len == q.size - 1
}

func(q *Queue) Length() int {
	return q.len;
}