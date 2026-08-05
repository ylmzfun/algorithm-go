package main

import (
	"fmt"

	"algorithm-go/bloomfilter"
	"algorithm-go/graph"
	"algorithm-go/hash"
	"algorithm-go/heap"
	"algorithm-go/linkedlist"
	"algorithm-go/queue"
	"algorithm-go/search"
	"algorithm-go/sort"
	"algorithm-go/stack"
	"algorithm-go/tree"
	"algorithm-go/trie"
	"algorithm-go/unionfind"
)

// main 演示所有数据结构的基本用法
func main() {
	fmt.Println("=== Go 数据结构实现演示 ===")
	fmt.Println()

	// 1. 单向链表演示
	fmt.Println("1. 单向链表 (Singly Linked List)")
	demoSinglyLinkedList()
	fmt.Println()

	// 2. 双向链表演示
	fmt.Println("2. 双向链表 (Doubly Linked List)")
	demoDoublyLinkedList()
	fmt.Println()

	// 3. 栈演示
	fmt.Println("3. 栈 (Stack)")
	demoStack()
	fmt.Println()

	// 4. 队列演示
	fmt.Println("4. 队列 (Queue)")
	demoQueue()
	fmt.Println()

	// 5. 二叉搜索树演示
	fmt.Println("5. 二叉搜索树 (Binary Search Tree)")
	demoBST()
	fmt.Println()

	// 6. AVL树演示
	fmt.Println("6. AVL树 (AVL Tree)")
	demoAVLTree()
	fmt.Println()

	// 7. 红黑树演示
	fmt.Println("7. 红黑树 (Red-Black Tree)")
	demoRedBlackTree()
	fmt.Println()

	// 8. B树演示
	fmt.Println("8. B树 (B-Tree)")
	demoBTree()
	fmt.Println()

	// 9. 哈希表演示
	fmt.Println("9. 哈希表 (Hash Table)")
	demoHashTable()
	fmt.Println()

	// 10. 堆演示
	fmt.Println("10. 堆 (Heap)")
	demoHeap()
	fmt.Println()

	// 11. 图演示
	fmt.Println("11. 图 (Graph)")
	demoGraph()
	fmt.Println()

	// 12. 字典树演示
	fmt.Println("12. 字典树 (Trie)")
	demoTrie()
	fmt.Println()

	// 13. 并查集演示
	fmt.Println("13. 并查集 (Union-Find)")
	demoUnionFind()
	fmt.Println()

	// 14. 布隆过滤器演示
	fmt.Println("14. 布隆过滤器 (Bloom Filter)")
	demoBloomFilter()
	fmt.Println()

	// 15. 排序算法演示
	fmt.Println("15. 排序算法 (Sort)")
	demoSort()
	fmt.Println()

	// 16. 搜索算法演示
	fmt.Println("16. 搜索算法 (Search)")
	demoSearch()
	fmt.Println()

	fmt.Println("=== 演示完成 ===")
}

// demoSinglyLinkedList 演示单向链表
func demoSinglyLinkedList() {
	list := linkedlist.NewSinglyLinkedList()

	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddFirst(0)

	fmt.Printf("  链表内容: %s\n", list.String())
	fmt.Printf("  大小: %d\n", list.Size())

	list.RemoveFirst()
	fmt.Printf("  删除第一个后: %s\n", list.String())

	list.Reverse()
	fmt.Printf("  反转后: %s\n", list.String())
}

// demoDoublyLinkedList 演示双向链表
func demoDoublyLinkedList() {
	dll := linkedlist.NewDoublyLinkedList()

	dll.AddLast(1)
	dll.AddLast(2)
	dll.AddLast(3)
	dll.AddFirst(0)

	fmt.Printf("  链表内容: %s\n", dll.String())
	fmt.Printf("  大小: %d\n", dll.Size())

	first, _ := dll.GetFirst()
	last, _ := dll.GetLast()
	fmt.Printf("  第一个: %v, 最后一个: %v\n", first, last)

	dll.RemoveFirst()
	dll.RemoveLast()
	fmt.Printf("  删除首尾后: %s\n", dll.String())

	dll.Insert(1, 5)
	fmt.Printf("  在索引1插入5后: %s\n", dll.String())

	// 反向遍历
	fmt.Printf("  反向遍历: %v\n", dll.ToSliceReverse())
}

// demoStack 演示栈
func demoStack() {
	s := stack.NewStack(10)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Printf("  栈内容: %s\n", s.String())

	top, _ := s.Peek()
	fmt.Printf("  栈顶: %v\n", top)

	s.Pop()
	fmt.Printf("  出栈后: %s\n", s.String())

	expr := "((()))"
	fmt.Printf("  表达式 '%s' 括号匹配: %v\n", expr, stack.IsValidParentheses(expr))
}

