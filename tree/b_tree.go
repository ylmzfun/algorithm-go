package tree

import (
	"errors"
	"fmt"
	"strings"
)

// BTreeNode B树节点
type BTreeNode struct {
	keys     []Comparable  // 键数组
	values   []interface{} // 值数组
	children []*BTreeNode  // 子节点数组
	leaf     bool          // 是否为叶节点
	n        int           // 当前键的数量
}

// BTree B树实现（一种自平衡的多路搜索树）
// 思路：每个节点可包含多个键和子节点（阶数 t >= 2 时，节点包含 t-1 ~ 2t-1 个键），
//
//	通过分裂和合并操作保持平衡，所有叶节点在同一深度
//
// 作用：减少磁盘 I/O，优化大规模数据的存储和检索
// 业务场景：
// 1. 数据库索引：MySQL InnoDB 的 B+树索引基础
// 2. 文件系统：NTFS、HFS+、Ext4 等文件系统的目录索引
// 3. Key-Value 存储：LevelDB、RocksDB 的数据组织
// 4. 搜索引擎：倒排索引的存储
// 5. 缓存系统：大块数据的组织管理
type BTree struct {
	root *BTreeNode
	t    int // 最小度数
	size int
}

// NewBTree 创建新的 B树
// degree: 最小度数 t，每个节点至少有 t-1 个键，最多有 2t-1 个键
func NewBTree(degree int) *BTree {
	if degree < 2 {
		degree = 2
	}
	return &BTree{
		root: &BTreeNode{
			keys:     make([]Comparable, 2*degree-1),
			values:   make([]interface{}, 2*degree-1),
			children: make([]*BTreeNode, 2*degree),
			leaf:     true,
			n:        0,
		},
		t:    degree,
		size: 0,
	}
}

// Insert 插入键值对
// 时间复杂度：O(log n)
func (bt *BTree) Insert(key Comparable, value interface{}) {
	root := bt.root
	if root.n == 2*bt.t-1 {
		// 根节点已满，需要分裂
		newRoot := &BTreeNode{
			keys:     make([]Comparable, 2*bt.t-1),
			values:   make([]interface{}, 2*bt.t-1),
			children: make([]*BTreeNode, 2*bt.t),
			leaf:     false,
			n:        0,
		}
		bt.root = newRoot
		newRoot.children[0] = root
		bt.splitChild(newRoot, 0, root)
		bt.insertNonFull(newRoot, key, value)
	} else {
		bt.insertNonFull(root, key, value)
	}
}

// splitChild 分裂子节点
func (bt *BTree) splitChild(parent *BTreeNode, i int, child *BTreeNode) {
	t := bt.t
	newNode := &BTreeNode{
		keys:     make([]Comparable, 2*t-1),
		values:   make([]interface{}, 2*t-1),
		children: make([]*BTreeNode, 2*t),
		leaf:     child.leaf,
		n:        t - 1,
	}

	// 将 child 的后 t-1 个键复制到 newNode
	for j := 0; j < t-1; j++ {
		newNode.keys[j] = child.keys[j+t]
		newNode.values[j] = child.values[j+t]
	}

	// 如果不是叶节点，复制 t 个子节点
	if !child.leaf {
		for j := 0; j < t; j++ {
			newNode.children[j] = child.children[j+t]
		}
	}

	child.n = t - 1

	// 在父节点中为新节点腾出空间
	for j := parent.n; j > i; j-- {
		parent.children[j+1] = parent.children[j]
	}
	parent.children[i+1] = newNode

	// 将 child 的中间键提升到父节点
	for j := parent.n - 1; j >= i; j-- {
		parent.keys[j+1] = parent.keys[j]
		parent.values[j+1] = parent.values[j]
	}
	parent.keys[i] = child.keys[t-1]
	parent.values[i] = child.values[t-1]
	parent.n++
}

// insertNonFull 向非满节点插入
func (bt *BTree) insertNonFull(node *BTreeNode, key Comparable, value interface{}) {
	i := node.n - 1

	if node.leaf {
		// 找到插入位置
		for i >= 0 && key.CompareTo(node.keys[i]) < 0 {
			node.keys[i+1] = node.keys[i]
			node.values[i+1] = node.values[i]
			i--
		}
		if i >= 0 && key.CompareTo(node.keys[i]) == 0 {
			node.values[i] = value
			return
		}
		node.keys[i+1] = key
		node.values[i+1] = value
		node.n++
		bt.size++
	} else {
		// 找到合适的子节点
		for i >= 0 && key.CompareTo(node.keys[i]) < 0 {
			i--
		}
		i++
		if node.children[i].n == 2*bt.t-1 {
			bt.splitChild(node, i, node.children[i])
			if key.CompareTo(node.keys[i]) > 0 {
				i++
			}
		}
		bt.insertNonFull(node.children[i], key, value)
	}
}

