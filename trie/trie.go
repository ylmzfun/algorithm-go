package trie

import (
	"strings"
)

// TrieNode 字典树节点
type TrieNode struct {
	children map[rune]*TrieNode // 子节点映射
	isEnd    bool               // 是否为单词结尾
	value    interface{}        // 存储的值（可选）
	count    int                // 经过此节点的单词数量
}

// NewTrieNode 创建新的字典树节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		value:    nil,
		count:    0,
	}
}

// Trie 字典树数据结构
// 思路：使用树形结构存储字符串，每个节点代表一个字符，从根到叶子的路径构成完整单词
// 作用：高效地进行字符串的插入、查找、删除和前缀匹配
// 业务场景：
// 1. 搜索引擎：自动补全，拼写检查
// 2. 输入法：词汇联想，智能输入
// 3. 路由系统：URL路径匹配
// 4. 编程语言：关键字识别，语法分析
// 5. 数据库：索引优化，模糊查询
// 6. 网络安全：恶意URL检测，敏感词过滤
// 7. 生物信息学：DNA序列匹配
type Trie struct {
	root *TrieNode // 根节点
	size int       // 存储的单词数量
}

// NewTrie 创建新的字典树
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

// Insert 插入单词
// 时间复杂度：O(m)，其中m是单词长度
func (t *Trie) Insert(word string) {
	t.InsertWithValue(word, nil)
}

// InsertWithValue 插入单词并关联值
func (t *Trie) InsertWithValue(word string, value interface{}) {
	if word == "" {
		return
	}

	current := t.root

	// 遍历单词的每个字符
	for _, char := range word {
		// 如果子节点不存在，创建新节点
		if current.children[char] == nil {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
		current.count++ // 增加经过此节点的单词计数
	}

	// 标记单词结尾
	if !current.isEnd {
		current.isEnd = true
		t.size++
	}
	current.value = value
}

// Search 搜索单词是否存在
// 时间复杂度：O(m)，其中m是单词长度
func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEnd
}

// SearchWithValue 搜索单词并返回关联的值
func (t *Trie) SearchWithValue(word string) (interface{}, bool) {
	node := t.searchNode(word)
	if node != nil && node.isEnd {
		return node.value, true
	}
	return nil, false
}

// searchNode 搜索到指定单词的节点
func (t *Trie) searchNode(word string) *TrieNode {
	if word == "" {
		return t.root
	}

	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			return nil
		}
		current = current.children[char]
	}

	return current
}

// StartsWith 检查是否存在以指定前缀开头的单词
// 时间复杂度：O(m)，其中m是前缀长度
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

// GetWordsWithPrefix 获取所有以指定前缀开头的单词
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	node := t.searchNode(prefix)
	if node == nil {
		return []string{}
	}

	result := make([]string, 0)
	t.collectWords(node, prefix, &result)
	return result
}

// collectWords 递归收集以当前节点为根的所有单词
func (t *Trie) collectWords(node *TrieNode, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}

	for char, child := range node.children {
		t.collectWords(child, prefix+string(char), result)
	}
}

// GetWordsWithPrefixAndValue 获取所有以指定前缀开头的单词及其值
func (t *Trie) GetWordsWithPrefixAndValue(prefix string) []WordValue {
	node := t.searchNode(prefix)
	if node == nil {
		return []WordValue{}
	}

	result := make([]WordValue, 0)
	t.collectWordsWithValue(node, prefix, &result)
	return result
}

// WordValue 单词和值的组合
type WordValue struct {
	Word  string
	Value interface{}
}

// collectWordsWithValue 递归收集以当前节点为根的所有单词和值
func (t *Trie) collectWordsWithValue(node *TrieNode, prefix string, result *[]WordValue) {
	if node.isEnd {
		*result = append(*result, WordValue{
			Word:  prefix,
			Value: node.value,
		})
	}

	for char, child := range node.children {
		t.collectWordsWithValue(child, prefix+string(char), result)
	}
}

// Delete 删除单词
// 时间复杂度：O(m)，其中m是单词长度
func (t *Trie) Delete(word string) bool {
	if word == "" || !t.Search(word) {
		return false
	}

	// deleteHelper 的返回值表示"父节点是否需要剪掉该子节点"（剪枝语义），
	// 并非"删除是否成功"；成功与否已由上面的 Search 判定
	t.deleteHelper(t.root, word, 0)
	t.size--
	return true
}

