<!--
 * @Author: ylmzfun ylmzfun@foxmail.com
 * @Date: 2025-08-07 14:49:34
 * @LastEditors: ylmzfun ylmzfun@foxmail.com
 * @LastEditTime: 2025-08-07 15:29:42
 * @FilePath: /go-algorithm/README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Go 数据结构实现

本项目实现了常见的数据结构与算法，使用Go语言编写，包含详细的注释和使用示例。

## 项目结构

```
go-algorithm/
├── README.md
├── go.mod
├── main.go
├── array/          # 数组
├── linkedlist/     # 链表
├── stack/          # 栈
├── queue/          # 队列
├── tree/           # 树结构
├── graph/          # 图
├── hash/           # 哈希表
├── heap/           # 堆
├── trie/           # 字典树
├── unionfind/      # 并查集
├── bloomfilter/    # 布隆过滤器
├── sort/           # 排序算法
└── search/         # 搜索算法
```

## 已实现的数据结构与算法

### 基础数据结构
- [x] 动态数组 (Dynamic Array)
- [x] 单向链表 (Singly Linked List)
- [x] 双向链表 (Doubly Linked List)
- [x] 栈 (Stack) — 切片实现 & 链表实现
- [x] 队列 (Queue) — 循环队列 & 链表队列
- [x] 双端队列 (Deque)

### 树结构
- [x] 二叉搜索树 (Binary Search Tree)
- [x] AVL树 (AVL Tree)
- [x] 红黑树 (Red-Black Tree)
- [x] B树 (B-Tree)
- [x] 字典树 (Trie)

### 图结构
- [x] 邻接表图 (Adjacency List Graph)
- [x] 邻接矩阵图 (Adjacency Matrix Graph)

### 高级数据结构
- [x] 哈希表 (Hash Table) — 链地址法 & 开放地址法
- [x] 最小堆 (Min Heap)
- [x] 最大堆 (Max Heap)
- [x] 优先级队列 (Priority Queue)
- [x] 并查集 (Union Find)
- [x] 布隆过滤器 (Bloom Filter)

### 排序算法
- [x] 冒泡排序 (Bubble Sort)
- [x] 选择排序 (Selection Sort)
- [x] 插入排序 (Insertion Sort)
- [x] 希尔排序 (Shell Sort)
- [x] 归并排序 (Merge Sort)
- [x] 快速排序 (Quick Sort) — 三数取中 + 小数组插入排序优化
- [x] 堆排序 (Heap Sort)
- [x] 计数排序 (Counting Sort)
- [x] 基数排序 (Radix Sort)
- [x] 桶排序 (Bucket Sort)
- [x] 快速选择 (Quick Select) — 第 K 小元素

### 搜索算法
- [x] 线性搜索 (Linear Search)
- [x] 二分搜索 (Binary Search) — 迭代 & 递归
- [x] 二分搜索变体 — 首次/末次出现位置、LowerBound、UpperBound
- [x] 跳跃搜索 (Jump Search)
- [x] 插值搜索 (Interpolation Search)
- [x] 指数搜索 (Exponential Search)
- [x] 斐波那契搜索 (Fibonacci Search)
- [x] 三分搜索 (Ternary Search)

### 工具函数
- [x] 括号匹配 (`stack.IsValidParentheses`)
- [x] 后缀表达式求值 (`stack.EvaluatePostfix`)
- [x] 环检测 (`unionfind.DetectCycle`)
- [x] 社交网络分析 (`unionfind.SocialNetwork`)
- [x] Kruskal 最小生成树 (`unionfind.KruskalMST`)
- [x] 合并 K 个有序数组 (`heap.MergeKSortedArrays`)
- [x] 第 K 大元素 (`heap.FindKthLargest`)
- [x] 数组反转 (`sort.Reverse`)
- [x] 排序校验 (`sort.IsSorted`)

## 使用方法

每个数据结构和算法都包含：
1. 完整的实现代码
2. 详细的注释说明（思路、作用、业务场景）
3. 时间复杂度分析
4. 单元测试
5. `main.go` 中提供了完整的演示

## 运行测试

```bash
go test ./...
```

## 运演示

```bash
go run main.go
```

## 业务应用场景

每个数据结构的实现都包含了详细的业务应用场景说明，帮助理解在实际项目中的使用方式。覆盖场景包括：

- **数据库系统**：B树/B+树索引、哈希索引、内存缓存
- **操作系统**：进程调度（堆）、页面置换（链表）、文件系统（B树）
- **网络系统**：路由表（哈希）、拓扑排序（图）、IP 前缀匹配（Trie）
- **搜索引擎**：倒排索引（哈希）、URL 去重（布隆过滤器）、自动补全（Trie）
- **社交网络**：好友关系（并查集）、推荐系统（图）、影响力传播
- **金融系统**：订单匹配（堆）、风险传播（图）、交易去重（布隆过滤器）
- **游戏开发**：碰撞检测（树）、AI 寻路（图+A*）、事件队列
- **编译器**：符号表（哈希+树）、语法分析（树）、依赖解析（图）
