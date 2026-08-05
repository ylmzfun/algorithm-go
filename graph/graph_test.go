package graph

import (
	"reflect"
	"sort"
	"testing"
)

// TestAdjacencyListGraphBasicOperations 测试邻接表图基本操作
func TestAdjacencyListGraphBasicOperations(t *testing.T) {
	// 测试有向图
	graph := NewAdjacencyListGraph(true)

	// 测试添加顶点
	err := graph.AddVertex("A")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = graph.AddVertex("B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = graph.AddVertex("C")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if graph.GetVertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", graph.GetVertexCount())
	}

	// 测试重复添加顶点
	err = graph.AddVertex("A")
	if err == nil {
		t.Error("Expected error for duplicate vertex")
	}

	// 测试HasVertex
	if !graph.HasVertex("A") {
		t.Error("Expected graph to have vertex A")
	}
	if graph.HasVertex("D") {
		t.Error("Expected graph not to have vertex D")
	}

	// 测试添加边
	err = graph.AddEdge("A", "B", 1.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = graph.AddEdge("B", "C", 2.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = graph.AddEdge("A", "C", 3.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if graph.GetEdgeCount() != 3 {
		t.Errorf("Expected 3 edges, got %d", graph.GetEdgeCount())
	}

	// 测试HasEdge
	if !graph.HasEdge("A", "B") {
		t.Error("Expected edge A->B")
	}
	if graph.HasEdge("B", "A") {
		t.Error("Expected no edge B->A in directed graph")
	}

	// 测试GetWeight
	weight, err := graph.GetWeight("A", "B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if weight != 1.0 {
		t.Errorf("Expected weight 1.0, got %f", weight)
	}

	// 测试GetNeighbors
	neighbors, err := graph.GetNeighbors("A")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors, got %d", len(neighbors))
	}
}

// TestAdjacencyListGraphUndirected 测试无向图
func TestAdjacencyListGraphUndirected(t *testing.T) {
	graph := NewAdjacencyListGraph(false)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")

	// 添加无向边
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("B", "C", 2.0)

	// 无向图应该有双向边
	if !graph.HasEdge("A", "B") {
		t.Error("Expected edge A->B")
	}
	if !graph.HasEdge("B", "A") {
		t.Error("Expected edge B->A in undirected graph")
	}

	// 边数应该只计算一次
	if graph.GetEdgeCount() != 2 {
		t.Errorf("Expected 2 edges, got %d", graph.GetEdgeCount())
	}
}

// TestAdjacencyListGraphRemoveOperations 测试删除操作
func TestAdjacencyListGraphRemoveOperations(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("B", "C", 2.0)
	graph.AddEdge("A", "C", 3.0)

	// 测试删除边
	err := graph.RemoveEdge("A", "B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if graph.HasEdge("A", "B") {
		t.Error("Expected edge A->B to be removed")
	}
	if graph.GetEdgeCount() != 2 {
		t.Errorf("Expected 2 edges, got %d", graph.GetEdgeCount())
	}

	// 测试删除不存在的边
	err = graph.RemoveEdge("A", "B")
	if err == nil {
		t.Error("Expected error for removing non-existent edge")
	}

	// 测试删除顶点
	err = graph.RemoveVertex("B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if graph.HasVertex("B") {
		t.Error("Expected vertex B to be removed")
	}
	if graph.GetVertexCount() != 2 {
		t.Errorf("Expected 2 vertices, got %d", graph.GetVertexCount())
	}
	if graph.GetEdgeCount() != 1 {
		t.Errorf("Expected 1 edge, got %d", graph.GetEdgeCount())
	}
}

// TestAdjacencyMatrixGraphBasicOperations 测试邻接矩阵图基本操作
func TestAdjacencyMatrixGraphBasicOperations(t *testing.T) {
	graph := NewAdjacencyMatrixGraph(10, true)

	// 测试添加顶点
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")

	if graph.GetVertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", graph.GetVertexCount())
	}

	// 测试添加边
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("B", "C", 2.0)
	graph.AddEdge("A", "C", 3.0)

	if graph.GetEdgeCount() != 3 {
		t.Errorf("Expected 3 edges, got %d", graph.GetEdgeCount())
	}

	// 测试HasEdge
	if !graph.HasEdge("A", "B") {
		t.Error("Expected edge A->B")
	}
	if graph.HasEdge("B", "A") {
		t.Error("Expected no edge B->A in directed graph")
	}

	// 测试GetWeight
	weight, err := graph.GetWeight("A", "B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if weight != 1.0 {
		t.Errorf("Expected weight 1.0, got %f", weight)
	}
}