// deleteHelper 递归删除单词的辅助函数
func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len([]rune(word)) {
		// 到达单词末尾
		if !node.isEnd {
			return false // 单词不存在
		}
		node.isEnd = false
		node.value = nil
		// 如果节点没有子节点，可以删除
		return len(node.children) == 0
	}

	runes := []rune(word)
	char := runes[index]
	child := node.children[char]

	if child == nil {
		return false // 单词不存在
	}

	// 递归删除
	shouldDeleteChild := t.deleteHelper(child, word, index+1)

	if shouldDeleteChild {
		// 删除子节点
		delete(node.children, char)
		// 如果当前节点不是单词结尾且没有其他子节点，也可以删除
		return !node.isEnd && len(node.children) == 0
	}

	// 更新计数
	child.count--
	return false
}

// Size 返回字典树中单词的数量
func (t *Trie) Size() int {
	return t.size
}

// IsEmpty 检查字典树是否为空
func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

// Clear 清空字典树
func (t *Trie) Clear() {
	t.root = NewTrieNode()
	t.size = 0
}

// GetAllWords 获取所有单词
func (t *Trie) GetAllWords() []string {
	return t.GetWordsWithPrefix("")
}

// GetAllWordsWithValue 获取所有单词及其值
func (t *Trie) GetAllWordsWithValue() []WordValue {
	return t.GetWordsWithPrefixAndValue("")
}

// GetLongestCommonPrefix 获取所有单词的最长公共前缀
func (t *Trie) GetLongestCommonPrefix() string {
	if t.IsEmpty() {
		return ""
	}

	var prefix strings.Builder
	current := t.root

	// 当只有一个子节点且不是单词结尾时，继续
	for len(current.children) == 1 && !current.isEnd {
		for char, child := range current.children {
			prefix.WriteRune(char)
			current = child
			break
		}
	}

	return prefix.String()
}

// CountWordsWithPrefix 统计以指定前缀开头的单词数量
func (t *Trie) CountWordsWithPrefix(prefix string) int {
	node := t.searchNode(prefix)
	if node == nil {
		return 0
	}
	return t.countWords(node)
}

// countWords 递归统计以当前节点为根的单词数量
func (t *Trie) countWords(node *TrieNode) int {
	count := 0
	if node.isEnd {
		count = 1
	}

	for _, child := range node.children {
		count += t.countWords(child)
	}

	return count
}

// GetShortestUniquePrefix 获取单词的最短唯一前缀
func (t *Trie) GetShortestUniquePrefix(word string) string {
	if !t.Search(word) {
		return ""
	}

	current := t.root
	var prefix strings.Builder

	for _, char := range word {
		prefix.WriteRune(char)
		current = current.children[char]

		// 如果当前节点的计数为1，说明这是唯一路径
		if current.count == 1 {
			return prefix.String()
		}
	}

	return word // 如果没有找到唯一前缀，返回完整单词
}

// AutoComplete 自动补全功能
func (t *Trie) AutoComplete(prefix string, maxResults int) []string {
	words := t.GetWordsWithPrefix(prefix)
	if maxResults > 0 && len(words) > maxResults {
		return words[:maxResults]
	}
	return words
}

// FuzzySearch 模糊搜索（允许一个字符的差异）
func (t *Trie) FuzzySearch(word string, maxDistance int) []string {
	result := make([]string, 0)
	t.fuzzySearchHelper(t.root, "", word, 0, maxDistance, &result)
	return result
}

// fuzzySearchHelper 模糊搜索的递归辅助函数
func (t *Trie) fuzzySearchHelper(node *TrieNode, current, target string, distance, maxDistance int, result *[]string) {
	if distance > maxDistance {
		return
	}

	if node.isEnd && len(current) >= len(target)-maxDistance && len(current) <= len(target)+maxDistance {
		*result = append(*result, current)
	}

	targetRunes := []rune(target)
	currentRunes := []rune(current)

	for char, child := range node.children {
		newCurrent := current + string(char)
		newDistance := distance

		// 如果当前字符与目标字符不匹配，增加距离
		if len(currentRunes) < len(targetRunes) && char != targetRunes[len(currentRunes)] {
			newDistance++
		}

		t.fuzzySearchHelper(child, newCurrent, target, newDistance, maxDistance, result)
	}
}

// String 字符串表示
func (t *Trie) String() string {
	if t.IsEmpty() {
		return "Trie{}"
	}

	words := t.GetAllWords()
	return "Trie{words: [" + strings.Join(words, ", ") + "]}"
}

// PrintTrie 打印字典树结构
func (t *Trie) PrintTrie() {
	t.printTrieHelper(t.root, "", "")
}

