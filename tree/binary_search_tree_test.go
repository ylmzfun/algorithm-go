package tree

import (
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	
	if bst.Size() != 0 {
		t.Errorf("Expected size 0, got %d", bst.Size())
	}
	
	if !bst.IsEmpty() {
		t.Error("New tree should be empty")
	}
	
	if bst.Height() != -1 {
		t.Errorf("Expected height -1 for empty tree, got %d", bst.Height())
	}
}

func TestInsertAndSearch(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 插入一些值
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(9), "nine")
	
	if bst.Size() != 5 {
		t.Errorf("Expected size 5, got %d", bst.Size())
	}
	
	if bst.IsEmpty() {
		t.Error("Tree should not be empty")
	}
	
	// 测试搜索
	val, err := bst.Search(IntValue(3))
	if err != nil {
		t.Errorf("Search failed: %v", err)
	}
	if val != "three" {
		t.Errorf("Expected 'three', got %v", val)
	}
	
	val, err = bst.Search(IntValue(7))
	if err != nil {
		t.Errorf("Search failed: %v", err)
	}
	if val != "seven" {
		t.Errorf("Expected 'seven', got %v", val)
	}
	
	// 搜索不存在的键
	_, err = bst.Search(IntValue(10))
	if err == nil {
		t.Error("Expected error for non-existent key")
	}
}

func TestContains(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	
	if !bst.Contains(IntValue(5)) {
		t.Error("Expected to contain 5")
	}
	
	if !bst.Contains(IntValue(3)) {
		t.Error("Expected to contain 3")
	}
	
	if bst.Contains(IntValue(10)) {
		t.Error("Should not contain 10")
	}
}

func TestUpdateValue(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(IntValue(5), "five")
	
	// 更新已存在键的值
	bst.Insert(IntValue(5), "FIVE")
	
	if bst.Size() != 1 {
		t.Errorf("Expected size 1 after update, got %d", bst.Size())
	}
	
	val, _ := bst.Search(IntValue(5))
	if val != "FIVE" {
		t.Errorf("Expected 'FIVE', got %v", val)
	}
}

func TestMinMax(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树测试
	_, err := bst.Min()
	if err == nil {
		t.Error("Expected error for empty tree")
	}
	
	_, err = bst.Max()
	if err == nil {
		t.Error("Expected error for empty tree")
	}
	
	// 插入值
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(9), "nine")
	
	min, err := bst.Min()
	if err != nil {
		t.Errorf("Min failed: %v", err)
	}
	if min != IntValue(1) {
		t.Errorf("Expected min 1, got %v", min)
	}
	
	max, err := bst.Max()
	if err != nil {
		t.Errorf("Max failed: %v", err)
	}
	if max != IntValue(9) {
		t.Errorf("Expected max 9, got %v", max)
	}
}

func TestDelete(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 删除不存在的键
	err := bst.Delete(IntValue(5))
	if err == nil {
		t.Error("Expected error when deleting non-existent key")
	}
	
	// 构建测试树
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(4), "four")
	bst.Insert(IntValue(6), "six")
	bst.Insert(IntValue(9), "nine")
	
	initialSize := bst.Size()
	
	// 删除叶子节点
	err = bst.Delete(IntValue(1))
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	
	if bst.Size() != initialSize-1 {
		t.Errorf("Expected size %d, got %d", initialSize-1, bst.Size())
	}
	
	if bst.Contains(IntValue(1)) {
		t.Error("Should not contain deleted key 1")
	}
	
	// 删除有一个子节点的节点
	err = bst.Delete(IntValue(6))
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	
	if bst.Contains(IntValue(6)) {
		t.Error("Should not contain deleted key 6")
	}
	
	// 删除有两个子节点的节点（根节点）
	err = bst.Delete(IntValue(5))
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	
	if bst.Contains(IntValue(5)) {
		t.Error("Should not contain deleted key 5")
	}
	
	// 验证树仍然是有效的BST
	if !bst.IsValidBST() {
		t.Error("Tree should still be a valid BST after deletions")
	}
}

func TestHeight(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树高度
	if bst.Height() != -1 {
		t.Errorf("Expected height -1 for empty tree, got %d", bst.Height())
	}
	
	// 单节点树
	bst.Insert(IntValue(5), "five")
	if bst.Height() != 0 {
		t.Errorf("Expected height 0 for single node, got %d", bst.Height())
	}
	
	// 添加更多节点
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	if bst.Height() != 1 {
		t.Errorf("Expected height 1, got %d", bst.Height())
	}
	
	bst.Insert(IntValue(1), "one")
	if bst.Height() != 2 {
		t.Errorf("Expected height 2, got %d", bst.Height())
	}
}

func TestInorderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树遍历
	result := bst.InorderTraversal()
	if len(result) != 0 {
		t.Errorf("Expected empty result for empty tree, got %v", result)
	}
	
	// 插入值
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(9), "nine")
	
	result = bst.InorderTraversal()
	expected := []Comparable{IntValue(1), IntValue(3), IntValue(5), IntValue(7), IntValue(9)}
	
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, result[i])
		}
	}
}

func TestPreorderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(9), "nine")
	
	result := bst.PreorderTraversal()
	expected := []Comparable{IntValue(5), IntValue(3), IntValue(1), IntValue(7), IntValue(9)}
	
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, result[i])
		}
	}
}

func TestPostorderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(9), "nine")
	
	result := bst.PostorderTraversal()
	expected := []Comparable{IntValue(1), IntValue(3), IntValue(9), IntValue(7), IntValue(5)}
	
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, result[i])
		}
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树测试
	result := bst.LevelOrderTraversal()
	if len(result) != 0 {
		t.Errorf("Expected empty result for empty tree, got %v", result)
	}
	
	// 构建树
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(4), "four")
	bst.Insert(IntValue(6), "six")
	bst.Insert(IntValue(9), "nine")
	
	result = bst.LevelOrderTraversal()
	expected := []Comparable{IntValue(5), IntValue(3), IntValue(7), IntValue(1), IntValue(4), IntValue(6), IntValue(9)}
	
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, result[i])
		}
	}
}

func TestRangeQuery(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树测试
	result := bst.RangeQuery(IntValue(1), IntValue(10))
	if len(result) != 0 {
		t.Errorf("Expected empty result for empty tree, got %v", result)
	}
	
	// 插入值
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(4), "four")
	bst.Insert(IntValue(6), "six")
	bst.Insert(IntValue(9), "nine")
	
	// 范围查询 [3, 7]
	result = bst.RangeQuery(IntValue(3), IntValue(7))
	expected := []Comparable{IntValue(3), IntValue(4), IntValue(5), IntValue(6), IntValue(7)}
	
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, result[i])
		}
	}
	
	// 范围查询 [10, 20] (无结果)
	result = bst.RangeQuery(IntValue(10), IntValue(20))
	if len(result) != 0 {
		t.Errorf("Expected empty result for range [10, 20], got %v", result)
	}
	
	// 范围查询 [1, 1] (单个值)
	result = bst.RangeQuery(IntValue(1), IntValue(1))
	if len(result) != 1 || result[0] != IntValue(1) {
		t.Errorf("Expected [1], got %v", result)
	}
}

func TestClear(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	
	bst.Clear()
	
	if bst.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", bst.Size())
	}
	
	if !bst.IsEmpty() {
		t.Error("Tree should be empty after clear")
	}
	
	if bst.Height() != -1 {
		t.Errorf("Expected height -1 after clear, got %d", bst.Height())
	}
}

func TestIsValidBST(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树是有效的BST
	if !bst.IsValidBST() {
		t.Error("Empty tree should be valid BST")
	}
	
	// 正常插入的树应该是有效的BST
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	bst.Insert(IntValue(1), "one")
	bst.Insert(IntValue(9), "nine")
	
	if !bst.IsValidBST() {
		t.Error("Properly constructed tree should be valid BST")
	}
}

func TestStringValue(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 测试字符串类型的键
	bst.Insert(StringValue("banana"), 2)
	bst.Insert(StringValue("apple"), 1)
	bst.Insert(StringValue("cherry"), 3)
	bst.Insert(StringValue("date"), 4)
	
	if bst.Size() != 4 {
		t.Errorf("Expected size 4, got %d", bst.Size())
	}
	
	// 中序遍历应该是按字典序排列的
	result := bst.InorderTraversal()
	expected := []Comparable{StringValue("apple"), StringValue("banana"), StringValue("cherry"), StringValue("date")}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, result[i])
		}
	}
	
	// 测试搜索
	val, err := bst.Search(StringValue("cherry"))
	if err != nil {
		t.Errorf("Search failed: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestString(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 空树
	str := bst.String()
	if str != "BinarySearchTree{}" {
		t.Errorf("Expected empty tree string, got %s", str)
	}
	
	// 非空树
	bst.Insert(IntValue(5), "five")
	bst.Insert(IntValue(3), "three")
	bst.Insert(IntValue(7), "seven")
	
	str = bst.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	t.Logf("String representation: %s", str)
}

// 测试边界情况
func TestEdgeCases(t *testing.T) {
	bst := NewBinarySearchTree()
	
	// 插入相同的键多次
	bst.Insert(IntValue(5), "first")
	bst.Insert(IntValue(5), "second")
	bst.Insert(IntValue(5), "third")
	
	if bst.Size() != 1 {
		t.Errorf("Expected size 1 for duplicate keys, got %d", bst.Size())
	}
	
	val, _ := bst.Search(IntValue(5))
	if val != "third" {
		t.Errorf("Expected 'third' (last inserted), got %v", val)
	}
	
	// 删除唯一节点
	bst.Delete(IntValue(5))
	if !bst.IsEmpty() {
		t.Error("Tree should be empty after deleting only node")
	}
}

// 性能测试
func BenchmarkInsert(b *testing.B) {
	bst := NewBinarySearchTree()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		bst.Insert(IntValue(i), i)
	}
}

func BenchmarkSearch(b *testing.B) {
	bst := NewBinarySearchTree()
	for i := 0; i < 1000; i++ {
		bst.Insert(IntValue(i), i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Search(IntValue(i % 1000))
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	bst := NewBinarySearchTree()
	for i := 0; i < 1000; i++ {
		bst.Insert(IntValue(i), i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.InorderTraversal()
	}
}