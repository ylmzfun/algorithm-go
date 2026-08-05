package tree

import (
	"errors"
	"fmt"
	"strings"
)

// AVLNode AVL树节点
type AVLNode struct {
	Key    Comparable
	Value  interface{}
	Left   *AVLNode
	Right  *AVLNode
	height int // 节点高度
}

// AVLTree AVL树实现（自平衡二叉搜索树）
// 思路：通过维护每个节点的平衡因子（左子树高度-右子树高度），在插入和删除后通过旋转操作保持 |balanceFactor| <= 1
// 作用：保证 O(log n) 的查找、插入、删除时间，避免 BST 退化为链表
// 业务场景：
// 1. 数据库索引：需要频繁查找且数据有序的场景
// 2. 文件系统目录：需要快速定位的文件查找
// 3. 内存分配器：管理空闲内存块
// 4. 网络路由表：IP 地址前缀匹配
// 5. 游戏引擎：场景物体排序与查找
type AVLTree struct {
	root *AVLNode
	size int
}

// NewAVLTree 创建新的 AVL 树
func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
		size: 0,
	}
}

// height 获取节点高度（空节点高度为 -1）
func (avl *AVLTree) height(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

// updateHeight 更新节点高度
func (avl *AVLTree) updateHeight(node *AVLNode) {
	leftH := avl.height(node.Left)
	rightH := avl.height(node.Right)
	if leftH > rightH {
		node.height = leftH + 1
	} else {
		node.height = rightH + 1
	}
}

// balanceFactor 获取平衡因子：左子树高度 - 右子树高度
func (avl *AVLTree) balanceFactor(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return avl.height(node.Left) - avl.height(node.Right)
}

// rotateRight 右旋（左子树过高）
func (avl *AVLTree) rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	avl.updateHeight(y)
	avl.updateHeight(x)
	return x
}

// rotateLeft 左旋（右子树过高）
func (avl *AVLTree) rotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	avl.updateHeight(x)
	avl.updateHeight(y)
	return y
}

// balance 平衡节点
func (avl *AVLTree) balance(node *AVLNode) *AVLNode {
	if node == nil {
		return nil
	}

	avl.updateHeight(node)
	bf := avl.balanceFactor(node)

	// 左重 (LL 或 LR)
	if bf > 1 {
		if avl.balanceFactor(node.Left) < 0 {
			// LR：先左旋左子节点
			node.Left = avl.rotateLeft(node.Left)
		}
		// LL / LR：右旋当前节点
		return avl.rotateRight(node)
	}

	// 右重 (RR 或 RL)
	if bf < -1 {
		if avl.balanceFactor(node.Right) > 0 {
			// RL：先右旋右子节点
			node.Right = avl.rotateRight(node.Right)
		}
		// RR / RL：左旋当前节点
		return avl.rotateLeft(node)
	}

	return node
}

// Insert 插入键值对
// 时间复杂度：O(log n)
func (avl *AVLTree) Insert(key Comparable, value interface{}) {
	avl.root = avl.insertNode(avl.root, key, value)
}

func (avl *AVLTree) insertNode(node *AVLNode, key Comparable, value interface{}) *AVLNode {
	if node == nil {
		avl.size++
		return &AVLNode{Key: key, Value: value, height: 0}
	}

	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		node.Left = avl.insertNode(node.Left, key, value)
	} else if cmp > 0 {
		node.Right = avl.insertNode(node.Right, key, value)
	} else {
		node.Value = value
		return node
	}

	return avl.balance(node)
}

// Search 搜索指定键的值
// 时间复杂度：O(log n)
func (avl *AVLTree) Search(key Comparable) (interface{}, error) {
	node := avl.searchNode(avl.root, key)
	if node == nil {
		return nil, errors.New("key not found")
	}
	return node.Value, nil
}

func (avl *AVLTree) searchNode(node *AVLNode, key Comparable) *AVLNode {
	if node == nil {
		return nil
	}
	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		return avl.searchNode(node.Left, key)
	} else if cmp > 0 {
		return avl.searchNode(node.Right, key)
	}
	return node
}

