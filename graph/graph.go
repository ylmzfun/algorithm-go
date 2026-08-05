package graph

import (
	"errors"
	"fmt"
	"strings"
)

// Graph 图数据结构接口
// 思路：定义图的基本操作接口，支持有向图和无向图
// 作用：表示和操作复杂的关系网络
// 业务场景：
// 1. 社交网络：用户关系图，好友推荐
// 2. 地图导航：路径规划，最短路径
// 3. 网络拓扑：计算机网络，互联网结构
// 4. 依赖管理：软件包依赖，任务调度
// 5. 推荐系统：商品关联，内容推荐
// 6. 生物信息：基因网络，蛋白质相互作用
// 7. 金融分析：交易网络，风险传播
type Graph interface {
	AddVertex(vertex interface{}) error
	RemoveVertex(vertex interface{}) error
	AddEdge(from, to interface{}, weight float64) error
	RemoveEdge(from, to interface{}) error
	HasVertex(vertex interface{}) bool
	HasEdge(from, to interface{}) bool
	GetWeight(from, to interface{}) (float64, error)
	GetVertices() []interface{}
	GetNeighbors(vertex interface{}) ([]interface{}, error)
	GetVertexCount() int
	GetEdgeCount() int
	IsDirected() bool
	Clear()
	String() string
}

// Edge 边结构
type Edge struct {
	From   interface{} // 起始顶点
	To     interface{} // 目标顶点
	Weight float64     // 权重
}

// AdjacencyListGraph 邻接表实现的图
type AdjacencyListGraph struct {
	adjList   map[interface{}]map[interface{}]float64 // 邻接表：顶点 -> {邻居顶点 -> 权重}
	vertices  map[interface{}]bool                    // 顶点集合
	edgeCount int                                     // 边的数量
	directed  bool                                    // 是否为有向图
}

// NewAdjacencyListGraph 创建邻接表图
func NewAdjacencyListGraph(directed bool) *AdjacencyListGraph {
	return &AdjacencyListGraph{
		adjList:   make(map[interface{}]map[interface{}]float64),
		vertices:  make(map[interface{}]bool),
		edgeCount: 0,
		directed:  directed,
	}
}

// AddVertex 添加顶点
func (g *AdjacencyListGraph) AddVertex(vertex interface{}) error {
	if g.HasVertex(vertex) {
		return errors.New("vertex already exists")
	}
	
	g.vertices[vertex] = true
	g.adjList[vertex] = make(map[interface{}]float64)
	return nil
}

// RemoveVertex 删除顶点
func (g *AdjacencyListGraph) RemoveVertex(vertex interface{}) error {
	if !g.HasVertex(vertex) {
		return errors.New("vertex not found")
	}
	
	// 删除所有与该顶点相关的边
	for neighbor := range g.adjList[vertex] {
		g.RemoveEdge(vertex, neighbor)
	}
	
	// 删除其他顶点到该顶点的边
	for v := range g.vertices {
		if v != vertex {
			g.RemoveEdge(v, vertex)
		}
	}
	
	// 删除顶点
	delete(g.vertices, vertex)
	delete(g.adjList, vertex)
	
	return nil
}

// AddEdge 添加边
func (g *AdjacencyListGraph) AddEdge(from, to interface{}, weight float64) error {
	if !g.HasVertex(from) {
		return errors.New("from vertex not found")
	}
	if !g.HasVertex(to) {
		return errors.New("to vertex not found")
	}
	
	// 检查边是否已存在
	if g.HasEdge(from, to) {
		// 更新权重
		g.adjList[from][to] = weight
		if !g.directed {
			g.adjList[to][from] = weight
		}
		return nil
	}
	
	// 添加边
	g.adjList[from][to] = weight
	g.edgeCount++
	
	// 如果是无向图，添加反向边
	if !g.directed {
		g.adjList[to][from] = weight
	}
	
	return nil
}

