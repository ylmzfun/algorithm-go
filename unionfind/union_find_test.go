package unionfind

import (
	"sort"
	"testing"
)

// TestNewUnionFind 测试创建并查集
func TestNewUnionFind(t *testing.T) {
	// 测试正常创建
	uf := NewUnionFind(5)
	if uf == nil {
		t.Error("NewUnionFind() returned nil")
	}
	
	if uf.Count() != 5 {
		t.Errorf("Expected count 5, got %d", uf.Count())
	}
	
	// 测试无效参数
	uf = NewUnionFind(0)
	if uf != nil {
		t.Error("NewUnionFind(0) should return nil")
	}
	
	uf = NewUnionFind(-1)
	if uf != nil {
		t.Error("NewUnionFind(-1) should return nil")
	}
}

// TestFind 测试查找操作
func TestFind(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 初始状态下，每个元素的根是自己
	for i := 0; i < 5; i++ {
		if uf.Find(i) != i {
			t.Errorf("Expected Find(%d) = %d, got %d", i, i, uf.Find(i))
		}
	}
	
	// 测试无效索引
	if uf.Find(-1) != -1 {
		t.Error("Find(-1) should return -1")
	}
	
	if uf.Find(5) != -1 {
		t.Error("Find(5) should return -1")
	}
}

// TestUnion 测试合并操作
func TestUnion(t *testing.T) {
	uf := NewUnionFind(5)
	
	initialCount := uf.Count()
	
	// 合并0和1
	if !uf.Union(0, 1) {
		t.Error("Union(0, 1) should return true")
	}
	
	if uf.Count() != initialCount-1 {
		t.Errorf("Expected count %d, got %d", initialCount-1, uf.Count())
	}
	
	// 再次合并0和1，应该返回false
	if uf.Union(0, 1) {
		t.Error("Union(0, 1) should return false when already connected")
	}
	
	// 合并2和3
	uf.Union(2, 3)
	
	// 合并(0,1)组和(2,3)组
	uf.Union(1, 2)
	
	if uf.Count() != 2 {
		t.Errorf("Expected count 2, got %d", uf.Count())
	}
	
	// 测试无效索引
	if uf.Union(-1, 0) {
		t.Error("Union(-1, 0) should return false")
	}
	
	if uf.Union(0, 5) {
		t.Error("Union(0, 5) should return false")
	}
}

// TestConnected 测试连通性检查
func TestConnected(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 初始状态下，元素不连通
	if uf.Connected(0, 1) {
		t.Error("0 and 1 should not be connected initially")
	}
	
	// 合并后应该连通
	uf.Union(0, 1)
	if !uf.Connected(0, 1) {
		t.Error("0 and 1 should be connected after union")
	}
	
	// 传递性测试
	uf.Union(1, 2)
	if !uf.Connected(0, 2) {
		t.Error("0 and 2 should be connected through 1")
	}
	
	// 测试无效索引
	if uf.Connected(-1, 0) {
		t.Error("Connected(-1, 0) should return false")
	}
	
	if uf.Connected(0, 5) {
		t.Error("Connected(0, 5) should return false")
	}
}

// TestSize 测试集合大小
func TestSize(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 初始大小都是1
	for i := 0; i < 5; i++ {
		if uf.Size(i) != 1 {
			t.Errorf("Expected size 1 for element %d, got %d", i, uf.Size(i))
		}
	}
	
	// 合并后大小应该增加
	uf.Union(0, 1)
	if uf.Size(0) != 2 {
		t.Errorf("Expected size 2 for element 0, got %d", uf.Size(0))
	}
	
	if uf.Size(1) != 2 {
		t.Errorf("Expected size 2 for element 1, got %d", uf.Size(1))
	}
	
	// 继续合并
	uf.Union(2, 3)
	uf.Union(0, 2)
	
	expectedSize := 4
	for i := 0; i < 4; i++ {
		if uf.Size(i) != expectedSize {
			t.Errorf("Expected size %d for element %d, got %d", expectedSize, i, uf.Size(i))
		}
	}
	
	// 测试无效索引
	if uf.Size(-1) != 0 {
		t.Error("Size(-1) should return 0")
	}
	
	if uf.Size(5) != 0 {
		t.Error("Size(5) should return 0")
	}
}

