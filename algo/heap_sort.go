package algo

// 堆
// 堆是一颗完全二叉树
// 大顶堆：父节点值比子节点大
// 小顶堆：父节点值比子节点小
// 满二叉树：高度 h 的二叉树元素个数为 2^(h+1) - 1。如高度 2 的满二叉树元素个数为2^3 - 1 = 7。除了最后一层没有子节点，其它层都有两个节点（满子节点）
// 树的高度：从根节点到叶节点的最长简单路径的边数
// 完全二叉树：层数 h 的二叉树，其 h - 1 层为满二叉树，第 h 层元素从左到右排列
// 完全二叉树可以用数组表示
// 此处实现大顶堆

// 建堆
// 时间复杂度：O(n)
func MaxHeap(lst []int) {
	l := len(lst)
	// 从 i 开始，从下往上堆化，保证在处理第 i 个元素时，大于 i 的所有元素都已经完成堆化
	// 这里是个循环不变式：在下一次循环前， 所生成的是一个堆，下一次循环会在此基础上生成新的堆
	for i := l/2 - 1; i >= 0; i-- {
		maxHeapifiy(lst, i, l)
	}
}

// 堆排序
// 原址排序
// 时间复杂度：O(nlgn)
// 首先将 lst 堆化。堆化完毕后，堆顶一定是最大元素，然后将堆顶元素和最后元素交换位置，并重新堆化。如此反复，直到堆大小变成 2
func HeapSort(lst []int) {
	if len(lst) < 2 {
		return
	}

	MaxHeap(lst)

	for i := len(lst) - 1; i > 1; i-- {
		// 交换 lst[0] 和 lst[i]，使得 i 位置的元素成为当前堆中的最大元素
		lst[0], lst[i] = lst[i], lst[0]
		// 对新堆做堆化处理
		maxHeapifiy(lst, 0, i)
	}

	// 将最后 2 元素的堆元素互换即可
	lst[0], lst[1] = lst[1], lst[0]
}

// 对第 i 个元素堆化
// 前提：在对第 i 个元素堆化时，大于 i 的元素已经完成了堆化（其子元素子树）
// 注意：i 是数组下表，数组下标从 0 开始（有些地方为了方便计算是用 1 开始的）
func maxHeapifiy(lst []int, i int, hSize int) {
	l := i * 2 + 1
	r := l + 1
	max := hSize - 1
	largest := i

	if l <= max && lst[largest] < lst[l] {
		largest = l
	}

	if r <= max && lst[largest] < lst[r] {
		largest = r
	}

	// i 需要和 largest 做交换（父节点至少比其中一个子节点小）
	if largest != i {
		lst[largest], lst[i] = lst[i], lst[largest]

		// 交换后，需要对被交换的子节点继续做 heapify（因为该交换可能破坏了子节点的堆性质）
		maxHeapifiy(lst, largest, hSize)
	}
}