// demoQueue 演示队列
func demoQueue() {
	q := queue.NewQueue(10)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Printf("  队列: %s\n", q.String())

	front, _ := q.Front()
	fmt.Printf("  队首: %v\n", front)

	q.Dequeue()
	fmt.Printf("  出队后: %s\n", q.String())

	// 双端队列演示
	deque := queue.NewDeque(10)
	deque.AddFront(1)
	deque.AddRear(2)
	deque.AddFront(0)
	fmt.Printf("  双端队列: %s\n", deque.String())
}

// demoBST 演示二叉搜索树
func demoBST() {
	bst := tree.NewBinarySearchTree()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		bst.Insert(tree.IntValue(val), val)
	}

	fmt.Printf("  BST大小: %d, 高度: %d\n", bst.Size(), bst.Height())

	min, _ := bst.Min()
	max, _ := bst.Max()
	fmt.Printf("  最小值: %v, 最大值: %v\n", min, max)
	fmt.Printf("  中序遍历: %v\n", bst.InorderTraversal())
	fmt.Printf("  范围查询[30,60]: %v\n", bst.RangeQuery(tree.IntValue(30), tree.IntValue(60)))
}

// demoAVLTree 演示AVL树
func demoAVLTree() {
	avl := tree.NewAVLTree()

	// 顺序插入会自动平衡
	for i := 1; i <= 7; i++ {
		avl.Insert(tree.IntValue(i*10), i*10)
	}

	fmt.Printf("  AVL大小: %d, 高度: %d\n", avl.Size(), avl.Height())
	fmt.Printf("  是否平衡: %v\n", avl.IsBalanced())
	fmt.Printf("  中序遍历: %v\n", avl.InorderTraversal())

	val, _ := avl.Search(tree.IntValue(30))
	fmt.Printf("  搜索30: %v\n", val)
}

// demoRedBlackTree 演示红黑树
func demoRedBlackTree() {
	rbt := tree.NewRedBlackTree()

	for i := 1; i <= 7; i++ {
		rbt.Insert(tree.IntValue(i*10), i*10)
	}

	fmt.Printf("  红黑树大小: %d\n", rbt.Size())
	fmt.Printf("  是否有效红黑树: %v\n", rbt.IsValid())

	min, _ := rbt.Min()
	max, _ := rbt.Max()
	fmt.Printf("  最小值: %v, 最大值: %v\n", min, max)
	fmt.Printf("  中序遍历: %v\n", rbt.InorderTraversal())
}

// demoBTree 演示B树
func demoBTree() {
	bt := tree.NewBTree(3)

	for i := 1; i <= 15; i++ {
		bt.Insert(tree.IntValue(i), i)
	}

	fmt.Printf("  B树大小: %d\n", bt.Size())

	min, _ := bt.Min()
	max, _ := bt.Max()
	fmt.Printf("  最小值: %v, 最大值: %v\n", min, max)

	val, _ := bt.Search(tree.IntValue(8))
	fmt.Printf("  搜索8: %v\n", val)
}

// demoHashTable 演示哈希表
func demoHashTable() {
	ht := hash.NewHashTable(16, 0.75)

	ht.Put("name", "Alice")
	ht.Put("age", 25)
	ht.Put("city", "Beijing")

	fmt.Printf("  哈希表大小: %d, 负载因子: %.2f\n", ht.Size(), ht.LoadFactor())

	if val, err := ht.Get("name"); err == nil {
		fmt.Printf("  name: %v\n", val)
	}

	fmt.Printf("  所有键: %v\n", ht.Keys())
}

// demoHeap 演示堆
func demoHeap() {
	// 最大堆
	maxHeap := heap.NewIntMaxHeap()
	values := []int{10, 20, 15, 30, 40}
	for _, val := range values {
		maxHeap.Insert(val)
	}

	top, _ := maxHeap.Peek()
	fmt.Printf("  最大堆顶: %v\n", top)

	max, _ := maxHeap.ExtractTop()
	fmt.Printf("  提取最大值: %v\n", max)

	// 堆排序
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	unsortedInterface := make([]interface{}, len(unsorted))
	for i, v := range unsorted {
		unsortedInterface[i] = v
	}
	compareFunc := func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	}
	heap.HeapSort(unsortedInterface, compareFunc)
	sorted := make([]int, len(unsortedInterface))
	for i, v := range unsortedInterface {
		sorted[i] = v.(int)
	}
	fmt.Printf("  堆排序: %v -> %v\n", unsorted, sorted)
}