// Search 搜索指定键的值
// 时间复杂度：O(log n)
func (bt *BTree) Search(key Comparable) (interface{}, error) {
	node := bt.root
	for {
		i := 0
		for i < node.n && key.CompareTo(node.keys[i]) > 0 {
			i++
		}
		if i < node.n && key.CompareTo(node.keys[i]) == 0 {
			return node.values[i], nil
		}
		if node.leaf {
			return nil, errors.New("key not found")
		}
		node = node.children[i]
	}
}

// Contains 检查键是否存在
func (bt *BTree) Contains(key Comparable) bool {
	_, err := bt.Search(key)
	return err == nil
}

// Delete 删除指定键
// 时间复杂度：O(log n)
func (bt *BTree) Delete(key Comparable) error {
	if !bt.Contains(key) {
		return errors.New("key not found")
	}
	bt.deleteKey(bt.root, key)
	bt.size--
	if bt.root.n == 0 && !bt.root.leaf {
		bt.root = bt.root.children[0]
	}
	return nil
}

func (bt *BTree) deleteKey(node *BTreeNode, key Comparable) {
	t := bt.t
	idx := 0
	for idx < node.n && key.CompareTo(node.keys[idx]) > 0 {
		idx++
	}

	if idx < node.n && key.CompareTo(node.keys[idx]) == 0 {
		// 键在当前节点
		if node.leaf {
			bt.removeFromLeaf(node, idx)
		} else {
			bt.removeFromNonLeaf(node, idx)
		}
	} else {
		// 键在子树中
		if node.leaf {
			return // 不应到达这里
		}
		flag := idx == node.n
		if node.children[idx].n < t {
			bt.fill(node, idx)
		}
		if flag && idx > node.n {
			bt.deleteKey(node.children[idx-1], key)
		} else {
			bt.deleteKey(node.children[idx], key)
		}
	}
}

func (bt *BTree) removeFromLeaf(node *BTreeNode, idx int) {
	for i := idx + 1; i < node.n; i++ {
		node.keys[i-1] = node.keys[i]
		node.values[i-1] = node.values[i]
	}
	node.n--
}

func (bt *BTree) removeFromNonLeaf(node *BTreeNode, idx int) {
	t := bt.t
	key := node.keys[idx]

	if node.children[idx].n >= t {
		// 用前驱替换
		pred := bt.getPred(node, idx)
		node.keys[idx] = pred.keys[pred.n-1]
		node.values[idx] = pred.values[pred.n-1]
		bt.deleteKey(node.children[idx], pred.keys[pred.n-1])
	} else if node.children[idx+1].n >= t {
		// 用后继替换
		succ := bt.getSucc(node, idx)
		node.keys[idx] = succ.keys[0]
		node.values[idx] = succ.values[0]
		bt.deleteKey(node.children[idx+1], succ.keys[0])
	} else {
		// 合并
		bt.merge(node, idx)
		bt.deleteKey(node.children[idx], key)
	}
}

func (bt *BTree) getPred(node *BTreeNode, idx int) *BTreeNode {
	cur := node.children[idx]
	for !cur.leaf {
		cur = cur.children[cur.n]
	}
	return cur
}

func (bt *BTree) getSucc(node *BTreeNode, idx int) *BTreeNode {
	cur := node.children[idx+1]
	for !cur.leaf {
		cur = cur.children[0]
	}
	return cur
}

func (bt *BTree) fill(node *BTreeNode, idx int) {
	t := bt.t
	if idx != 0 && node.children[idx-1].n >= t {
		bt.borrowFromPrev(node, idx)
	} else if idx != node.n && node.children[idx+1].n >= t {
		bt.borrowFromNext(node, idx)
	} else {
		if idx != node.n {
			bt.merge(node, idx)
		} else {
			bt.merge(node, idx-1)
		}
	}
}

