package bloomfilter

import (
	"fmt"
	"hash/fnv"
	"math"
)

// BloomFilter 布隆过滤器实现
// 思路：使用多个哈希函数将元素映射到比特数组的不同位置。查询时，若所有映射位都为 1，则认为元素"可能存在"；
//
//	若任一位为 0，则元素"一定不存在"。是一种空间高效的概率性数据结构。
//
// 作用：在允许少量误判的情况下，用极少的空间快速判断元素是否存在
// 业务场景：
// 1. 缓存穿透防护：快速判断请求的 key 是否在数据库中
// 2. URL 去重：爬虫系统中判断 URL 是否已爬取
// 3. 垃圾邮件过滤：判断邮件地址是否在黑名单中
// 4. 数据库去重：BigTable、HBase、Cassandra 中判断 key 是否存在
// 5. 推荐系统：判断用户是否已看过某个内容
// 6. 网络路由器：包过滤和速率限制
// 7. 区块链：SPV 节点判断交易是否在区块中
type BloomFilter struct {
	bitSet    []bool  // 比特数组
	size      uint64  // 比特数组大小 (m)
	hashCount int     // 哈希函数个数 (k)
	count     uint64  // 已插入元素数量
}

// NewBloomFilter 创建布隆过滤器
// expectedElements: 预期元素数量
// falsePositiveRate: 可接受的误判率 (0 < p < 1)
func NewBloomFilter(expectedElements uint64, falsePositiveRate float64) *BloomFilter {
	if expectedElements == 0 {
		expectedElements = 1000
	}
	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		falsePositiveRate = 0.01
	}

	// 计算最优的比特数组大小 m 和哈希函数个数 k
	// m = -n * ln(p) / (ln(2))^2
	// k = (m / n) * ln(2)
	size := uint64(math.Ceil(-float64(expectedElements) * math.Log(falsePositiveRate) / (math.Pow(math.Ln2, 2))))
	hashCount := int(math.Ceil(float64(size) / float64(expectedElements) * math.Ln2))

	if hashCount < 1 {
		hashCount = 1
	}

	return &BloomFilter{
		bitSet:    make([]bool, size),
		size:      size,
		hashCount: hashCount,
		count:     0,
	}
}

// NewBloomFilterWithSize 使用指定大小和哈希函数个数创建布隆过滤器
func NewBloomFilterWithSize(size uint64, hashCount int) *BloomFilter {
	if size == 0 {
		size = 10000
	}
	if hashCount < 1 {
		hashCount = 3
	}
	return &BloomFilter{
		bitSet:    make([]bool, size),
		size:      size,
		hashCount: hashCount,
		count:     0,
	}
}

// hash 计算第 i 个哈希值
// 使用双重哈希技术：h(i, key) = (hash1(key) + i * hash2(key)) % size
func (bf *BloomFilter) hash(data []byte, seed uint32) uint64 {
	// 使用 FNV 哈希 + 不同的种子来模拟多个哈希函数
	h := fnv.New64a()
	h.Write(data)
	// 将种子混入
	seedBytes := make([]byte, 4)
	seedBytes[0] = byte(seed)
	seedBytes[1] = byte(seed >> 8)
	seedBytes[2] = byte(seed >> 16)
	seedBytes[3] = byte(seed >> 24)
	h.Write(seedBytes)
	return h.Sum64() % bf.size
}

// Add 添加元素
// 时间复杂度：O(k)，k 为哈希函数个数
func (bf *BloomFilter) Add(data []byte) {
	for i := 0; i < bf.hashCount; i++ {
		index := bf.hash(data, uint32(i))
		bf.bitSet[index] = true
	}
	bf.count++
}

// AddString 添加字符串元素
func (bf *BloomFilter) AddString(s string) {
	bf.Add([]byte(s))
}

// Contains 检查元素是否可能存在
// 返回 true 表示元素"可能存在"（可能误判），false 表示元素"一定不存在"
// 时间复杂度：O(k)
func (bf *BloomFilter) Contains(data []byte) bool {
	for i := 0; i < bf.hashCount; i++ {
		index := bf.hash(data, uint32(i))
		if !bf.bitSet[index] {
			return false
		}
	}
	return true
}

// ContainsString 检查字符串元素是否可能存在
func (bf *BloomFilter) ContainsString(s string) bool {
	return bf.Contains([]byte(s))
}

// Count 返回已插入元素数量（近似）
func (bf *BloomFilter) Count() uint64 {
	return bf.count
}

// Size 返回比特数组大小
func (bf *BloomFilter) Size() uint64 {
	return bf.size
}

// HashCount 返回哈希函数个数
func (bf *BloomFilter) HashCount() int {
	return bf.hashCount
}

// EstimatedFalsePositiveRate 估算当前误判率
// p ≈ (1 - e^(-kn/m))^k
func (bf *BloomFilter) EstimatedFalsePositiveRate() float64 {
	k := float64(bf.hashCount)
	n := float64(bf.count)
	m := float64(bf.size)
	return math.Pow(1-math.Exp(-k*n/m), k)
}

// FillRatio 返回比特数组填充比例
func (bf *BloomFilter) FillRatio() float64 {
	if bf.size == 0 {
		return 0
	}
	ones := uint64(0)
	for _, bit := range bf.bitSet {
		if bit {
			ones++
		}
	}
	return float64(ones) / float64(bf.size)
}

// Clear 清空布隆过滤器
func (bf *BloomFilter) Clear() {
	for i := range bf.bitSet {
		bf.bitSet[i] = false
	}
	bf.count = 0
}

// Merge 合并另一个布隆过滤器（两个过滤器必须具有相同的 size 和 hashCount）
func (bf *BloomFilter) Merge(other *BloomFilter) error {
	if bf.size != other.size || bf.hashCount != other.hashCount {
		return fmt.Errorf("bloom filters must have same size and hash count")
	}
	for i := range bf.bitSet {
		bf.bitSet[i] = bf.bitSet[i] || other.bitSet[i]
	}
	bf.count += other.count
	return nil
}

// String 字符串表示
func (bf *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter{size: %d, hashCount: %d, elements: %d, fillRatio: %.2f%%, estimatedFPR: %.4f%%}",
		bf.size, bf.hashCount, bf.count, bf.FillRatio()*100, bf.EstimatedFalsePositiveRate()*100)
}
