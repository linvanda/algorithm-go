package algo

// 二叉树

// 节点定义
type Node struct {
	id int
	val int
	left *Node
	right *Node
}

func (n *Node) Val() int {
	return n.val
}

func (n *Node) Equal(node *Node) bool {
	if n == nil {
		if node == nil {
			return true
		}

		return false
	} else if node == nil {
		return false
	}

	return n.id == node.id
}

// 二叉搜索树
// 特点：节点 node 的左子树的值都比自己的小，右子树的值都>=自己的
type BSTree struct {
	idGen int
	node *Node
	cnt int
}

func NewBSTree() *BSTree {
	return &BSTree{idGen: 1}
}

// 节点数量
func (t *BSTree) Count() int {
	return t.cnt
}

// 添加节点
func (t *BSTree) Add(val int) {
	node := &Node{id: t.idGen, val: val}

	t.idGen++
	t.cnt++

	// 当前树是空的，直接加进去
	if t.node == nil {
		t.node = node
		return
	}

	// 插入到合适的叶子位置
	curr := t.node
	for {
		if node.val < curr.val {
			// 比 curr 的值小，放到它的左边
			if curr.left == nil {
				// 没有左节点，直接插入并结束
				curr.left = node
				return
			} else {
				// curr 有左节点，则继续往下
				curr = curr.left
			}
		} else {
			// 不小于 curr 的值，放到 curr 右边
			if curr.right == nil {
				curr.right = node
				return
			} else {
				curr = curr.right
			}
		}
	}
}

// 删除第一次找到的 val 相同的 node
// 删除成功则返回 true，没有找到对应的 node 返回 false
func (t *BSTree) Del(val int) bool {
	if t.node == nil {
		return false
	}

	curr, parent := t.findWithParent(val)
	if curr == nil {
		return false
	}

	// 找到了节点，删除该结点之前要找一个合适的新节点顶替该结点
	t.replaceAndDelNode(parent, curr, t.findReplaceNode(curr))

	t.cnt--

	return true
}

// 查找 value 等于 val 的第一个 Node 并返回
// 没有找到返回 nil
func (t *BSTree) Find(val int) *Node {
	curr, _ := t.findWithParent(val)

	return curr
}

// 前序遍历
func (t *BSTree) Preorder() []*Node {
	return preorder(t.node)
}

// 中序遍历
// 可以实现排序
func (t *BSTree) MiddleOrder() []*Node {
	return middleorder(t.node)
}

// 后序遍历
func (t *BSTree) Postorder() []*Node {
	return postorder(t.node)
}

// 递归实现前序遍历
// 特点：先处理父节点，在处理左节点，最后处理右节点
func preorder(node *Node) []*Node {
	if node == nil {
		return []*Node{}
	}

	arr := make([]*Node, 0, 1)

	// 先处理自己
	arr = append(arr, node)

	// 处理左节点
	if node.left != nil {
		arr = append(arr, preorder(node.left)...)
	}

	// 处理右节点
	if node.right != nil {
		arr = append(arr, preorder(node.right)...)
	}

	return arr
}

// 递归实现中序遍历
// 特点：先处理左节点，再处理父节点，最后处理右节点
func middleorder(node *Node) []*Node {
	if node == nil {
		return []*Node{}
	}

	arr := make([]*Node, 0, 1)

	// 处理左节点
	if node.left != nil {
		arr = append(arr, middleorder(node.left)...)
	}

	// 处理自己
	arr = append(arr, node)

	// 处理右节点
	if node.right != nil {
		arr = append(arr, middleorder(node.right)...)
	}

	return arr
}

// 递归实现后序遍历
// 特点：先处理左节点，再处理右节点，最后处理父节点
func postorder(node *Node) []*Node {
	if node == nil {
		return []*Node{}
	}

	arr := make([]*Node, 0, 1)

	// 处理左节点
	if node.left != nil {
		arr = append(arr, postorder(node.left)...)
	}

	// 处理右节点
	if node.right != nil {
		arr = append(arr, postorder(node.right)...)
	}

	// 处理自己
	arr = append(arr, node)

	return arr
}

// 查找用来替代被删除节点 curr 的新节点
func (t *BSTree) findReplaceNode(curr *Node) *Node {
	if curr.left == nil && curr.right == nil {
		// curr 本身是叶子节点
		return nil
	}

	// 如果 curr 只有一个子节点，则将这个子节点提上去即可
	if curr.left == nil {
		// 只有右子节点
		return curr.right
	}

	if curr.right == nil {
		// 只有左子节点
		return curr.left
	}

	// 左右子节点都存在:
	// 1. 如果其左节点没有右节点，则直接用其左节点
	// 2. 如果其右节点没有左节点，则直接用其右节点
	// 3. 从其左节点的右子树找，直到找到叶子节点
	if curr.left.right == nil {
		return curr.left
	}

	if curr.right.left == nil {
		return curr.right
	}

	node := curr.left
	p := curr
	for {
		// 如果 node 是叶子节点，则返回
		if node.left == nil && node.right == nil {
			// 先解除 parent 对它的引用
			t.replaceAndDelNode(p, node, nil)
			return node
		}

		// 不是叶子节点，如果有右节点，则继续看右节点
		p = node
		if node.right != nil {
			node = node.right
		} else {
			node = node.left
		}
	}
}

func (t *BSTree) replaceAndDelNode(parent, oldChild, newChild *Node) {
	if oldChild == nil {
		return
	}

	if newChild != nil {
		// 处理新旧节点的左孩子
		if newChild.left == nil && oldChild.left != nil && !oldChild.left.Equal(newChild) {
			newChild.left = oldChild.left
			oldChild.left = nil
		}

		// 处理新旧节点的右孩子
		if newChild.right == nil && oldChild.right != nil && !oldChild.right.Equal(newChild) {
			newChild.right = oldChild.right
			oldChild.right = nil
		}
	}

	if parent != nil {
		if oldChild.val >= parent.val {
			// oldChild 是 parent 的右节点
			parent.right = newChild
		} else {
			parent.left = newChild
		}
	} else {
		// 没有 parent，则调整 BSTree 指针的指向
		t.node = newChild
	}
}

func (t *BSTree) findWithParent(val int) (curr *Node, parent *Node) {
	curr = t.node
	parent = nil

	for {
		if curr == nil || curr.val == val {
			return
		}

		// 和当前节点不相等，比较大小以决定接下来和谁比较
		parent = curr
		if val < curr.val {
			// 小于当前节点的，跟其左节点比较
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
}