// printTrieHelper 递归打印字典树的辅助函数
func (t *Trie) printTrieHelper(node *TrieNode, prefix, indent string) {
	if node.isEnd {
		println(indent + prefix + " (END)")
	} else if prefix != "" {
		println(indent + prefix)
	}

	for char, child := range node.children {
		t.printTrieHelper(child, string(char), indent+"  ")
	}
}

// println 简单的打印函数
func println(s string) {
	// 这里可以使用fmt.Println，但为了避免导入fmt包，使用简单实现
	_ = s
}

// CompressedTrie 压缩字典树（Patricia Trie）
// 将只有一个子节点的节点链压缩为单个节点
type CompressedTrie struct {
	root *CompressedTrieNode
	size int
}

// CompressedTrieNode 压缩字典树节点
type CompressedTrieNode struct {
	children map[rune]*CompressedTrieNode
	label    string      // 节点标签（可能包含多个字符）
	isEnd    bool        // 是否为单词结尾
	value    interface{} // 存储的值
}

// NewCompressedTrie 创建压缩字典树
func NewCompressedTrie() *CompressedTrie {
	return &CompressedTrie{
		root: &CompressedTrieNode{
			children: make(map[rune]*CompressedTrieNode),
			label:    "",
			isEnd:    false,
		},
		size: 0,
	}
}

// Insert 在压缩字典树中插入单词
func (ct *CompressedTrie) Insert(word string) {
	if word == "" {
		return
	}

	current := ct.root
	i := 0
	runes := []rune(word)

	for i < len(runes) {
		char := runes[i]

		if child, exists := current.children[char]; exists {
			// 找到匹配的子节点
			labelRunes := []rune(child.label)
			j := 0

			// 比较标签和剩余单词
			for j < len(labelRunes) && i+j < len(runes) && labelRunes[j] == runes[i+j] {
				j++
			}

			if j == len(labelRunes) {
				// 完全匹配标签
				current = child
				i += j
			} else {
				// 部分匹配，需要分割节点
				ct.splitNode(child, j)
				current = child
				i += j
			}
		} else {
			// 创建新的子节点
			newNode := &CompressedTrieNode{
				children: make(map[rune]*CompressedTrieNode),
				label:    string(runes[i:]),
				isEnd:    true,
			}
			current.children[char] = newNode
			ct.size++
			return
		}
	}

	// 标记单词结尾
	if !current.isEnd {
		current.isEnd = true
		ct.size++
	}
}

// splitNode 分割压缩字典树节点
func (ct *CompressedTrie) splitNode(node *CompressedTrieNode, splitIndex int) {
	labelRunes := []rune(node.label)

	// 创建新的子节点
	newChild := &CompressedTrieNode{
		children: node.children,
		label:    string(labelRunes[splitIndex:]),
		isEnd:    node.isEnd,
		value:    node.value,
	}

	// 更新当前节点
	node.children = make(map[rune]*CompressedTrieNode)
	node.children[labelRunes[splitIndex]] = newChild
	node.label = string(labelRunes[:splitIndex])
	node.isEnd = false
	node.value = nil
}

// Search 在压缩字典树中搜索单词
func (ct *CompressedTrie) Search(word string) bool {
	node := ct.searchNode(word)
	return node != nil && node.isEnd
}

// searchNode 在压缩字典树中搜索节点
func (ct *CompressedTrie) searchNode(word string) *CompressedTrieNode {
	if word == "" {
		return ct.root
	}

	current := ct.root
	runes := []rune(word)
	i := 0

	for i < len(runes) {
		char := runes[i]

		if child, exists := current.children[char]; exists {
			labelRunes := []rune(child.label)

			// 检查标签是否匹配
			if i+len(labelRunes) > len(runes) {
				return nil // 剩余字符不足
			}

			for j, labelChar := range labelRunes {
				if runes[i+j] != labelChar {
					return nil // 不匹配
				}
			}

			current = child
			i += len(labelRunes)
		} else {
			return nil // 没有匹配的子节点
		}
	}

	return current
}

// Size 返回压缩字典树中单词的数量
func (ct *CompressedTrie) Size() int {
	return ct.size
}

// 业务应用示例：
// 1. 搜索引擎：关键词自动补全，查询建议
// 2. 输入法：拼音输入，词汇联想
// 3. 编程IDE：代码自动补全，语法高亮
// 4. 路由系统：URL路径匹配，RESTful API
// 5. 数据库：前缀索引，模糊查询优化
// 6. 网络安全：恶意域名检测，敏感词过滤
// 7. 生物信息学：基因序列分析，蛋白质匹配
// 8. 自然语言处理：词典构建，文本分析
