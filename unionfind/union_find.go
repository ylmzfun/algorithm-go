package unionfind

import (
	"fmt"
	"strings"
)

// UnionFind 并查集数据结构
// 思路：使用树形结构表示集合，每个节点指向其父节点，根节点指向自己
// 作用：高效地进行集合的合并和查找操作，判断两个元素是否属于同一集合
// 业务场景：
// 1. 社交网络：朋友圈分析，社区发现
// 2. 网络连通性：判断网络节点是否连通
// 3. 图像处理：连通区域标记，图像分割
// 4. 编译器：变量等价类分析
// 5. 数据库：关系型数据的聚类分析
// 6. 游戏开发：地图连通性检测
// 7. 生物信息学：基因家族分类
// 8. 机器学习：聚类算法，特征选择
type UnionFind struct {
	parent []int // 父节点数组
	rank   []int // 秩（树的高度）数组，用于优化
	size   []int // 每个集合的大小
	count  int   // 集合的数量
	n      int   // 元素总数
}

// NewUnionFind 创建新的并查集
// n: 元素的数量（0到n-1）
func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
		count:  n,
		n:      n,
	}
	
	// 初始化：每个元素都是独立的集合
	for i := 0; i < n; i++ {
		uf.parent[i] = i // 每个元素的父节点是自己
		uf.rank[i] = 0   // 初始秩为0
		uf.size[i] = 1   // 初始大小为1
	}
	
	return uf
}

// Find 查找元素x所属集合的根节点（代表元素）
// 使用路径压缩优化：将路径上的所有节点直接指向根节点
// 时间复杂度：O(α(n))，其中α是阿克曼函数的反函数，实际上近似常数
func (uf *UnionFind) Find(x int) int {
	if !uf.isValid(x) {
		return -1
	}
	
	// 路径压缩：递归查找根节点，并将路径上所有节点直接指向根节点
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	
	return uf.parent[x]
}

// Union 合并两个元素所属的集合
// 使用按秩合并优化：将秩小的树合并到秩大的树下
// 时间复杂度：O(α(n))
func (uf *UnionFind) Union(x, y int) bool {
	if !uf.isValid(x) || !uf.isValid(y) {
		return false
	}
	
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	// 如果已经在同一集合中，无需合并
	if rootX == rootY {
		return false
	}
	
	// 按秩合并：将秩小的树合并到秩大的树下
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	} else {
		// 秩相等时，任选一个作为根，并增加其秩
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
		uf.rank[rootX]++
	}
	
	uf.count-- // 集合数量减1
	return true
}

// Connected 判断两个元素是否属于同一集合
// 时间复杂度：O(α(n))
func (uf *UnionFind) Connected(x, y int) bool {
	if !uf.isValid(x) || !uf.isValid(y) {
		return false
	}
	
	return uf.Find(x) == uf.Find(y)
}

// Count 返回集合的数量
func (uf *UnionFind) Count() int {
	return uf.count
}

// Size 返回元素x所属集合的大小
func (uf *UnionFind) Size(x int) int {
	if !uf.isValid(x) {
		return 0
	}
	
	root := uf.Find(x)
	return uf.size[root]
}

// GetSets 获取所有集合
func (uf *UnionFind) GetSets() map[int][]int {
	sets := make(map[int][]int)
	
	for i := 0; i < uf.n; i++ {
		root := uf.Find(i)
		sets[root] = append(sets[root], i)
	}
	
	return sets
}

// GetRoots 获取所有根节点
func (uf *UnionFind) GetRoots() []int {
	roots := make([]int, 0, uf.count)
	rootSet := make(map[int]bool)
	
	for i := 0; i < uf.n; i++ {
		root := uf.Find(i)
		if !rootSet[root] {
			roots = append(roots, root)
			rootSet[root] = true
		}
	}
	
	return roots
}

// Reset 重置并查集
func (uf *UnionFind) Reset() {
	for i := 0; i < uf.n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
		uf.size[i] = 1
	}
	uf.count = uf.n
}

// isValid 检查索引是否有效
func (uf *UnionFind) isValid(x int) bool {
	return x >= 0 && x < uf.n
}

// String 字符串表示
func (uf *UnionFind) String() string {
	if uf.n == 0 {
		return "UnionFind{}"
	}
	
	sets := uf.GetSets()
	var parts []string
	
	for root, elements := range sets {
		var elemStrs []string
		for _, elem := range elements {
			elemStrs = append(elemStrs, fmt.Sprintf("%d", elem))
		}
		parts = append(parts, fmt.Sprintf("{%s}[root:%d]", strings.Join(elemStrs, ","), root))
	}
	
	return fmt.Sprintf("UnionFind{sets: [%s], count: %d}", strings.Join(parts, ", "), uf.count)
}