// TestAdjacencyMatrixGraphCapacity 测试邻接矩阵图容量限制
func TestAdjacencyMatrixGraphCapacity(t *testing.T) {
	graph := NewAdjacencyMatrixGraph(2, true)

	graph.AddVertex("A")
	graph.AddVertex("B")

	// 超出容量
	err := graph.AddVertex("C")
	if err == nil {
		t.Error("Expected error for exceeding capacity")
	}
}

// TestGraphClear 测试清空图
func TestGraphClear(t *testing.T) {
	// 测试邻接表图
	graph1 := NewAdjacencyListGraph(true)
	graph1.AddVertex("A")
	graph1.AddVertex("B")
	graph1.AddEdge("A", "B", 1.0)

	graph1.Clear()
	if graph1.GetVertexCount() != 0 {
		t.Errorf("Expected 0 vertices after clear, got %d", graph1.GetVertexCount())
	}
	if graph1.GetEdgeCount() != 0 {
		t.Errorf("Expected 0 edges after clear, got %d", graph1.GetEdgeCount())
	}

	// 测试邻接矩阵图
	graph2 := NewAdjacencyMatrixGraph(10, true)
	graph2.AddVertex("A")
	graph2.AddVertex("B")
	graph2.AddEdge("A", "B", 1.0)

	graph2.Clear()
	if graph2.GetVertexCount() != 0 {
		t.Errorf("Expected 0 vertices after clear, got %d", graph2.GetVertexCount())
	}
	if graph2.GetEdgeCount() != 0 {
		t.Errorf("Expected 0 edges after clear, got %d", graph2.GetEdgeCount())
	}
}