// TestGetSets 测试获取所有集合
func TestGetSets(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 初始状态：5个独立集合
	sets := uf.GetSets()
	if len(sets) != 5 {
		t.Errorf("Expected 5 sets, got %d", len(sets))
	}
	
	// 合并一些元素
	uf.Union(0, 1)
	uf.Union(2, 3)
	
	sets = uf.GetSets()
	if len(sets) != 3 {
		t.Errorf("Expected 3 sets, got %d", len(sets))
	}
	
	// 验证集合内容
	found01 := false
	found23 := false
	found4 := false
	
	for _, set := range sets {
		sort.Ints(set)
		if len(set) == 2 && set[0] == 0 && set[1] == 1 {
			found01 = true
		} else if len(set) == 2 && set[0] == 2 && set[1] == 3 {
			found23 = true
		} else if len(set) == 1 && set[0] == 4 {
			found4 = true
		}
	}
	
	if !found01 || !found23 || !found4 {
		t.Error("Sets content is incorrect")
	}
}

// TestGetRoots 测试获取根节点
func TestGetRoots(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 初始状态：5个根节点
	roots := uf.GetRoots()
	if len(roots) != 5 {
		t.Errorf("Expected 5 roots, got %d", len(roots))
	}
	
	// 合并后根节点减少
	uf.Union(0, 1)
	uf.Union(2, 3)
	
	roots = uf.GetRoots()
	if len(roots) != 3 {
		t.Errorf("Expected 3 roots, got %d", len(roots))
	}
}

// TestReset 测试重置操作
func TestReset(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 进行一些合并操作
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(1, 2)
	
	if uf.Count() == 5 {
		t.Error("Count should not be 5 after unions")
	}
	
	// 重置
	uf.Reset()
	
	if uf.Count() != 5 {
		t.Errorf("Expected count 5 after reset, got %d", uf.Count())
	}
	
	// 验证所有元素都是独立的
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if uf.Connected(i, j) {
				t.Errorf("Elements %d and %d should not be connected after reset", i, j)
			}
		}
	}
}

// TestString 测试字符串表示
func TestString(t *testing.T) {
	uf := NewUnionFind(3)
	
	// 测试初始状态
	str := uf.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	
	// 测试合并后
	uf.Union(0, 1)
	str = uf.String()
	if str == "" {
		t.Error("String representation should not be empty after union")
	}
	
	// 测试空并查集
	emptyUF := NewUnionFind(0)
	if emptyUF != nil {
		t.Error("Empty UnionFind should be nil")
	}
}

// TestWeightedUnionFind 测试加权并查集
func TestWeightedUnionFind(t *testing.T) {
	wuf := NewWeightedUnionFind(5)
	
	if wuf == nil {
		t.Error("NewWeightedUnionFind() returned nil")
	}
	
	if wuf.Count() != 5 {
		t.Errorf("Expected count 5, got %d", wuf.Count())
	}
	
	// 测试无效参数
	wuf = NewWeightedUnionFind(0)
	if wuf != nil {
		t.Error("NewWeightedUnionFind(0) should return nil")
	}
}

// TestWeightedUnionFindOperations 测试加权并查集操作
func TestWeightedUnionFindOperations(t *testing.T) {
	wuf := NewWeightedUnionFind(4)
	
	// 建立权重关系：0 --5--> 1 --3--> 2
	wuf.Union(0, 1, 5.0)
	wuf.Union(1, 2, 3.0)
	
	// 检查连通性
	if !wuf.Connected(0, 2) {
		t.Error("0 and 2 should be connected")
	}
	
	// 检查权重
	weight, ok := wuf.GetWeight(0, 2)
	if !ok {
		t.Error("Should be able to get weight between 0 and 2")
	}
	
	expected := 8.0 // 5 + 3
	if weight != expected {
		t.Errorf("Expected weight %f, got %f", expected, weight)
	}
	
	// 测试不连通的元素
	_, ok = wuf.GetWeight(0, 3)
	if ok {
		t.Error("Should not be able to get weight between disconnected elements")
	}
	
	// 测试无效索引
	if wuf.Find(-1) != -1 {
		t.Error("Find(-1) should return -1")
	}
	
	if wuf.Union(-1, 0, 1.0) {
		t.Error("Union with invalid index should return false")
	}
}