// RemoveEdge 删除边
func (g *AdjacencyListGraph) RemoveEdge(from, to interface{}) error {
	if !g.HasVertex(from) || !g.HasVertex(to) {
		return errors.New("vertex not found")
	}
	
	if !g.HasEdge(from, to) {
		return errors.New("edge not found")
	}
	
	// 删除边
	delete(g.adjList[from], to)
	g.edgeCount--
	
	// 如果是无向图，删除反向边
	if !g.directed {
		delete(g.adjList[to], from)
	}
	
	return nil
}

// HasVertex 检查顶点是否存在
func (g *AdjacencyListGraph) HasVertex(vertex interface{}) bool {
	return g.vertices[vertex]
}

// HasEdge 检查边是否存在
func (g *AdjacencyListGraph) HasEdge(from, to interface{}) bool {
	if !g.HasVertex(from) || !g.HasVertex(to) {
		return false
	}
	_, exists := g.adjList[from][to]
	return exists
}

// GetWeight 获取边的权重
func (g *AdjacencyListGraph) GetWeight(from, to interface{}) (float64, error) {
	if !g.HasEdge(from, to) {
		return 0, errors.New("edge not found")
	}
	return g.adjList[from][to], nil
}

// GetVertices 获取所有顶点
func (g *AdjacencyListGraph) GetVertices() []interface{} {
	vertices := make([]interface{}, 0, len(g.vertices))
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// GetNeighbors 获取顶点的邻居
func (g *AdjacencyListGraph) GetNeighbors(vertex interface{}) ([]interface{}, error) {
	if !g.HasVertex(vertex) {
		return nil, errors.New("vertex not found")
	}
	
	neighbors := make([]interface{}, 0, len(g.adjList[vertex]))
	for neighbor := range g.adjList[vertex] {
		neighbors = append(neighbors, neighbor)
	}
	return neighbors, nil
}

// GetVertexCount 获取顶点数量
func (g *AdjacencyListGraph) GetVertexCount() int {
	return len(g.vertices)
}

// GetEdgeCount 获取边数量
func (g *AdjacencyListGraph) GetEdgeCount() int {
	return g.edgeCount
}

// IsDirected 是否为有向图
func (g *AdjacencyListGraph) IsDirected() bool {
	return g.directed
}

// Clear 清空图
func (g *AdjacencyListGraph) Clear() {
	g.adjList = make(map[interface{}]map[interface{}]float64)
	g.vertices = make(map[interface{}]bool)
	g.edgeCount = 0
}

// String 字符串表示
func (g *AdjacencyListGraph) String() string {
	if len(g.vertices) == 0 {
		return "Graph{}"
	}
	
	var edges []string
	for from, neighbors := range g.adjList {
		for to, weight := range neighbors {
			if g.directed || fmt.Sprintf("%v", from) <= fmt.Sprintf("%v", to) {
				edges = append(edges, fmt.Sprintf("%v->%v(%.1f)", from, to, weight))
			}
		}
	}
	
	graphType := "Directed"
	if !g.directed {
		graphType = "Undirected"
	}
	
	return fmt.Sprintf("%s Graph{vertices: %d, edges: %d, connections: [%s]}", 
		graphType, g.GetVertexCount(), g.GetEdgeCount(), strings.Join(edges, ", "))
}

// AdjacencyMatrixGraph 邻接矩阵实现的图
type AdjacencyMatrixGraph struct {
	matrix      [][]float64            // 邻接矩阵
	vertexMap   map[interface{}]int    // 顶点到索引的映射
	indexMap    map[int]interface{}    // 索引到顶点的映射
	vertexCount int                    // 顶点数量
	edgeCount   int                    // 边数量
	capacity    int                    // 矩阵容量
	directed    bool                   // 是否为有向图
	noEdgeValue float64               // 表示无边的值
}

// NewAdjacencyMatrixGraph 创建邻接矩阵图
func NewAdjacencyMatrixGraph(capacity int, directed bool) *AdjacencyMatrixGraph {
	if capacity <= 0 {
		capacity = 10
	}
	
	noEdgeValue := -1.0 // 使用-1表示无边
	matrix := make([][]float64, capacity)
	for i := range matrix {
		matrix[i] = make([]float64, capacity)
		for j := range matrix[i] {
			matrix[i][j] = noEdgeValue
		}
	}
	
	return &AdjacencyMatrixGraph{
		matrix:      matrix,
		vertexMap:   make(map[interface{}]int),
		indexMap:    make(map[int]interface{}),
		vertexCount: 0,
		edgeCount:   0,
		capacity:    capacity,
		directed:    directed,
		noEdgeValue: noEdgeValue,
	}
}

// AddVertex 添加顶点
func (g *AdjacencyMatrixGraph) AddVertex(vertex interface{}) error {
	if g.HasVertex(vertex) {
		return errors.New("vertex already exists")
	}
	
	if g.vertexCount >= g.capacity {
		return errors.New("graph capacity exceeded")
	}
	
	index := g.vertexCount
	g.vertexMap[vertex] = index
	g.indexMap[index] = vertex
	g.vertexCount++
	
	return nil
}

// RemoveVertex 删除顶点
func (g *AdjacencyMatrixGraph) RemoveVertex(vertex interface{}) error {
	if !g.HasVertex(vertex) {
		return errors.New("vertex not found")
	}
	
	index := g.vertexMap[vertex]
	
	// 删除所有与该顶点相关的边
	for i := 0; i < g.vertexCount; i++ {
		if g.matrix[index][i] != g.noEdgeValue {
			g.edgeCount--
			g.matrix[index][i] = g.noEdgeValue
			if !g.directed {
				g.matrix[i][index] = g.noEdgeValue
			}
		}
		if g.matrix[i][index] != g.noEdgeValue && g.directed {
			g.edgeCount--
			g.matrix[i][index] = g.noEdgeValue
		}
	}
	
	// 移动最后一个顶点到删除位置
	lastIndex := g.vertexCount - 1
	if index != lastIndex {
		lastVertex := g.indexMap[lastIndex]
		
		// 更新映射
		g.vertexMap[lastVertex] = index
		g.indexMap[index] = lastVertex
		
		// 移动矩阵行和列
		for i := 0; i < g.vertexCount; i++ {
			g.matrix[index][i] = g.matrix[lastIndex][i]
			g.matrix[i][index] = g.matrix[i][lastIndex]
			g.matrix[lastIndex][i] = g.noEdgeValue
			g.matrix[i][lastIndex] = g.noEdgeValue
		}
	}
	
	// 删除顶点
	delete(g.vertexMap, vertex)
	delete(g.indexMap, lastIndex)
	g.vertexCount--
	
	return nil
}

// AddEdge 添加边
func (g *AdjacencyMatrixGraph) AddEdge(from, to interface{}, weight float64) error {
	if !g.HasVertex(from) {
		return errors.New("from vertex not found")
	}
	if !g.HasVertex(to) {
		return errors.New("to vertex not found")
	}
	
	fromIndex := g.vertexMap[from]
	toIndex := g.vertexMap[to]
	
	// 检查边是否已存在
	if g.matrix[fromIndex][toIndex] != g.noEdgeValue {
		// 更新权重
		g.matrix[fromIndex][toIndex] = weight
		if !g.directed {
			g.matrix[toIndex][fromIndex] = weight
		}
		return nil
	}
	
	// 添加边
	g.matrix[fromIndex][toIndex] = weight
	g.edgeCount++
	
	// 如果是无向图，添加反向边
	if !g.directed {
		g.matrix[toIndex][fromIndex] = weight
	}
	
	return nil
}

// RemoveEdge 删除边
func (g *AdjacencyMatrixGraph) RemoveEdge(from, to interface{}) error {
	if !g.HasVertex(from) || !g.HasVertex(to) {
		return errors.New("vertex not found")
	}
	
	fromIndex := g.vertexMap[from]
	toIndex := g.vertexMap[to]
	
	if g.matrix[fromIndex][toIndex] == g.noEdgeValue {
		return errors.New("edge not found")
	}
	
	// 删除边
	g.matrix[fromIndex][toIndex] = g.noEdgeValue
	g.edgeCount--
	
	// 如果是无向图，删除反向边
	if !g.directed {
		g.matrix[toIndex][fromIndex] = g.noEdgeValue
	}
	
	return nil
}

// HasVertex 检查顶点是否存在
func (g *AdjacencyMatrixGraph) HasVertex(vertex interface{}) bool {
	_, exists := g.vertexMap[vertex]
	return exists
}

// HasEdge 检查边是否存在
func (g *AdjacencyMatrixGraph) HasEdge(from, to interface{}) bool {
	if !g.HasVertex(from) || !g.HasVertex(to) {
		return false
	}
	
	fromIndex := g.vertexMap[from]
	toIndex := g.vertexMap[to]
	return g.matrix[fromIndex][toIndex] != g.noEdgeValue
}

// GetWeight 获取边的权重
func (g *AdjacencyMatrixGraph) GetWeight(from, to interface{}) (float64, error) {
	if !g.HasEdge(from, to) {
		return 0, errors.New("edge not found")
	}
	
	fromIndex := g.vertexMap[from]
	toIndex := g.vertexMap[to]
	return g.matrix[fromIndex][toIndex], nil
}

// GetVertices 获取所有顶点
func (g *AdjacencyMatrixGraph) GetVertices() []interface{} {
	vertices := make([]interface{}, 0, g.vertexCount)
	for i := 0; i < g.vertexCount; i++ {
		vertices = append(vertices, g.indexMap[i])
	}
	return vertices
}

// GetNeighbors 获取顶点的邻居
func (g *AdjacencyMatrixGraph) GetNeighbors(vertex interface{}) ([]interface{}, error) {
	if !g.HasVertex(vertex) {
		return nil, errors.New("vertex not found")
	}
	
	index := g.vertexMap[vertex]
	neighbors := make([]interface{}, 0)
	
	for i := 0; i < g.vertexCount; i++ {
		if g.matrix[index][i] != g.noEdgeValue {
			neighbors = append(neighbors, g.indexMap[i])
		}
	}
	
	return neighbors, nil
}

// GetVertexCount 获取顶点数量
func (g *AdjacencyMatrixGraph) GetVertexCount() int {
	return g.vertexCount
}

// GetEdgeCount 获取边数量
func (g *AdjacencyMatrixGraph) GetEdgeCount() int {
	return g.edgeCount
}

// IsDirected 是否为有向图
func (g *AdjacencyMatrixGraph) IsDirected() bool {
	return g.directed
}

// Clear 清空图
func (g *AdjacencyMatrixGraph) Clear() {
	for i := 0; i < g.capacity; i++ {
		for j := 0; j < g.capacity; j++ {
			g.matrix[i][j] = g.noEdgeValue
		}
	}
	g.vertexMap = make(map[interface{}]int)
	g.indexMap = make(map[int]interface{})
	g.vertexCount = 0
	g.edgeCount = 0
}

// String 字符串表示
func (g *AdjacencyMatrixGraph) String() string {
	if g.vertexCount == 0 {
		return "Graph{}"
	}
	
	var edges []string
	for i := 0; i < g.vertexCount; i++ {
		for j := 0; j < g.vertexCount; j++ {
			if g.matrix[i][j] != g.noEdgeValue {
				if g.directed || i <= j {
					from := g.indexMap[i]
					to := g.indexMap[j]
					edges = append(edges, fmt.Sprintf("%v->%v(%.1f)", from, to, g.matrix[i][j]))
				}
			}
		}
	}
	
	graphType := "Directed"
	if !g.directed {
		graphType = "Undirected"
	}
	
	return fmt.Sprintf("%s Graph{vertices: %d, edges: %d, connections: [%s]}", 
		graphType, g.GetVertexCount(), g.GetEdgeCount(), strings.Join(edges, ", "))
}

// 图算法实现

// DFS 深度优先搜索
func DFS(graph Graph, start interface{}, visit func(interface{})) error {
	if !graph.HasVertex(start) {
		return errors.New("start vertex not found")
	}
	
	visited := make(map[interface{}]bool)
	stack := []interface{}{start}
	
	for len(stack) > 0 {
		// 弹出栈顶元素
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if !visited[current] {
			visited[current] = true
			visit(current)
			
			// 将邻居加入栈
			neighbors, _ := graph.GetNeighbors(current)
			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	
	return nil
}

// BFS 广度优先搜索
func BFS(graph Graph, start interface{}, visit func(interface{})) error {
	if !graph.HasVertex(start) {
		return errors.New("start vertex not found")
	}
	
	visited := make(map[interface{}]bool)
	queue := []interface{}{start}
	visited[start] = true
	
	for len(queue) > 0 {
		// 出队
		current := queue[0]
		queue = queue[1:]
		
		visit(current)
		
		// 将未访问的邻居加入队列
		neighbors, _ := graph.GetNeighbors(current)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	
	return nil
}

// HasPath 检查两个顶点之间是否存在路径
func HasPath(graph Graph, from, to interface{}) bool {
	if !graph.HasVertex(from) || !graph.HasVertex(to) {
		return false
	}
	
	if from == to {
		return true
	}
	
	found := false
	DFS(graph, from, func(vertex interface{}) {
		if vertex == to {
			found = true
		}
	})
	
	return found
}

// GetConnectedComponents 获取连通分量
func GetConnectedComponents(graph Graph) [][]interface{} {
	visited := make(map[interface{}]bool)
	components := make([][]interface{}, 0)
	
	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			component := make([]interface{}, 0)
			DFS(graph, vertex, func(v interface{}) {
				visited[v] = true
				component = append(component, v)
			})
			components = append(components, component)
		}
	}
	
	return components
}

// IsConnected 检查图是否连通
func IsConnected(graph Graph) bool {
	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return true
	}
	
	visitedCount := 0
	BFS(graph, vertices[0], func(interface{}) {
		visitedCount++
	})
	
	return visitedCount == graph.GetVertexCount()
}

// TopologicalSort 拓扑排序（仅适用于有向无环图）
func TopologicalSort(graph Graph) ([]interface{}, error) {
	if !graph.IsDirected() {
		return nil, errors.New("topological sort only applies to directed graphs")
	}
	
	// 计算每个顶点的入度
	inDegree := make(map[interface{}]int)
	for _, vertex := range graph.GetVertices() {
		inDegree[vertex] = 0
	}
	
	for _, vertex := range graph.GetVertices() {
		neighbors, _ := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}
	
	// 找到所有入度为0的顶点
	queue := make([]interface{}, 0)
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}
	
	result := make([]interface{}, 0)
	
	// Kahn算法
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		
		neighbors, _ := graph.GetNeighbors(current)
		for _, neighbor := range neighbors {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	
	// 检查是否存在环
	if len(result) != graph.GetVertexCount() {
		return nil, errors.New("graph contains cycle")
	}
	
	return result, nil
}

// 业务应用示例：
// 1. 社交网络分析：好友关系图，影响力传播
// 2. 路径规划：地图导航，最短路径算法
// 3. 依赖管理：软件包依赖解析，构建顺序
// 4. 网络分析：计算机网络拓扑，故障检测
// 5. 推荐系统：用户-商品关系图，协同过滤
// 6. 工作流管理：任务依赖图，调度优化
// 7. 知识图谱：实体关系建模，语义搜索
// 8. 生物信息学：基因调控网络，蛋白质相互作用