// WeightedUnionFind 加权并查集
// 支持带权重的合并操作，可以维护元素间的相对关系
type WeightedUnionFind struct {
	parent []int     // 父节点数组
	weight []float64 // 权重数组，表示到父节点的权重
	count  int       // 集合数量
	n      int       // 元素总数
}

// NewWeightedUnionFind 创建加权并查集
func NewWeightedUnionFind(n int) *WeightedUnionFind {
	if n <= 0 {
		return nil
	}
	
	wuf := &WeightedUnionFind{
		parent: make([]int, n),
		weight: make([]float64, n),
		count:  n,
		n:      n,
	}
	
	for i := 0; i < n; i++ {
		wuf.parent[i] = i
		wuf.weight[i] = 0.0
	}
	
	return wuf
}

// Find 查找根节点并更新权重
func (wuf *WeightedUnionFind) Find(x int) int {
	if !wuf.isValid(x) {
		return -1
	}
	
	if wuf.parent[x] != x {
		originalParent := wuf.parent[x]
		wuf.parent[x] = wuf.Find(wuf.parent[x])
		wuf.weight[x] += wuf.weight[originalParent]
	}
	
	return wuf.parent[x]
}

// Union 合并两个集合，指定权重关系
// weight: x到y的权重
func (wuf *WeightedUnionFind) Union(x, y int, weight float64) bool {
	if !wuf.isValid(x) || !wuf.isValid(y) {
		return false
	}
	
	rootX := wuf.Find(x)
	rootY := wuf.Find(y)
	
	if rootX == rootY {
		return false
	}
	
	// 将rootX合并到rootY下
	wuf.parent[rootX] = rootY
	wuf.weight[rootX] = wuf.weight[y] - wuf.weight[x] + weight
	wuf.count--
	
	return true
}

// GetWeight 获取两个元素间的权重差
func (wuf *WeightedUnionFind) GetWeight(x, y int) (float64, bool) {
	if !wuf.isValid(x) || !wuf.isValid(y) {
		return 0, false
	}
	
	if wuf.Find(x) != wuf.Find(y) {
		return 0, false // 不在同一集合中
	}
	
	return wuf.weight[x] - wuf.weight[y], true
}

// Connected 判断是否连通
func (wuf *WeightedUnionFind) Connected(x, y int) bool {
	if !wuf.isValid(x) || !wuf.isValid(y) {
		return false
	}
	
	return wuf.Find(x) == wuf.Find(y)
}

// Count 返回集合数量
func (wuf *WeightedUnionFind) Count() int {
	return wuf.count
}

// isValid 检查索引有效性
func (wuf *WeightedUnionFind) isValid(x int) bool {
	return x >= 0 && x < wuf.n
}

// 经典应用算法

// Kruskal 使用并查集实现Kruskal最小生成树算法
type Edge struct {
	From   int
	To     int
	Weight float64
}

