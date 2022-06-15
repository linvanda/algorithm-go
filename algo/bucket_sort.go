package algo

// 桶排序
// 非原址排序
// 期望时间复杂度：O(n)
// 稳定排序
// 思想：待排序元素均匀分布在[0, 1]上，或者是能够转化成 [0, 1]之间的值。我们将这些数均匀地放到 n 个桶中，然后对桶中的元素排序，最后按顺序
//		遍历桶中元素即可
func BucketSort(lst []float64) []float64 {
	// 这里假定 lst 中元素都在[0, 1]之间，这里不再做转化
	// 创建一个大小为 len + 1 的桶，桶中存放 link
	n := len(lst)
	bucket := newBucket(n + 1)

	// 元素择桶算法：bucket_index = floor(n*Val)
	for i := 0; i < n; i++ {
		idx :=int(float64(n) * lst[i])
		bucket.insert(idx, lst[i])
	}

	return bucket.traverse()
}

type bucket struct {
	buc []*node
	len int // bucket 大小
	count int // 元素数
}

func newBucket(l int) *bucket {
	b := make([]*node, l)
	// 初始化 -1（哨兵。因为真实数据都是正数）
	bSize := len(b)
	for i := 0; i < bSize; i++ {
		b[i] = &node{-1.0, nil}
	}
	return &bucket{b, l, 0}
}

// 将 Val 插入到 idx 位置
func(b *bucket) insert(idx int, val float64) {
	if idx >= b.len {
		panic("idx out of range")
	}

	cNode := b.buc[idx]
	// 插入到适当的位置（从低到高排序)。注意第一个元素是哨兵
	for {
		next := cNode.next
		if next == nil {
			// 最后了
			cNode.next = &node{val, nil}
			break
		}

		if val < next.val {
			// 比下一个小，则插入到当前位置
			n := node{val, nil}
			n.next = next
			cNode.next = &n
			break
		} else {
			// 继续
			cNode = next
		}
	}

	b.count++
}

// 遍历并返回排好序的 slice
func(b *bucket) traverse() []float64 {
	lst := make([]float64, b.count)
	p := 0
	for i := 0; i < b.len; i++ {
		// 遍历链表。注意链表第一个元素是哨兵要忽略掉
		c := b.buc[i]
		for {
			next := c.next
			if next == nil {
				break
			}

			lst[p] = next.val
			p++
			c = next
		}
	}

	return lst
}

type node struct {
	val float64
	next *node
}