func (bt *BTree) borrowFromPrev(node *BTreeNode, idx int) {
	child := node.children[idx]
	sibling := node.children[idx-1]

	for i := child.n - 1; i >= 0; i-- {
		child.keys[i+1] = child.keys[i]
		child.values[i+1] = child.values[i]
	}
	if !child.leaf {
		for i := child.n; i >= 0; i-- {
			child.children[i+1] = child.children[i]
		}
	}
	child.keys[0] = node.keys[idx-1]
	child.values[0] = node.values[idx-1]
	if !child.leaf {
		child.children[0] = sibling.children[sibling.n]
	}
	node.keys[idx-1] = sibling.keys[sibling.n-1]
	node.values[idx-1] = sibling.values[sibling.n-1]
	child.n++
	sibling.n--
}

func (bt *BTree) borrowFromNext(node *BTreeNode, idx int) {
	child := node.children[idx]
	sibling := node.children[idx+1]

	child.keys[child.n] = node.keys[idx]
	child.values[child.n] = node.values[idx]
	if !child.leaf {
		child.children[child.n+1] = sibling.children[0]
	}
	node.keys[idx] = sibling.keys[0]
	node.values[idx] = sibling.values[0]

	for i := 1; i < sibling.n; i++ {
		sibling.keys[i-1] = sibling.keys[i]
		sibling.values[i-1] = sibling.values[i]
	}
	if !sibling.leaf {
		for i := 1; i <= sibling.n; i++ {
			sibling.children[i-1] = sibling.children[i]
		}
	}
	child.n++
	sibling.n--
}

func (bt *BTree) merge(node *BTreeNode, idx int) {
	child := node.children[idx]
	sibling := node.children[idx+1]

	child.keys[child.n] = node.keys[idx]
	child.values[child.n] = node.values[idx]

	for i := 0; i < sibling.n; i++ {
		child.keys[i+child.n+1] = sibling.keys[i]
		child.values[i+child.n+1] = sibling.values[i]
	}
	if !child.leaf {
		for i := 0; i <= sibling.n; i++ {
			child.children[i+child.n+1] = sibling.children[i]
		}
	}

	for i := idx + 1; i < node.n; i++ {
		node.keys[i-1] = node.keys[i]
		node.values[i-1] = node.values[i]
	}
	for i := idx + 2; i <= node.n; i++ {
		node.children[i-1] = node.children[i]
	}
	child.n += sibling.n + 1
	node.n--
}

// Min 最小键
func (bt *BTree) Min() (Comparable, error) {
	if bt.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	node := bt.root
	for !node.leaf {
		node = node.children[0]
	}
	return node.keys[0], nil
}

// Max 最大键
func (bt *BTree) Max() (Comparable, error) {
	if bt.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	node := bt.root
	for !node.leaf {
		node = node.children[node.n]
	}
	return node.keys[node.n-1], nil
}

// Size 返回树的节点数
func (bt *BTree) Size() int {
	return bt.size
}

// IsEmpty 判断树是否为空
func (bt *BTree) IsEmpty() bool {
	return bt.size == 0
}

// InorderTraversal 中序遍历
func (bt *BTree) InorderTraversal() []Comparable {
	result := make([]Comparable, 0, bt.size)
	bt.inorderHelper(bt.root, &result)
	return result
}

func (bt *BTree) inorderHelper(node *BTreeNode, result *[]Comparable) {
	if node == nil {
		return
	}
	for i := 0; i < node.n; i++ {
		if !node.leaf {
			bt.inorderHelper(node.children[i], result)
		}
		*result = append(*result, node.keys[i])
	}
	if !node.leaf {
		bt.inorderHelper(node.children[node.n], result)
	}
}

// Clear 清空树
func (bt *BTree) Clear() {
	bt.root = &BTreeNode{
		keys:     make([]Comparable, 2*bt.t-1),
		values:   make([]interface{}, 2*bt.t-1),
		children: make([]*BTreeNode, 2*bt.t),
		leaf:     true,
		n:        0,
	}
	bt.size = 0
}

// String 字符串表示
func (bt *BTree) String() string {
	if bt.IsEmpty() {
		return "BTree{}"
	}
	inorder := bt.InorderTraversal()
	elements := make([]string, len(inorder))
	for i, k := range inorder {
		elements[i] = fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("BTree{degree: %d, size: %d, inorder: [%s]}",
		bt.t, bt.size, strings.Join(elements, ", "))
}