// TestKruskalMST 测试Kruskal最小生成树算法
func TestKruskalMST(t *testing.T) {
	// 测试简单图
	edges := []Edge{
		{0, 1, 4.0},
		{0, 2, 2.0},
		{1, 2, 1.0},
		{1, 3, 5.0},
		{2, 3, 3.0},
	}
	
	mst, totalWeight := KruskalMST(4, edges)
	
	if len(mst) != 3 {
		t.Errorf("Expected 3 edges in MST, got %d", len(mst))
	}
	
	expectedWeight := 6.0 // 1 + 2 + 3
	if totalWeight != expectedWeight {
		t.Errorf("Expected total weight %f, got %f", expectedWeight, totalWeight)
	}
	
	// 测试空输入
	mst, totalWeight = KruskalMST(0, []Edge{})
	if mst != nil || totalWeight != 0 {
		t.Error("Empty input should return nil MST and 0 weight")
	}
	
	mst, totalWeight = KruskalMST(3, []Edge{})
	if mst != nil || totalWeight != 0 {
		t.Error("No edges should return nil MST and 0 weight")
	}
}

// TestFindConnectedComponents 测试连通分量查找
func TestFindConnectedComponents(t *testing.T) {
	edges := [][2]int{
		{0, 1},
		{1, 2},
		{3, 4},
	}
	
	components := FindConnectedComponents(5, edges)
	
	if len(components) != 2 {
		t.Errorf("Expected 2 components, got %d", len(components))
	}
	
	// 验证组件内容
	foundComponent1 := false
	foundComponent2 := false
	
	for _, component := range components {
		sort.Ints(component)
		if len(component) == 3 && component[0] == 0 && component[1] == 1 && component[2] == 2 {
			foundComponent1 = true
		} else if len(component) == 2 && component[0] == 3 && component[1] == 4 {
			foundComponent2 = true
		}
	}
	
	if !foundComponent1 || !foundComponent2 {
		t.Error("Components content is incorrect")
	}
	
	// 测试空输入
	components = FindConnectedComponents(0, [][2]int{})
	if components != nil {
		t.Error("Empty input should return nil")
	}
}

// TestIsGraphConnected 测试图连通性判断
func TestIsGraphConnected(t *testing.T) {
	// 连通图
	edges := [][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	
	if !IsGraphConnected(4, edges) {
		t.Error("Graph should be connected")
	}
	
	// 非连通图
	edges = [][2]int{
		{0, 1},
		{2, 3},
	}
	
	if IsGraphConnected(4, edges) {
		t.Error("Graph should not be connected")
	}
	
	// 边界情况
	if !IsGraphConnected(1, [][2]int{}) {
		t.Error("Single node should be connected")
	}
	
	if !IsGraphConnected(0, [][2]int{}) {
		t.Error("Empty graph should be connected")
	}
}

// TestDetectCycle 测试环检测
func TestDetectCycle(t *testing.T) {
	// 无环图
	edges := [][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	
	if DetectCycle(4, edges) {
		t.Error("Should not detect cycle in tree")
	}
	
	// 有环图
	edges = [][2]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	
	if !DetectCycle(3, edges) {
		t.Error("Should detect cycle")
	}
	
	// 空图
	if DetectCycle(3, [][2]int{}) {
		t.Error("Empty graph should not have cycle")
	}
}

// TestSocialNetwork 测试社交网络应用
func TestSocialNetwork(t *testing.T) {
	sn := NewSocialNetwork(5)
	
	if sn == nil {
		t.Error("NewSocialNetwork() returned nil")
	}
	
	// 添加好友关系
	if !sn.AddFriendship(0, 1) {
		t.Error("AddFriendship should return true")
	}
	
	if !sn.AreFriends(0, 1) {
		t.Error("0 and 1 should be friends")
	}
	
	// 添加更多关系
	sn.AddFriendship(1, 2)
	sn.AddFriendship(3, 4)
	
	// 检查间接朋友关系
	if !sn.AreFriends(0, 2) {
		t.Error("0 and 2 should be friends through 1")
	}
	
	if sn.AreFriends(0, 3) {
		t.Error("0 and 3 should not be friends")
	}
	
	// 获取朋友圈
	groups := sn.GetFriendGroups()
	if len(groups) != 2 {
		t.Errorf("Expected 2 friend groups, got %d", len(groups))
	}
	
	// 获取最大朋友圈
	maxSize := sn.GetLargestFriendGroup()
	if maxSize != 3 {
		t.Errorf("Expected largest group size 3, got %d", maxSize)
	}
}

// TestNetworkConnectivity 测试网络连通性应用
func TestNetworkConnectivity(t *testing.T) {
	nc := NewNetworkConnectivity(4)
	
	if nc == nil {
		t.Error("NewNetworkConnectivity() returned nil")
	}
	
	// 初始状态：网络不完全连通
	if nc.IsNetworkFullyConnected() {
		t.Error("Network should not be fully connected initially")
	}
	
	if nc.GetNetworkComponents() != 4 {
		t.Errorf("Expected 4 components, got %d", nc.GetNetworkComponents())
	}
	
	// 连接节点
	nc.ConnectNodes(0, 1)
	nc.ConnectNodes(1, 2)
	nc.ConnectNodes(2, 3)
	
	// 检查连通性
	if !nc.AreConnected(0, 3) {
		t.Error("0 and 3 should be connected")
	}
	
	if !nc.IsNetworkFullyConnected() {
		t.Error("Network should be fully connected")
	}
	
	if nc.GetNetworkComponents() != 1 {
		t.Errorf("Expected 1 component, got %d", nc.GetNetworkComponents())
	}
}

// TestImageSegmentation 测试图像分割应用
func TestImageSegmentation(t *testing.T) {
	is := NewImageSegmentation(3, 3)
	
	if is == nil {
		t.Error("NewImageSegmentation() returned nil")
	}
	
	// 初始状态：9个独立区域
	if is.GetSegmentCount() != 9 {
		t.Errorf("Expected 9 segments, got %d", is.GetSegmentCount())
	}
	
	// 合并相邻像素
	is.MergePixels(0, 0, 0, 1) // 合并(0,0)和(0,1)
	is.MergePixels(0, 1, 1, 1) // 合并(0,1)和(1,1)
	
	if is.GetSegmentCount() != 7 {
		t.Errorf("Expected 7 segments after merging, got %d", is.GetSegmentCount())
	}
	
	// 检查区域大小
	size := is.GetSegmentSize(0, 0)
	if size != 3 {
		t.Errorf("Expected segment size 3, got %d", size)
	}
	
	// 测试无效坐标
	if is.MergePixels(-1, 0, 0, 0) {
		t.Error("MergePixels with invalid coordinates should return false")
	}
	
	if is.GetSegmentSize(-1, 0) != 0 {
		t.Error("GetSegmentSize with invalid coordinates should return 0")
	}
}

// TestPathCompression 测试路径压缩效果
func TestPathCompression(t *testing.T) {
	uf := NewUnionFind(10)
	
	// 创建一个长链：0-1-2-3-4-5-6-7-8-9
	for i := 0; i < 9; i++ {
		uf.Union(i, i+1)
	}
	
	// 多次查找应该触发路径压缩
	for i := 0; i < 5; i++ {
		uf.Find(9)
	}
	
	// 验证所有元素都连通
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if !uf.Connected(i, j) {
				t.Errorf("Elements %d and %d should be connected", i, j)
			}
		}
	}
}

