package tree

import (
	"errors"
	"fmt"
	"strings"
)

// Color 红黑树节点颜色
type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

// RBNode 红黑树节点
type RBNode struct {
	Key    Comparable
	Value  interface{}
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
	color  Color
}

// RedBlackTree 红黑树实现
// 思路：通过节点着色和旋转操作维持以下性质：
//   1. 每个节点是红色或黑色
//   2. 根节点是黑色
//   3. 叶子节点（NIL）是黑色
//   4. 红色节点的两个子节点都是黑色（不能有两个连续的红色节点）
//   5. 从任意节点到其每个叶子节点的路径包含相同数量的黑色节点
// 作用：保证 O(log n) 的最坏情况时间复杂度，相比 AVL 树插入删除的旋转次数更少
// 业务场景：
// 1. Linux 内核：完全公平调度器（CFS）的进程调度
// 2. C++ STL：map 和 set 的底层实现
// 3. Java：TreeMap 和 TreeSet 的实现
// 4. 数据库：内存索引结构
// 5. 定时器管理：epoll 的事件管理
// 6. Nginx：连接和定时器管理
type RedBlackTree struct {
	root *RBNode
	size int
	leaf *RBNode // 哨兵叶子节点
}

// NewRedBlackTree 创建新的红黑树
func NewRedBlackTree() *RedBlackTree {
	leaf := &RBNode{color: BLACK}
	return &RedBlackTree{
		root: leaf,
		size: 0,
		leaf: leaf,
	}
}

func (rbt *RedBlackTree) isRed(node *RBNode) bool {
	return node != nil && node.color == RED
}

func (rbt *RedBlackTree) isBlack(node *RBNode) bool {
	return node == nil || node.color == BLACK
}