// KruskalMST 计算最小生成树
func KruskalMST(n int, edges []Edge) ([]Edge, float64) {
	if n <= 0 || len(edges) == 0 {
		return nil, 0
	}
	
	// 按权重排序边
	sortedEdges := make([]Edge, len(edges))
	copy(sortedEdges, edges)
	
	// 简单的冒泡排序（实际应用中应使用更高效的排序算法）
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-1-i; j++ {
			if sortedEdges[j].Weight > sortedEdges[j+1].Weight {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	uf := NewUnionFind(n)
	mst := make([]Edge, 0, n-1)
	totalWeight := 0.0
	
	for _, edge := range sortedEdges {
		if !uf.Connected(edge.From, edge.To) {
			uf.Union(edge.From, edge.To)
			mst = append(mst, edge)
			totalWeight += edge.Weight
			
			// 如果已经有n-1条边，MST完成
			if len(mst) == n-1 {
				break
			}
		}
	}
	
	return mst, totalWeight
}

// FindConnectedComponents 查找图的连通分量
func FindConnectedComponents(n int, edges [][2]int) [][]int {
	if n <= 0 {
		return nil
	}
	
	uf := NewUnionFind(n)
	
	// 合并所有边连接的节点
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	
	// 获取所有连通分量
	sets := uf.GetSets()
	components := make([][]int, 0, len(sets))
	
	for _, component := range sets {
		components = append(components, component)
	}
	
	return components
}

// IsGraphConnected 判断图是否连通
func IsGraphConnected(n int, edges [][2]int) bool {
	if n <= 1 {
		return true
	}
	
	uf := NewUnionFind(n)
	
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	
	return uf.Count() == 1
}

// DetectCycle 检测无向图中的环
func DetectCycle(n int, edges [][2]int) bool {
	uf := NewUnionFind(n)
	
	for _, edge := range edges {
		if uf.Connected(edge[0], edge[1]) {
			return true // 发现环
		}
		uf.Union(edge[0], edge[1])
	}
	
	return false
}

// 业务应用示例：

// SocialNetwork 社交网络分析
type SocialNetwork struct {
	uf *UnionFind
	userCount int
}

// NewSocialNetwork 创建社交网络
func NewSocialNetwork(userCount int) *SocialNetwork {
	return &SocialNetwork{
		uf: NewUnionFind(userCount),
		userCount: userCount,
	}
}

// AddFriendship 添加好友关系
func (sn *SocialNetwork) AddFriendship(user1, user2 int) bool {
	return sn.uf.Union(user1, user2)
}

// AreFriends 判断是否为朋友（直接或间接）
func (sn *SocialNetwork) AreFriends(user1, user2 int) bool {
	return sn.uf.Connected(user1, user2)
}

// GetFriendGroups 获取所有朋友圈
func (sn *SocialNetwork) GetFriendGroups() [][]int {
	sets := sn.uf.GetSets()
	groups := make([][]int, 0, len(sets))
	
	for _, group := range sets {
		groups = append(groups, group)
	}
	
	return groups
}

// GetLargestFriendGroup 获取最大朋友圈大小
func (sn *SocialNetwork) GetLargestFriendGroup() int {
	maxSize := 0
	for i := 0; i < sn.userCount; i++ {
		size := sn.uf.Size(i)
		if size > maxSize {
			maxSize = size
		}
	}
	return maxSize
}

// NetworkConnectivity 网络连通性分析
type NetworkConnectivity struct {
	uf *UnionFind
	nodeCount int
}

// NewNetworkConnectivity 创建网络连通性分析器
func NewNetworkConnectivity(nodeCount int) *NetworkConnectivity {
	return &NetworkConnectivity{
		uf: NewUnionFind(nodeCount),
		nodeCount: nodeCount,
	}
}

// ConnectNodes 连接两个网络节点
func (nc *NetworkConnectivity) ConnectNodes(node1, node2 int) bool {
	return nc.uf.Union(node1, node2)
}

// AreConnected 判断两个节点是否连通
func (nc *NetworkConnectivity) AreConnected(node1, node2 int) bool {
	return nc.uf.Connected(node1, node2)
}

// GetNetworkComponents 获取网络连通分量
func (nc *NetworkConnectivity) GetNetworkComponents() int {
	return nc.uf.Count()
}

// IsNetworkFullyConnected 判断网络是否完全连通
func (nc *NetworkConnectivity) IsNetworkFullyConnected() bool {
	return nc.uf.Count() == 1
}

// ImageSegmentation 图像分割应用
type ImageSegmentation struct {
	uf *UnionFind
	width, height int
}

// NewImageSegmentation 创建图像分割器
func NewImageSegmentation(width, height int) *ImageSegmentation {
	return &ImageSegmentation{
		uf: NewUnionFind(width * height),
		width: width,
		height: height,
	}
}

// getIndex 将2D坐标转换为1D索引
func (is *ImageSegmentation) getIndex(x, y int) int {
	return y*is.width + x
}

// MergePixels 合并相似像素
func (is *ImageSegmentation) MergePixels(x1, y1, x2, y2 int) bool {
	if x1 < 0 || x1 >= is.width || y1 < 0 || y1 >= is.height ||
		x2 < 0 || x2 >= is.width || y2 < 0 || y2 >= is.height {
		return false
	}
	
	index1 := is.getIndex(x1, y1)
	index2 := is.getIndex(x2, y2)
	
	return is.uf.Union(index1, index2)
}

// GetSegmentCount 获取分割区域数量
func (is *ImageSegmentation) GetSegmentCount() int {
	return is.uf.Count()
}

// GetSegmentSize 获取指定像素所在区域的大小
func (is *ImageSegmentation) GetSegmentSize(x, y int) int {
	if x < 0 || x >= is.width || y < 0 || y >= is.height {
		return 0
	}
	
	index := is.getIndex(x, y)
	return is.uf.Size(index)
}