// Contains 检查键是否存在
func (avl *AVLTree) Contains(key Comparable) bool {
	return avl.searchNode(avl.root, key) != nil
}

// Delete 删除指定键
// 时间复杂度：O(log n)
func (avl *AVLTree) Delete(key Comparable) error {
	if !avl.Contains(key) {
		return errors.New("key not found")
	}
	avl.root = avl.deleteNode(avl.root, key)
	avl.size--
	return nil
}

func (avl *AVLTree) deleteNode(node *AVLNode, key Comparable) *AVLNode {
	if node == nil {
		return nil
	}

	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		node.Left = avl.deleteNode(node.Left, key)
	} else if cmp > 0 {
		node.Right = avl.deleteNode(node.Right, key)
	} else {
		// 无子节点或只有一个子节点
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// 有两个子节点，找中序后继
		successor := avl.findMinNode(node.Right)
		node.Key = successor.Key
		node.Value = successor.Value
		node.Right = avl.deleteNode(node.Right, successor.Key)
	}

	return avl.balance(node)
}

func (avl *AVLTree) findMinNode(node *AVLNode) *AVLNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Min 获取最小键
func (avl *AVLTree) Min() (Comparable, error) {
	if avl.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	return avl.findMinNode(avl.root).Key, nil
}

// Max 获取最大键
func (avl *AVLTree) Max() (Comparable, error) {
	if avl.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	node := avl.root
	for node.Right != nil {
		node = node.Right
	}
	return node.Key, nil
}

// Size 返回树的节点数
func (avl *AVLTree) Size() int {
	return avl.size
}

// IsEmpty 判断树是否为空
func (avl *AVLTree) IsEmpty() bool {
	return avl.size == 0
}

// Height 返回树的高度
func (avl *AVLTree) Height() int {
	return avl.height(avl.root)
}

// InorderTraversal 中序遍历
func (avl *AVLTree) InorderTraversal() []Comparable {
	result := make([]Comparable, 0, avl.size)
	avl.inorderHelper(avl.root, &result)
	return result
}

func (avl *AVLTree) inorderHelper(node *AVLNode, result *[]Comparable) {
	if node == nil {
		return
	}
	avl.inorderHelper(node.Left, result)
	*result = append(*result, node.Key)
	avl.inorderHelper(node.Right, result)
}

// PreorderTraversal 前序遍历
func (avl *AVLTree) PreorderTraversal() []Comparable {
	result := make([]Comparable, 0, avl.size)
	avl.preorderHelper(avl.root, &result)
	return result
}

func (avl *AVLTree) preorderHelper(node *AVLNode, result *[]Comparable) {
	if node == nil {
		return
	}
	*result = append(*result, node.Key)
	avl.preorderHelper(node.Left, result)
	avl.preorderHelper(node.Right, result)
}

// IsBalanced 检查树是否平衡
func (avl *AVLTree) IsBalanced() bool {
	return avl.isBalancedHelper(avl.root)
}

func (avl *AVLTree) isBalancedHelper(node *AVLNode) bool {
	if node == nil {
		return true
	}
	bf := avl.balanceFactor(node)
	if bf < -1 || bf > 1 {
		return false
	}
	return avl.isBalancedHelper(node.Left) && avl.isBalancedHelper(node.Right)
}

// Clear 清空树
func (avl *AVLTree) Clear() {
	avl.root = nil
	avl.size = 0
}

// String 字符串表示
func (avl *AVLTree) String() string {
	if avl.IsEmpty() {
		return "AVLTree{}"
	}
	inorder := avl.InorderTraversal()
	elements := make([]string, len(inorder))
	for i, k := range inorder {
		elements[i] = fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("AVLTree{size: %d, height: %d, inorder: [%s]}",
		avl.size, avl.Height(), strings.Join(elements, ", "))
}
