package corealgo

// --- 字符串算法 ---

// KMPSearch KMP 字符串匹配：返回 pattern 在 text 中所有出现位置的下标
// 思路：先对模式串构建 next（前缀函数）数组，next[i] 表示 pattern[:i+1]
// 的最长相等前后缀长度。匹配失败时利用 next 数组回退模式串指针，主串指针
// 不回退，从而将朴素匹配的 O(n*m) 降为线性
// 时间复杂度：O(n+m)，n、m 分别为主串与模式串长度
// 空间复杂度：O(m)
// 适用场景：
// 1. 文本编辑器中的查找与替换
// 2. 日志、基因序列中的模式匹配
// 3. 敏感词过滤与内容审核
// 注意：pattern 为空时返回空切片
func KMPSearch(text, pattern string) []int {
	result := make([]int, 0)
	m := len(pattern)
	if m == 0 {
		return result
	}
	next := buildNext(pattern)
	n := len(text)
	i, j := 0, 0
	for i < n {
		if text[i] == pattern[j] {
			i++
			j++
			if j == m {
				result = append(result, i-j)
				j = next[j-1]
			}
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
	}
	return result
}

// buildNext 构建模式串的 next（前缀函数）数组
// next[i] 表示 pattern[0:i+1] 的最长相等前后缀长度（不含整个子串自身）
func buildNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

// RabinKarpSearch Rabin-Karp 滚动哈希字符串匹配：返回 pattern 在 text 中
// 所有出现位置的下标
// 思路：将字符串视为 base 进制数并取模得到哈希值。先计算模式串哈希与主串
// 第一个窗口的哈希，之后用滚动公式 O(1) 更新窗口哈希；哈希相等时再做一次
// 逐字符比对，消除哈希冲突带来的误匹配
// 时间复杂度：平均 O(n+m)，最坏 O(n*m)（哈希冲突严重时）
// 空间复杂度：O(1)
// 适用场景：
// 1. 多模式串批量匹配（多个敏感词同时扫描）
// 2. 大文本中的子串查找与指纹比对
// 3. 抄袭检测、文档查重
// 注意：pattern 为空或长于 text 时返回空切片
func RabinKarpSearch(text, pattern string) []int {
	const (
		base int64 = 131
		mod  int64 = 1000000007
	)
	result := make([]int, 0)
	n, m := len(text), len(pattern)
	if m == 0 || n < m {
		return result
	}
	// 计算模式串哈希与主串首窗口哈希
	var pHash, wHash int64
	for i := 0; i < m; i++ {
		pHash = (pHash*base + int64(pattern[i])) % mod
		wHash = (wHash*base + int64(text[i])) % mod
	}
	// pow = base^(m-1) % mod
	pow := int64(1)
	for i := 1; i < m; i++ {
		pow = pow * base % mod
	}
	for i := 0; i+m <= n; i++ {
		// 哈希相等时二次验证，避免哈希冲突误判
		if wHash == pHash && text[i:i+m] == pattern {
			result = append(result, i)
		}
		if i+m < n {
			// 滚动窗口：减去最左字符贡献，加入新字符
			wHash = ((wHash-int64(text[i])*pow%mod+mod)%mod*base + int64(text[i+m])) % mod
		}
	}
	return result
}