// rotateLeft 左旋
func (rbt *RedBlackTree) rotateLeft(x *RBNode) {
	y := x.Right
	x.Right = y.Left
	if y.Left != rbt.leaf {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == rbt.leaf {
		rbt.root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// rotateRight 右旋
func (rbt *RedBlackTree) rotateRight(y *RBNode) {
	x := y.Left
	y.Left = x.Right
	if x.Right != rbt.leaf {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == rbt.leaf {
		rbt.root = x
	} else if y == y.Parent.Right {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}
	x.Right = y
	y.Parent = x
}

// Insert 插入键值对
// 时间复杂度：O(log n)
func (rbt *RedBlackTree) Insert(key Comparable, value interface{}) {
	node := &RBNode{
		Key:    key,
		Value:  value,
		Left:   rbt.leaf,
		Right:  rbt.leaf,
		Parent: rbt.leaf,
		color:  RED,
	}

	// BST 插入
	parent := rbt.leaf
	current := rbt.root
	for current != rbt.leaf {
		parent = current
		cmp := key.CompareTo(current.Key)
		if cmp < 0 {
			current = current.Left
		} else if cmp > 0 {
			current = current.Right
		} else {
			current.Value = value
			return
		}
	}

	node.Parent = parent
	if parent == rbt.leaf {
		rbt.root = node
	} else if key.CompareTo(parent.Key) < 0 {
		parent.Left = node
	} else {
		parent.Right = node
	}

	rbt.size++
	rbt.insertFixup(node)
}

// insertFixup 插入后修复红黑树性质
func (rbt *RedBlackTree) insertFixup(z *RBNode) {
	for rbt.isRed(z.Parent) {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right // 叔节点
			if rbt.isRed(y) {
				// Case 1：叔节点为红色
				z.Parent.color = BLACK
				y.color = BLACK
				z.Parent.Parent.color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					// Case 2：叔节点黑色，z 是右孩子
					z = z.Parent
					rbt.rotateLeft(z)
				}
				// Case 3：叔节点黑色，z 是左孩子
				z.Parent.color = BLACK
				z.Parent.Parent.color = RED
				rbt.rotateRight(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if rbt.isRed(y) {
				z.Parent.color = BLACK
				y.color = BLACK
				z.Parent.Parent.color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rbt.rotateRight(z)
				}
				z.Parent.color = BLACK
				z.Parent.Parent.color = RED
				rbt.rotateLeft(z.Parent.Parent)
			}
		}
	}
	rbt.root.color = BLACK
}

// Search 搜索指定键的值
// 时间复杂度：O(log n)
func (rbt *RedBlackTree) Search(key Comparable) (interface{}, error) {
	node := rbt.searchNode(rbt.root, key)
	if node == rbt.leaf {
		return nil, errors.New("key not found")
	}
	return node.Value, nil
}

func (rbt *RedBlackTree) searchNode(node *RBNode, key Comparable) *RBNode {
	for node != rbt.leaf {
		cmp := key.CompareTo(node.Key)
		if cmp < 0 {
			node = node.Left
		} else if cmp > 0 {
			node = node.Right
		} else {
			return node
		}
	}
	return rbt.leaf
}

// Contains 检查键是否存在
func (rbt *RedBlackTree) Contains(key Comparable) bool {
	return rbt.searchNode(rbt.root, key) != rbt.leaf
}

// Delete 删除指定键
// 时间复杂度：O(log n)
func (rbt *RedBlackTree) Delete(key Comparable) error {
	z := rbt.searchNode(rbt.root, key)
	if z == rbt.leaf {
		return errors.New("key not found")
	}

	y := z
	yOriginalColor := y.color
	var x *RBNode

	if z.Left == rbt.leaf {
		x = z.Right
		rbt.transplant(z, z.Right)
	} else if z.Right == rbt.leaf {
		x = z.Left
		rbt.transplant(z, z.Left)
	} else {
		y = rbt.minimum(z.Right)
		yOriginalColor = y.color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			rbt.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		rbt.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		rbt.deleteFixup(x)
	}

	rbt.size--
	return nil
}

func (rbt *RedBlackTree) transplant(u, v *RBNode) {
	if u.Parent == rbt.leaf {
		rbt.root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

func (rbt *RedBlackTree) minimum(node *RBNode) *RBNode {
	for node.Left != rbt.leaf {
		node = node.Left
	}
	return node
}

// deleteFixup 删除后修复红黑树性质
func (rbt *RedBlackTree) deleteFixup(x *RBNode) {
	for x != rbt.root && x.color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right // 兄弟节点
			if rbt.isRed(w) {
				w.color = BLACK
				x.Parent.color = RED
				rbt.rotateLeft(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.color == BLACK && w.Right.color == BLACK {
				w.color = RED
				x = x.Parent
			} else {
				if w.Right.color == BLACK {
					w.Left.color = BLACK
					w.color = RED
					rbt.rotateRight(w)
					w = x.Parent.Right
				}
				w.color = x.Parent.color
				x.Parent.color = BLACK
				w.Right.color = BLACK
				rbt.rotateLeft(x.Parent)
				x = rbt.root
			}
		} else {
			w := x.Parent.Left
			if rbt.isRed(w) {
				w.color = BLACK
				x.Parent.color = RED
				rbt.rotateRight(x.Parent)
				w = x.Parent.Left
			}
			if w.Right.color == BLACK && w.Left.color == BLACK {
				w.color = RED
				x = x.Parent
			} else {
				if w.Left.color == BLACK {
					w.Right.color = BLACK
					w.color = RED
					rbt.rotateLeft(w)
					w = x.Parent.Left
				}
				w.color = x.Parent.color
				x.Parent.color = BLACK
				w.Left.color = BLACK
				rbt.rotateRight(x.Parent)
				x = rbt.root
			}
		}
	}
	x.color = BLACK
}

// Min 最小键
func (rbt *RedBlackTree) Min() (Comparable, error) {
	if rbt.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	return rbt.minimum(rbt.root).Key, nil
}

// Max 最大键
func (rbt *RedBlackTree) Max() (Comparable, error) {
	if rbt.IsEmpty() {
		return nil, errors.New("tree is empty")
	}
	node := rbt.root
	for node.Right != rbt.leaf {
		node = node.Right
	}
	return node.Key, nil
}

// Size 返回树大小
func (rbt *RedBlackTree) Size() int {
	return rbt.size
}

// IsEmpty 判断是否为空
func (rbt *RedBlackTree) IsEmpty() bool {
	return rbt.size == 0
}

// InorderTraversal 中序遍历
func (rbt *RedBlackTree) InorderTraversal() []Comparable {
	result := make([]Comparable, 0, rbt.size)
	rbt.inorderHelper(rbt.root, &result)
	return result
}

func (rbt *RedBlackTree) inorderHelper(node *RBNode, result *[]Comparable) {
	if node == rbt.leaf {
		return
	}
	rbt.inorderHelper(node.Left, result)
	*result = append(*result, node.Key)
	rbt.inorderHelper(node.Right, result)
}

// IsValid 验证红黑树性质
func (rbt *RedBlackTree) IsValid() bool {
	if rbt.root == rbt.leaf {
		return true
	}
	// 根节点必须是黑色
	if rbt.root.color != BLACK {
		return false
	}
	// 计算第一条路径的黑高
	blackHeight := -1
	return rbt.isValidHelper(rbt.root, 0, &blackHeight)
}

func (rbt *RedBlackTree) isValidHelper(node *RBNode, blackCount int, expectedBlackHeight *int) bool {
	if node == rbt.leaf {
		if *expectedBlackHeight == -1 {
			*expectedBlackHeight = blackCount
		}
		return blackCount == *expectedBlackHeight
	}
	// 不能有两个连续的红色节点
	if rbt.isRed(node) && (rbt.isRed(node.Left) || rbt.isRed(node.Right)) {
		return false
	}
	if node.color == BLACK {
		blackCount++
	}
	return rbt.isValidHelper(node.Left, blackCount, expectedBlackHeight) &&
		rbt.isValidHelper(node.Right, blackCount, expectedBlackHeight)
}

// Clear 清空树
func (rbt *RedBlackTree) Clear() {
	rbt.root = rbt.leaf
	rbt.size = 0
}

// String 字符串表示
func (rbt *RedBlackTree) String() string {
	if rbt.IsEmpty() {
		return "RedBlackTree{}"
	}
	inorder := rbt.InorderTraversal()
	elements := make([]string, len(inorder))
	for i, k := range inorder {
		elements[i] = fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("RedBlackTree{size: %d, inorder: [%s]}",
		rbt.size, strings.Join(elements, ", "))
}
