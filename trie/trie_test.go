package trie

import (
	"reflect"
	"sort"
	"testing"
)

// TestNewTrie 测试创建新的字典树
func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	
	if trie == nil {
		t.Error("NewTrie() returned nil")
	}
	
	if trie.Size() != 0 {
		t.Errorf("Expected size 0, got %d", trie.Size())
	}
	
	if !trie.IsEmpty() {
		t.Error("Expected empty trie")
	}
}

// TestInsertAndSearch 测试插入和搜索
func TestInsertAndSearch(t *testing.T) {
	trie := NewTrie()
	
	// 测试插入单个单词
	trie.Insert("hello")
	if !trie.Search("hello") {
		t.Error("Expected to find 'hello'")
	}
	
	if trie.Search("hell") {
		t.Error("Should not find 'hell'")
	}
	
	if trie.Search("hello world") {
		t.Error("Should not find 'hello world'")
	}
	
	// 测试插入多个单词
	words := []string{"world", "hell", "help", "hero"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find '%s'", word)
		}
	}
	
	if trie.Size() != 5 {
		t.Errorf("Expected size 5, got %d", trie.Size())
	}
}

// TestInsertWithValue 测试插入带值的单词
func TestInsertWithValue(t *testing.T) {
	trie := NewTrie()
	
	trie.InsertWithValue("apple", 100)
	trie.InsertWithValue("app", 50)
	trie.InsertWithValue("application", 200)
	
	// 测试搜索带值
	value, found := trie.SearchWithValue("apple")
	if !found || value != 100 {
		t.Errorf("Expected value 100 for 'apple', got %v", value)
	}
	
	value, found = trie.SearchWithValue("app")
	if !found || value != 50 {
		t.Errorf("Expected value 50 for 'app', got %v", value)
	}
	
	_, found = trie.SearchWithValue("appl")
	if found {
		t.Error("Should not find value for 'appl'")
	}
}

// TestStartsWith 测试前缀检查
func TestStartsWith(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "app", "application", "apply"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	if !trie.StartsWith("app") {
		t.Error("Expected to find prefix 'app'")
	}
	
	if !trie.StartsWith("appl") {
		t.Error("Expected to find prefix 'appl'")
	}
	
	if trie.StartsWith("banana") {
		t.Error("Should not find prefix 'banana'")
	}
	
	if !trie.StartsWith("") {
		t.Error("Empty prefix should always be found")
	}
}

// TestGetWordsWithPrefix 测试获取前缀单词
func TestGetWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "app", "application", "apply", "banana", "band"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	// 测试前缀 "app"
	appWords := trie.GetWordsWithPrefix("app")
	sort.Strings(appWords)
	expected := []string{"app", "apple", "application", "apply"}
	sort.Strings(expected)
	
	if !reflect.DeepEqual(appWords, expected) {
		t.Errorf("Expected %v, got %v", expected, appWords)
	}
	
	// 测试前缀 "ban"
	banWords := trie.GetWordsWithPrefix("ban")
	sort.Strings(banWords)
	expected = []string{"banana", "band"}
	sort.Strings(expected)
	
	if !reflect.DeepEqual(banWords, expected) {
		t.Errorf("Expected %v, got %v", expected, banWords)
	}
	
	// 测试不存在的前缀
	noWords := trie.GetWordsWithPrefix("xyz")
	if len(noWords) != 0 {
		t.Errorf("Expected empty slice, got %v", noWords)
	}
}

// TestGetWordsWithPrefixAndValue 测试获取前缀单词和值
func TestGetWordsWithPrefixAndValue(t *testing.T) {
	trie := NewTrie()
	
	trie.InsertWithValue("apple", 100)
	trie.InsertWithValue("app", 50)
	trie.InsertWithValue("application", 200)
	
	results := trie.GetWordsWithPrefixAndValue("app")
	
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}
	
	// 验证结果
	wordValues := make(map[string]interface{})
	for _, wv := range results {
		wordValues[wv.Word] = wv.Value
	}
	
	if wordValues["apple"] != 100 {
		t.Errorf("Expected value 100 for 'apple', got %v", wordValues["apple"])
	}
	
	if wordValues["app"] != 50 {
		t.Errorf("Expected value 50 for 'app', got %v", wordValues["app"])
	}
	
	if wordValues["application"] != 200 {
		t.Errorf("Expected value 200 for 'application', got %v", wordValues["application"])
	}
}

