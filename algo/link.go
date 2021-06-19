package algo

// 链表
// 此处实现包含哨兵的双向循环链表

type lNode struct {
	val interface{}
	prev *lNode
	next *lNode
}

type Link struct {
	// 哨兵
	nil *lNode
	len int
}

func NewLink() *Link {
	// 创建哨兵节点，初始化其 prev 和 next 都指向自身
	n := &lNode{val: nil}
	n.prev = n
	n.next = n

	return &Link{n, 0}
}

func(l *Link) Length() int {
	return l.len
}

// 将元素插入到表头
func(l *Link) Insert(item interface{}) {
	n := &lNode{val: item}
	n.prev = l.nil
	n.next = l.nil.next
	l.nil.next.prev = n
	l.nil.next = n
	l.len++
}

// 将元素追加到表尾
func(l *Link) Append(item interface{}) {
	n := &lNode{val:item}
	n.prev = l.nil.prev
	n.next = l.nil
	l.nil.prev.next = n
	l.nil.prev = n
	l.len++
}

// 删除元素 item（只删除第一个）
func(l *Link) Delete(item interface{})  {
	curr := l.nil.next
	for curr != l.nil {
		if curr.val == item {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			curr.prev = nil
			curr.next = nil
			l.len--
			break
		}

		curr = curr.next
	}
}

// 删除并返回表尾元素
func(l *Link) Pop() interface{} {
	if l.len == 0 {
		panic("link is empty")
	}

	c := l.nil.prev
	l.nil.prev = c.prev
	c.prev.next = l.nil
	c.prev = nil
	c.next = nil
	l.len--

	return c.val
}

// 删除并返回表头元素
func(l *Link) Shift() interface{} {
	if l.len == 0 {
		panic("link is empty")
	}

	c := l.nil.next
	l.nil.next = c.next
	c.next.prev = l.nil
	c.prev = nil
	c.next = nil
	l.len--

	return c.val
}
