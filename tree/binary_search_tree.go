package tree

import (
	"errors"
	"fmt"
	"strings"
)

// Comparable 可比较接口
type Comparable interface {
	CompareTo(other Comparable) int // 返回 -1, 0, 1 分别表示 小于, 等于, 大于
}

// IntValue 整数包装类型，实现Comparable接口
type IntValue int

func (i IntValue) CompareTo(other Comparable) int {
	if otherInt, ok := other.(IntValue); ok {
		if i < otherInt {
			return -1
		} else if i > otherInt {
			return 1
		}
		return 0
	}
	panic("cannot compare with non-IntValue")
}

// StringValue 字符串包装类型，实现Comparable接口
type StringValue string

func (s StringValue) CompareTo(other Comparable) int {
	if otherStr, ok := other.(StringValue); ok {
		if s < otherStr {
			return -1
		} else if s > otherStr {
			return 1
		}
		return 0
	}
	panic("cannot compare with non-StringValue")
}

// BSTNode 二叉搜索树节点
type BSTNode struct {
	Key   Comparable // 节点键值
	Value interface{} // 节点存储的数据
	Left  *BSTNode    // 左子节点
	Right *BSTNode    // 右子节点
}

// BinarySearchTree 二叉搜索树实现
// 思路：每个节点的左子树所有节点的键值都小于该节点，右子树所有节点的键值都大于该节点
// 作用：提供高效的搜索、插入、删除操作，支持有序遍历
// 业务场景：
// 1. 数据库索引：B+树的基础，提供快速查找
// 2. 文件系统：目录结构的组织
// 3. 表达式解析：语法树的构建
// 4. 游戏开发：场景管理，碰撞检测的空间划分
// 5. 编译器：符号表的实现
// 6. 缓存系统：有序数据的存储和查找
// 7. 范围查询：查找某个范围内的所有数据
type BinarySearchTree struct {
	root *BSTNode // 根节点
	size int      // 树的大小
}

// NewBinarySearchTree 创建新的二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
	}
}

// Insert 插入键值对
// 时间复杂度：平均O(log n)，最坏O(n)（退化为链表时）
func (bst *BinarySearchTree) Insert(key Comparable, value interface{}) {
	bst.root = bst.insertNode(bst.root, key, value)
}

// insertNode 递归插入节点
func (bst *BinarySearchTree) insertNode(node *BSTNode, key Comparable, value interface{}) *BSTNode {
	// 如果节点为空，创建新节点
	if node == nil {
		bst.size++
		return &BSTNode{
			Key:   key,
			Value: value,
			Left:  nil,
			Right: nil,
		}
	}
	
	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		// 插入到左子树
		node.Left = bst.insertNode(node.Left, key, value)
	} else if cmp > 0 {
		// 插入到右子树
		node.Right = bst.insertNode(node.Right, key, value)
	} else {
		// 键已存在，更新值
		node.Value = value
	}
	
	return node
}

// Search 搜索指定键的值
// 时间复杂度：平均O(log n)，最坏O(n)
func (bst *BinarySearchTree) Search(key Comparable) (interface{}, error) {
	node := bst.searchNode(bst.root, key)
	if node == nil {
		return nil, errors.New("key not found")
	}
	return node.Value, nil
}

// searchNode 递归搜索节点
func (bst *BinarySearchTree) searchNode(node *BSTNode, key Comparable) *BSTNode {
	if node == nil {
		return nil
	}
	
	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		return bst.searchNode(node.Left, key)
	} else if cmp > 0 {
		return bst.searchNode(node.Right, key)
	} else {
		return node
	}
}

// Contains 检查是否包含指定键
// 时间复杂度：平均O(log n)，最坏O(n)
func (bst *BinarySearchTree) Contains(key Comparable) bool {
	return bst.searchNode(bst.root, key) != nil
}

// Delete 删除指定键的节点
// 时间复杂度：平均O(log n)，最坏O(n)
func (bst *BinarySearchTree) Delete(key Comparable) error {
	if !bst.Contains(key) {
		return errors.New("key not found")
	}
	bst.root = bst.deleteNode(bst.root, key)
	bst.size--
	return nil
}