// TestDFS 测试深度优先搜索
func TestDFS(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("A", "C", 1.0)
	graph.AddEdge("B", "D", 1.0)
	graph.AddEdge("C", "D", 1.0)

	visited := make([]interface{}, 0)
	err := DFS(graph, "A", func(vertex interface{}) {
		visited = append(visited, vertex)
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(visited) != 4 {
		t.Errorf("Expected 4 visited vertices, got %d", len(visited))
	}
	if visited[0] != "A" {
		t.Errorf("Expected first visited vertex to be A, got %v", visited[0])
	}

	// 测试不存在的起始顶点
	err = DFS(graph, "E", func(vertex interface{}) {})
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

// TestBFS 测试广度优先搜索
func TestBFS(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("A", "C", 1.0)
	graph.AddEdge("B", "D", 1.0)
	graph.AddEdge("C", "D", 1.0)

	visited := make([]interface{}, 0)
	err := BFS(graph, "A", func(vertex interface{}) {
		visited = append(visited, vertex)
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(visited) != 4 {
		t.Errorf("Expected 4 visited vertices, got %d", len(visited))
	}
	if visited[0] != "A" {
		t.Errorf("Expected first visited vertex to be A, got %v", visited[0])
	}

	// 测试不存在的起始顶点
	err = BFS(graph, "E", func(vertex interface{}) {})
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

// TestHasPath 测试路径检查
func TestHasPath(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("B", "C", 1.0)
	// D是孤立顶点

	// 测试存在路径
	if !HasPath(graph, "A", "C") {
		t.Error("Expected path from A to C")
	}

	// 测试不存在路径
	if HasPath(graph, "A", "D") {
		t.Error("Expected no path from A to D")
	}

	// 测试自身路径
	if !HasPath(graph, "A", "A") {
		t.Error("Expected path from A to A")
	}

	// 测试不存在的顶点
	if HasPath(graph, "A", "E") {
		t.Error("Expected no path to non-existent vertex")
	}
}

// TestGetConnectedComponents 测试连通分量
func TestGetConnectedComponents(t *testing.T) {
	graph := NewAdjacencyListGraph(false) // 无向图

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("B", "C", 1.0)
	// D和E是孤立顶点

	components := GetConnectedComponents(graph)

	if len(components) != 3 {
		t.Errorf("Expected 3 connected components, got %d", len(components))
	}

	// 验证第一个连通分量包含A、B、C
	firstComponent := components[0]
	if len(firstComponent) != 3 {
		t.Errorf("Expected first component to have 3 vertices, got %d", len(firstComponent))
	}
}

// TestIsConnected 测试连通性检查
func TestIsConnected(t *testing.T) {
	// 测试连通图
	graph1 := NewAdjacencyListGraph(false)
	graph1.AddVertex("A")
	graph1.AddVertex("B")
	graph1.AddVertex("C")
	graph1.AddEdge("A", "B", 1.0)
	graph1.AddEdge("B", "C", 1.0)

	if !IsConnected(graph1) {
		t.Error("Expected graph to be connected")
	}

	// 测试非连通图
	graph2 := NewAdjacencyListGraph(false)
	graph2.AddVertex("A")
	graph2.AddVertex("B")
	graph2.AddVertex("C")
	graph2.AddVertex("D")
	graph2.AddEdge("A", "B", 1.0)
	// C和D是孤立顶点

	if IsConnected(graph2) {
		t.Error("Expected graph to be disconnected")
	}

	// 测试空图
	graph3 := NewAdjacencyListGraph(false)
	if !IsConnected(graph3) {
		t.Error("Expected empty graph to be connected")
	}
}

// TestTopologicalSort 测试拓扑排序
func TestTopologicalSort(t *testing.T) {
	// 测试有向无环图
	graph := NewAdjacencyListGraph(true)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddEdge("A", "B", 1.0)
	graph.AddEdge("A", "C", 1.0)
	graph.AddEdge("B", "D", 1.0)
	graph.AddEdge("C", "D", 1.0)

	result, err := TopologicalSort(graph)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 4 {
		t.Errorf("Expected 4 vertices in topological order, got %d", len(result))
	}
	if result[0] != "A" {
		t.Errorf("Expected A to be first in topological order, got %v", result[0])
	}
	if result[len(result)-1] != "D" {
		t.Errorf("Expected D to be last in topological order, got %v", result[len(result)-1])
	}

	// 测试有环图
	graphWithCycle := NewAdjacencyListGraph(true)
	graphWithCycle.AddVertex("A")
	graphWithCycle.AddVertex("B")
	graphWithCycle.AddVertex("C")
	graphWithCycle.AddEdge("A", "B", 1.0)
	graphWithCycle.AddEdge("B", "C", 1.0)
	graphWithCycle.AddEdge("C", "A", 1.0) // 形成环

	_, err = TopologicalSort(graphWithCycle)
	if err == nil {
		t.Error("Expected error for graph with cycle")
	}

	// 测试无向图
	undirectedGraph := NewAdjacencyListGraph(false)
	undirectedGraph.AddVertex("A")
	undirectedGraph.AddVertex("B")
	undirectedGraph.AddEdge("A", "B", 1.0)

	_, err = TopologicalSort(undirectedGraph)
	if err == nil {
		t.Error("Expected error for undirected graph")
	}
}

// TestGraphString 测试字符串表示
func TestGraphString(t *testing.T) {
	// 测试空图
	graph := NewAdjacencyListGraph(true)
	str := graph.String()
	if str != "Graph{}" {
		t.Errorf("Expected 'Graph{}', got '%s'", str)
	}

	// 测试非空图
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddEdge("A", "B", 1.0)

	str = graph.String()
	if !contains(str, "Directed Graph") {
		t.Errorf("Expected string to contain 'Directed Graph', got '%s'", str)
	}
	if !contains(str, "vertices: 2") {
		t.Errorf("Expected string to contain 'vertices: 2', got '%s'", str)
	}
	if !contains(str, "edges: 1") {
		t.Errorf("Expected string to contain 'edges: 1', got '%s'", str)
	}
}

// TestUpdateEdgeWeight 测试更新边权重
func TestUpdateEdgeWeight(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddEdge("A", "B", 1.0)

	// 更新边权重
	graph.AddEdge("A", "B", 2.0)

	// 边数不应该增加
	if graph.GetEdgeCount() != 1 {
		t.Errorf("Expected 1 edge, got %d", graph.GetEdgeCount())
	}

	// 权重应该更新
	weight, err := graph.GetWeight("A", "B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if weight != 2.0 {
		t.Errorf("Expected weight 2.0, got %f", weight)
	}
}

// TestGetVertices 测试获取所有顶点
func TestGetVertices(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	expectedVertices := []interface{}{"A", "B", "C"}
	for _, vertex := range expectedVertices {
		graph.AddVertex(vertex)
	}

	vertices := graph.GetVertices()
	if len(vertices) != len(expectedVertices) {
		t.Errorf("Expected %d vertices, got %d", len(expectedVertices), len(vertices))
	}

	// 将结果排序以便比较
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].(string) < vertices[j].(string)
	})
	sort.Slice(expectedVertices, func(i, j int) bool {
		return expectedVertices[i].(string) < expectedVertices[j].(string)
	})

	if !reflect.DeepEqual(vertices, expectedVertices) {
		t.Errorf("Expected vertices %v, got %v", expectedVertices, vertices)
	}
}

// TestErrorCases 测试错误情况
func TestErrorCases(t *testing.T) {
	graph := NewAdjacencyListGraph(true)

	// 测试在不存在的顶点之间添加边
	err := graph.AddEdge("A", "B", 1.0)
	if err == nil {
		t.Error("Expected error for adding edge between non-existent vertices")
	}

	graph.AddVertex("A")
	err = graph.AddEdge("A", "B", 1.0)
	if err == nil {
		t.Error("Expected error for adding edge to non-existent vertex")
	}

	// 测试获取不存在顶点的邻居
	_, err = graph.GetNeighbors("B")
	if err == nil {
		t.Error("Expected error for getting neighbors of non-existent vertex")
	}

	// 测试获取不存在边的权重
	graph.AddVertex("B")
	_, err = graph.GetWeight("A", "B")
	if err == nil {
		t.Error("Expected error for getting weight of non-existent edge")
	}

	// 测试删除不存在的顶点
	err = graph.RemoveVertex("C")
	if err == nil {
		t.Error("Expected error for removing non-existent vertex")
	}
}

// 辅助函数：检查字符串是否包含子串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || 
			s[len(s)-len(substr):] == substr || 
			containsHelper(s, substr))))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// BenchmarkAdjacencyListAddVertex 基准测试：邻接表添加顶点