// TestDelete 测试删除单词
func TestDelete(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "app", "application", "apply"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	initialSize := trie.Size()
	
	// 删除存在的单词
	if !trie.Delete("apple") {
		t.Error("Expected to delete 'apple'")
	}
	
	if trie.Search("apple") {
		t.Error("Should not find 'apple' after deletion")
	}
	
	if trie.Size() != initialSize-1 {
		t.Errorf("Expected size %d, got %d", initialSize-1, trie.Size())
	}
	
	// 确保其他单词仍然存在
	if !trie.Search("app") {
		t.Error("Should still find 'app'")
	}
	
	if !trie.Search("application") {
		t.Error("Should still find 'application'")
	}
	
	// 删除不存在的单词
	if trie.Delete("banana") {
		t.Error("Should not delete non-existent word")
	}
	
	// 删除空字符串
	if trie.Delete("") {
		t.Error("Should not delete empty string")
	}
}

// TestClear 测试清空字典树
func TestClear(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "banana", "cherry"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	if trie.IsEmpty() {
		t.Error("Trie should not be empty before clear")
	}
	
	trie.Clear()
	
	if !trie.IsEmpty() {
		t.Error("Trie should be empty after clear")
	}
	
	if trie.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", trie.Size())
	}
	
	for _, word := range words {
		if trie.Search(word) {
			t.Errorf("Should not find '%s' after clear", word)
		}
	}
}

// TestGetAllWords 测试获取所有单词
func TestGetAllWords(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "banana", "cherry", "app"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	allWords := trie.GetAllWords()
	sort.Strings(allWords)
	sort.Strings(words)
	
	if !reflect.DeepEqual(allWords, words) {
		t.Errorf("Expected %v, got %v", words, allWords)
	}
}

// TestGetLongestCommonPrefix 测试最长公共前缀
func TestGetLongestCommonPrefix(t *testing.T) {
	trie := NewTrie()
	
	// 测试空字典树
	if prefix := trie.GetLongestCommonPrefix(); prefix != "" {
		t.Errorf("Expected empty prefix for empty trie, got '%s'", prefix)
	}
	
	// 测试有公共前缀的情况
	words := []string{"application", "apply", "apple"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	prefix := trie.GetLongestCommonPrefix()
	expected := "app"
	if prefix != expected {
		t.Errorf("Expected prefix '%s', got '%s'", expected, prefix)
	}
	
	// 添加不同前缀的单词
	trie.Insert("banana")
	prefix = trie.GetLongestCommonPrefix()
	if prefix != "" {
		t.Errorf("Expected empty prefix after adding different word, got '%s'", prefix)
	}
}

// TestCountWordsWithPrefix 测试前缀单词计数
func TestCountWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "app", "application", "apply", "banana"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	// 测试前缀 "app"
	count := trie.CountWordsWithPrefix("app")
	if count != 4 {
		t.Errorf("Expected count 4 for prefix 'app', got %d", count)
	}
	
	// 测试前缀 "ban"
	count = trie.CountWordsWithPrefix("ban")
	if count != 1 {
		t.Errorf("Expected count 1 for prefix 'ban', got %d", count)
	}
	
	// 测试不存在的前缀
	count = trie.CountWordsWithPrefix("xyz")
	if count != 0 {
		t.Errorf("Expected count 0 for prefix 'xyz', got %d", count)
	}
}

// TestGetShortestUniquePrefix 测试最短唯一前缀
func TestGetShortestUniquePrefix(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "application", "banana", "band"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	// 测试 "apple" 的最短唯一前缀
	prefix := trie.GetShortestUniquePrefix("apple")
	expected := "appl"
	if prefix != expected {
		t.Errorf("Expected prefix '%s' for 'apple', got '%s'", expected, prefix)
	}
	
	// 测试 "banana" 的最短唯一前缀
	prefix = trie.GetShortestUniquePrefix("banana")
	expected = "ban"
	if prefix != expected {
		t.Errorf("Expected prefix '%s' for 'banana', got '%s'", expected, prefix)
	}
	
	// 测试不存在的单词
	prefix = trie.GetShortestUniquePrefix("xyz")
	if prefix != "" {
		t.Errorf("Expected empty prefix for non-existent word, got '%s'", prefix)
	}
}

// TestAutoComplete 测试自动补全
func TestAutoComplete(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "app", "application", "apply", "appreciate"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	// 测试无限制的自动补全
	results := trie.AutoComplete("app", 0)
	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}
	
	// 测试限制结果数量
	results = trie.AutoComplete("app", 3)
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}
	
	// 测试不存在的前缀
	results = trie.AutoComplete("xyz", 5)
	if len(results) != 0 {
		t.Errorf("Expected 0 results for non-existent prefix, got %d", len(results))
	}
}