// deleteNode 递归删除节点
func (bst *BinarySearchTree) deleteNode(node *BSTNode, key Comparable) *BSTNode {
	if node == nil {
		return nil
	}
	
	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		node.Left = bst.deleteNode(node.Left, key)
	} else if cmp > 0 {
		node.Right = bst.deleteNode(node.Right, key)
	} else {
		// 找到要删除的节点
		if node.Left == nil {
			// 只有右子树或无子树
			return node.Right
		} else if node.Right == nil {
			// 只有左子树
			return node.Left
		} else {
			// 有两个子树，找到右子树的最小节点（中序后继）
			minNode := bst.findMin(node.Right)
			node.Key = minNode.Key
			node.Value = minNode.Value
			// 删除中序后继节点
			node.Right = bst.deleteNode(node.Right, minNode.Key)
		}
	}
	
	return node
}

// findMin 找到子树中的最小节点
func (bst *BinarySearchTree) findMin(node *BSTNode) *BSTNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// findMax 找到子树中的最大节点
func (bst *BinarySearchTree) findMax(node *BSTNode) *BSTNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// Min 获取树中的最小键
// 时间复杂度：O(log n)
func (bst *BinarySearchTree) Min() (Comparable, error) {
	if bst.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	minNode := bst.findMin(bst.root)
	return minNode.Key, nil
}

// Max 获取树中的最大键
// 时间复杂度：O(log n)
func (bst *BinarySearchTree) Max() (Comparable, error) {
	if bst.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	maxNode := bst.findMax(bst.root)
	return maxNode.Key, nil
}

// Size 返回树的大小
func (bst *BinarySearchTree) Size() int {
	return bst.size
}

// IsEmpty 判断树是否为空
func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}

// Height 计算树的高度
// 时间复杂度：O(n)
func (bst *BinarySearchTree) Height() int {
	return bst.calculateHeight(bst.root)
}