func BenchmarkAdjacencyListAddVertex(b *testing.B) {
	graph := NewAdjacencyListGraph(true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		graph.AddVertex(i)
	}
}

// BenchmarkAdjacencyListAddEdge 基准测试：邻接表添加边
func BenchmarkAdjacencyListAddEdge(b *testing.B) {
	graph := NewAdjacencyListGraph(true)

	// 预先添加顶点
	for i := 0; i < 1000; i++ {
		graph.AddVertex(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		from := i % 1000
		to := (i + 1) % 1000
		graph.AddEdge(from, to, 1.0)
	}
}

// BenchmarkAdjacencyMatrixAddVertex 基准测试：邻接矩阵添加顶点
func BenchmarkAdjacencyMatrixAddVertex(b *testing.B) {
	graph := NewAdjacencyMatrixGraph(b.N, true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		graph.AddVertex(i)
	}
}

// BenchmarkAdjacencyMatrixAddEdge 基准测试：邻接矩阵添加边
func BenchmarkAdjacencyMatrixAddEdge(b *testing.B) {
	graph := NewAdjacencyMatrixGraph(1000, true)

	// 预先添加顶点
	for i := 0; i < 1000; i++ {
		graph.AddVertex(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		from := i % 1000
		to := (i + 1) % 1000
		graph.AddEdge(from, to, 1.0)
	}
}

// BenchmarkDFS 基准测试：深度优先搜索
func BenchmarkDFS(b *testing.B) {
	graph := NewAdjacencyListGraph(true)

	// 创建一个较大的图
	for i := 0; i < 1000; i++ {
		graph.AddVertex(i)
	}
	for i := 0; i < 999; i++ {
		graph.AddEdge(i, i+1, 1.0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DFS(graph, 0, func(interface{}) {})
	}
}

// BenchmarkBFS 基准测试：广度优先搜索
func BenchmarkBFS(b *testing.B) {
	graph := NewAdjacencyListGraph(true)

	// 创建一个较大的图
	for i := 0; i < 1000; i++ {
		graph.AddVertex(i)
	}
	for i := 0; i < 999; i++ {
		graph.AddEdge(i, i+1, 1.0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BFS(graph, 0, func(interface{}) {})
	}
}