// TestFuzzySearch 测试模糊搜索
func TestFuzzySearch(t *testing.T) {
	trie := NewTrie()
	
	words := []string{"apple", "apply", "happy", "help"}
	for _, word := range words {
		trie.Insert(word)
	}
	
	// 测试模糊搜索 "appl" (允许1个字符差异)
	results := trie.FuzzySearch("appl", 1)
	
	// 应该找到 "apple" 和 "apply"
	found := make(map[string]bool)
	for _, word := range results {
		found[word] = true
	}
	
	if !found["apple"] {
		t.Error("Expected to find 'apple' in fuzzy search")
	}
	
	if !found["apply"] {
		t.Error("Expected to find 'apply' in fuzzy search")
	}
}

// TestString 测试字符串表示
func TestString(t *testing.T) {
	trie := NewTrie()
	
	// 测试空字典树
	str := trie.String()
	if str != "Trie{}" {
		t.Errorf("Expected 'Trie{}' for empty trie, got '%s'", str)
	}
	
	// 测试非空字典树
	trie.Insert("apple")
	trie.Insert("app")
	
	str = trie.String()
	if str == "Trie{}" {
		t.Error("String representation should not be empty for non-empty trie")
	}
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	trie := NewTrie()
	
	// 测试空字符串
	trie.Insert("")
	if trie.Size() != 0 {
		t.Error("Should not insert empty string")
	}
	
	// 测试单字符单词
	trie.Insert("a")
	if !trie.Search("a") {
		t.Error("Should find single character word")
	}
	
	// 测试重复插入
	trie.Insert("apple")
	trie.Insert("apple")
	if trie.Size() != 2 {
		t.Errorf("Expected size 2 after duplicate insert, got %d", trie.Size())
	}
	
	// 测试Unicode字符
	trie.Insert("你好")
	if !trie.Search("你好") {
		t.Error("Should support Unicode characters")
	}
}

// TestCompressedTrie 测试压缩字典树
func TestCompressedTrie(t *testing.T) {
	ct := NewCompressedTrie()
	
	if ct == nil {
		t.Error("NewCompressedTrie() returned nil")
	}
	
	if ct.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ct.Size())
	}
	
	// 测试插入和搜索
	ct.Insert("apple")
	ct.Insert("app")
	ct.Insert("application")
	
	if !ct.Search("apple") {
		t.Error("Expected to find 'apple'")
	}
	
	if !ct.Search("app") {
		t.Error("Expected to find 'app'")
	}
	
	if !ct.Search("application") {
		t.Error("Expected to find 'application'")
	}
	
	if ct.Search("appl") {
		t.Error("Should not find 'appl'")
	}
	
	if ct.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ct.Size())
	}
}

// Benchmark tests

// BenchmarkTrieInsert 基准测试插入操作
func BenchmarkTrieInsert(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Insert(words[i%len(words)])
	}
}

// BenchmarkTrieSearch 基准测试搜索操作
func BenchmarkTrieSearch(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Search(words[i%len(words)])
	}
}

// BenchmarkTrieGetWordsWithPrefix 基准测试前缀搜索
func BenchmarkTrieGetWordsWithPrefix(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	prefixes := []string{"a", "ap", "app", "test"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.GetWordsWithPrefix(prefixes[i%len(prefixes)])
	}
}

// BenchmarkTrieAutoComplete 基准测试自动补全
func BenchmarkTrieAutoComplete(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	prefixes := []string{"a", "ap", "app", "test"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.AutoComplete(prefixes[i%len(prefixes)], 10)
	}
}

// BenchmarkCompressedTrieInsert 基准测试压缩字典树插入
func BenchmarkCompressedTrieInsert(b *testing.B) {
	ct := NewCompressedTrie()
	words := generateWords(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ct.Insert(words[i%len(words)])
	}
}

// BenchmarkCompressedTrieSearch 基准测试压缩字典树搜索
func BenchmarkCompressedTrieSearch(b *testing.B) {
	ct := NewCompressedTrie()
	words := generateWords(1000)
	
	for _, word := range words {
		ct.Insert(word)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ct.Search(words[i%len(words)])
	}
}

// generateWords 生成测试用的单词列表
func generateWords(count int) []string {
	words := make([]string, count)
	prefixes := []string{"app", "test", "hello", "world", "go", "lang", "data", "struct"}
	suffixes := []string{"le", "ing", "ed", "er", "ly", "tion", "ness", "ment"}
	
	for i := 0; i < count; i++ {
		prefix := prefixes[i%len(prefixes)]
		suffix := suffixes[i%len(suffixes)]
		words[i] = prefix + suffix
	}
	
	return words
}