// calculateHeight 递归计算高度
func (bst *BinarySearchTree) calculateHeight(node *BSTNode) int {
	if node == nil {
		return -1 // 空树高度为-1，单节点树高度为0
	}
	
	leftHeight := bst.calculateHeight(node.Left)
	rightHeight := bst.calculateHeight(node.Right)
	
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// InorderTraversal 中序遍历（左-根-右）
// 时间复杂度：O(n)
// 结果是有序的
func (bst *BinarySearchTree) InorderTraversal() []Comparable {
	var result []Comparable
	bst.inorderHelper(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) inorderHelper(node *BSTNode, result *[]Comparable) {
	if node != nil {
		bst.inorderHelper(node.Left, result)
		*result = append(*result, node.Key)
		bst.inorderHelper(node.Right, result)
	}
}

// PreorderTraversal 前序遍历（根-左-右）
// 时间复杂度：O(n)
func (bst *BinarySearchTree) PreorderTraversal() []Comparable {
	var result []Comparable
	bst.preorderHelper(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) preorderHelper(node *BSTNode, result *[]Comparable) {
	if node != nil {
		*result = append(*result, node.Key)
		bst.preorderHelper(node.Left, result)
		bst.preorderHelper(node.Right, result)
	}
}

// PostorderTraversal 后序遍历（左-右-根）
// 时间复杂度：O(n)
func (bst *BinarySearchTree) PostorderTraversal() []Comparable {
	var result []Comparable
	bst.postorderHelper(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) postorderHelper(node *BSTNode, result *[]Comparable) {
	if node != nil {
		bst.postorderHelper(node.Left, result)
		bst.postorderHelper(node.Right, result)
		*result = append(*result, node.Key)
	}
}

// LevelOrderTraversal 层序遍历（广度优先）
// 时间复杂度：O(n)
func (bst *BinarySearchTree) LevelOrderTraversal() []Comparable {
	if bst.root == nil {
		return []Comparable{}
	}
	
	var result []Comparable
	queue := []*BSTNode{bst.root}
	
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		
		result = append(result, node.Key)
		
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	
	return result
}

// RangeQuery 范围查询，返回在[minKey, maxKey]范围内的所有键
// 时间复杂度：O(log n + k)，其中k是结果数量
func (bst *BinarySearchTree) RangeQuery(minKey, maxKey Comparable) []Comparable {
	var result []Comparable
	bst.rangeQueryHelper(bst.root, minKey, maxKey, &result)
	return result
}

func (bst *BinarySearchTree) rangeQueryHelper(node *BSTNode, minKey, maxKey Comparable, result *[]Comparable) {
	if node == nil {
		return
	}
	
	// 如果当前节点的键大于minKey，则搜索左子树
	if node.Key.CompareTo(minKey) > 0 {
		bst.rangeQueryHelper(node.Left, minKey, maxKey, result)
	}
	
	// 如果当前节点在范围内，添加到结果中
	if node.Key.CompareTo(minKey) >= 0 && node.Key.CompareTo(maxKey) <= 0 {
		*result = append(*result, node.Key)
	}
	
	// 如果当前节点的键小于maxKey，则搜索右子树
	if node.Key.CompareTo(maxKey) < 0 {
		bst.rangeQueryHelper(node.Right, minKey, maxKey, result)
	}
}

// Clear 清空树
func (bst *BinarySearchTree) Clear() {
	bst.root = nil
	bst.size = 0
}

// IsValidBST 验证是否为有效的二叉搜索树
func (bst *BinarySearchTree) IsValidBST() bool {
	return bst.isValidBSTHelper(bst.root, nil, nil)
}

func (bst *BinarySearchTree) isValidBSTHelper(node *BSTNode, min, max Comparable) bool {
	if node == nil {
		return true
	}
	
	// 检查当前节点是否违反BST性质
	if (min != nil && node.Key.CompareTo(min) <= 0) || (max != nil && node.Key.CompareTo(max) >= 0) {
		return false
	}
	
	// 递归检查左右子树
	return bst.isValidBSTHelper(node.Left, min, node.Key) && 
		   bst.isValidBSTHelper(node.Right, node.Key, max)
}

// String 字符串表示
func (bst *BinarySearchTree) String() string {
	if bst.IsEmpty() {
		return "BinarySearchTree{}"
	}
	
	inorder := bst.InorderTraversal()
	var elements []string
	for _, key := range inorder {
		elements = append(elements, fmt.Sprintf("%v", key))
	}
	
	return fmt.Sprintf("BinarySearchTree{size: %d, inorder: [%s]}", 
		bst.size, strings.Join(elements, ", "))
}

// PrintTree 打印树的结构（简单的可视化）
func (bst *BinarySearchTree) PrintTree() {
	if bst.root == nil {
		fmt.Println("Empty tree")
		return
	}
	bst.printTreeHelper(bst.root, "", true)
}

func (bst *BinarySearchTree) printTreeHelper(node *BSTNode, prefix string, isLast bool) {
	if node == nil {
		return
	}
	
	// 打印当前节点
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	fmt.Printf("%s%s%v\n", prefix, connector, node.Key)
	
	// 计算子节点的前缀
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}
	
	// 递归打印子节点
	if node.Left != nil || node.Right != nil {
		if node.Right != nil {
			bst.printTreeHelper(node.Right, childPrefix, node.Left == nil)
		}
		if node.Left != nil {
			bst.printTreeHelper(node.Left, childPrefix, true)
		}
	}
}

// 业务应用示例：
// 1. 数据库系统：索引结构，快速查找记录
// 2. 文件系统：目录树结构，文件路径查找
// 3. 编译器：符号表管理，变量和函数查找
// 4. 游戏引擎：空间划分，碰撞检测优化
// 5. 网络路由：路由表的实现
// 6. 缓存系统：有序数据的存储和范围查询
// 7. 搜索引擎：倒排索引的组织
// 8. 金融系统：价格区间查询，订单管理