// Benchmark tests

// BenchmarkUnionFindUnion 基准测试合并操作
func BenchmarkUnionFindUnion(b *testing.B) {
	uf := NewUnionFind(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Union(i%999, (i+1)%1000)
	}
}

// BenchmarkUnionFindFind 基准测试查找操作
func BenchmarkUnionFindFind(b *testing.B) {
	uf := NewUnionFind(1000)
	
	// 预先进行一些合并操作
	for i := 0; i < 500; i++ {
		uf.Union(i, i+500)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Find(i % 1000)
	}
}

// BenchmarkUnionFindConnected 基准测试连通性检查
func BenchmarkUnionFindConnected(b *testing.B) {
	uf := NewUnionFind(1000)
	
	// 预先进行一些合并操作
	for i := 0; i < 500; i++ {
		uf.Union(i, i+500)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Connected(i%1000, (i+1)%1000)
	}
}

// BenchmarkWeightedUnionFind 基准测试加权并查集
func BenchmarkWeightedUnionFind(b *testing.B) {
	wuf := NewWeightedUnionFind(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wuf.Union(i%999, (i+1)%1000, float64(i))
	}
}

// BenchmarkKruskalMST 基准测试Kruskal算法
func BenchmarkKruskalMST(b *testing.B) {
	// 生成测试边
	edges := make([]Edge, 1000)
	for i := 0; i < 1000; i++ {
		edges[i] = Edge{
			From:   i % 100,
			To:     (i + 1) % 100,
			Weight: float64(i),
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KruskalMST(100, edges)
	}
}

// BenchmarkSocialNetwork 基准测试社交网络
func BenchmarkSocialNetwork(b *testing.B) {
	sn := NewSocialNetwork(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn.AddFriendship(i%1000, (i+1)%1000)
	}
}