// demoGraph 演示图
func demoGraph() {
	g := graph.NewAdjacencyListGraph(true)

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(0, 1, 1.0)
	g.AddEdge(0, 2, 1.0)
	g.AddEdge(1, 3, 1.0)
	g.AddEdge(2, 3, 1.0)
	g.AddEdge(3, 4, 1.0)

	fmt.Printf("  顶点数: %d, 边数: %d\n", g.GetVertexCount(), g.GetEdgeCount())
	fmt.Printf("  顶点0和4是否有路径: %v\n", graph.HasPath(g, 0, 4))

	// BFS遍历
	bfsResult := []interface{}{}
	graph.BFS(g, 0, func(vertex interface{}) {
		bfsResult = append(bfsResult, vertex)
	})
	fmt.Printf("  BFS: %v\n", bfsResult)

	// 拓扑排序
	topOrder, err := graph.TopologicalSort(g)
	if err == nil {
		fmt.Printf("  拓扑排序: %v\n", topOrder)
	}
}

// demoTrie 演示字典树
func demoTrie() {
	t := trie.NewTrie()

	words := []string{"apple", "app", "application", "apply", "banana"}
	for _, word := range words {
		t.Insert(word)
	}

	fmt.Printf("  字典树大小: %d\n", t.Size())
	fmt.Printf("  搜索'app': %v\n", t.Search("app"))
	fmt.Printf("  前缀'app'存在: %v\n", t.StartsWith("app"))

	// 自动补全
	suggestions := t.AutoComplete("app", 3)
	fmt.Printf("  'app'自动补全: %v\n", suggestions)
}

// demoUnionFind 演示并查集
func demoUnionFind() {
	uf := unionfind.NewUnionFind(6)

	fmt.Printf("  初始集合数: %d\n", uf.Count())

	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)

	fmt.Printf("  合并后集合数: %d\n", uf.Count())
	fmt.Printf("  0和1连通: %v\n", uf.Connected(0, 1))
	fmt.Printf("  0和2连通: %v\n", uf.Connected(0, 2))

	uf.Union(1, 2)
	fmt.Printf("  0和3连通(再合并后): %v\n", uf.Connected(0, 3))

	// 环检测
	edges := [][2]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Printf("  图环检测: %v\n", unionfind.DetectCycle(3, edges))
}

// demoBloomFilter 演示布隆过滤器
func demoBloomFilter() {
	bf := bloomfilter.NewBloomFilter(1000, 0.01)

	bf.AddString("apple")
	bf.AddString("banana")
	bf.AddString("orange")

	fmt.Printf("  布隆过滤器: %s\n", bf.String())
	fmt.Printf("  包含'apple': %v\n", bf.ContainsString("apple"))
	fmt.Printf("  包含'grape': %v\n", bf.ContainsString("grape"))
	fmt.Printf("  估算误判率: %.4f%%\n", bf.EstimatedFalsePositiveRate()*100)
}

// demoSort 演示排序算法
func demoSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	// 快速排序
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.QuickSort(arrCopy)
	fmt.Printf("  快速排序: %v -> %v\n", arr, arrCopy)

	// 归并排序
	copy(arrCopy, arr)
	sort.MergeSort(arrCopy)
	fmt.Printf("  归并排序: %v\n", sort.FormatArray(arrCopy))

	// 堆排序
	copy(arrCopy, arr)
	sort.HeapSort(arrCopy)
	fmt.Printf("  堆排序:   %v\n", sort.FormatArray(arrCopy))

	// 第K小元素
	kth := sort.FindKthSmallest(arr, 3)
	fmt.Printf("  第3小的元素: %d\n", kth)
}

// demoSearch 演示搜索算法
func demoSearch() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	fmt.Printf("  数组: %v\n", arr)

	idx := search.BinarySearch(arr, 7)
	fmt.Printf("  二分搜索 7: 索引=%d\n", idx)

	idx = search.JumpSearch(arr, 15)
	fmt.Printf("  跳跃搜索 15: 索引=%d\n", idx)

	idx = search.InterpolationSearch(arr, 3)
	fmt.Printf("  插值搜索 3: 索引=%d\n", idx)

	idx = search.ExponentialSearch(arr, 19)
	fmt.Printf("  指数搜索 19: 索引=%d\n", idx)

	idx = search.LinearSearch(arr, 5)
	fmt.Printf("  线性搜索 5: 索引=%d\n", idx)

	// LowerBound / UpperBound
	fmt.Printf("  LowerBound(8): %d\n", search.LowerBound(arr, 8))
	fmt.Printf("  UpperBound(8): %d\n", search.UpperBound(arr, 8))
}
