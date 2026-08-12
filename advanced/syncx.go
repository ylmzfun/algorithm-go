package advanced

import (
	"sync"
	"sync/atomic"
)

// sync 包提供的并发原语：
// 1. Mutex：互斥锁，保证临界区同一时刻只有一个 goroutine 进入
// 2. RWMutex：读写锁，读多写少场景下多个读者可并发，写者独占
// 3. Once：保证某个函数只执行一次（单例初始化、懒加载）
// 4. Pool：对象池，复用临时对象以减少 GC 压力（注意：池中对象可能被随时回收）

// Counter 基于 Mutex 的并发安全计数器
// 思路：用互斥锁保护 count 字段，Inc 与 Value 都加锁访问
// 作用：演示 Mutex 保护共享状态的基本模式
// 业务场景：并发请求统计、在线人数计数、流量统计
type Counter struct {
	mu    sync.Mutex
	count int
}

// NewCounter 创建计数器
func NewCounter() *Counter {
	return &Counter{}
}

// Inc 计数器加一
// 时间复杂度：O(1)
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value 返回当前计数值
// 时间复杂度：O(1)
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// RWCache 基于 RWMutex 的并发安全缓存
// 思路：读操作加读锁（RLock，可并发），写操作加写锁（Lock，独占）
// 作用：演示 RWMutex 的读多写少优化
// 业务场景：配置缓存、热点数据缓存、字典表
type RWCache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewRWCache 创建缓存
func NewRWCache() *RWCache {
	return &RWCache{data: make(map[string]string)}
}

// Set 写入键值对（写锁，独占）
// 时间复杂度：O(1)
func (c *RWCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get 读取键对应的值（读锁，多个读者可并发）
// 时间复杂度：O(1)
func (c *RWCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

// Size 返回缓存条目数（读锁）
// 时间复杂度：O(1)
func (c *RWCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

// 单例实例：由 sync.Once 保证只初始化一次
var (
	singletonOnce sync.Once
	singleton     *Singleton
)

// Singleton 演示 sync.Once 实现的单例
// 思路：多个 goroutine 同时调用 GetInstance 时，Once 保证初始化函数只执行一次，
// 其余 goroutine 直接拿到已初始化好的实例
// 作用：演示单例模式的并发安全实现
// 业务场景：全局配置对象、数据库连接池、日志器
type Singleton struct {
	name string
}

// GetInstance 获取单例实例
func GetInstance() *Singleton {
	singletonOnce.Do(func() {
		singleton = &Singleton{name: "singleton-demo"}
	})
	return singleton
}

// Name 返回单例名称
func (s *Singleton) Name() string {
	return s.name
}

// lazyInitCount 初始化函数实际执行次数（供测试与演示观察 once 语义）
var lazyInitCount int64

// LazyInit 演示 sync.Once 的惰性初始化（懒加载）
// 思路：once.Do 保证初始化函数在多个 goroutine 并发调用时也只执行一次
// 作用：演示"首次访问时初始化"的延迟计算模式
// 业务场景：昂贵资源的懒加载（大对象、网络连接、随机数种子）
type LazyInit struct {
	once  sync.Once
	value int
}

// NewLazyInit 创建惰性初始化对象
func NewLazyInit() *LazyInit {
	return &LazyInit{}
}

// Value 返回初始化后的值，初始化只执行一次
// 时间复杂度：首次调用 O(初始化开销)，之后 O(1)
func (l *LazyInit) Value() int {
	l.once.Do(func() {
		atomic.AddInt64(&lazyInitCount, 1)
		l.value = 42 // 模拟昂贵初始化产出的值（真实场景可能是加载配置、建立连接等）
	})
	return l.value
}

// LazyInitCount 返回初始化函数实际执行次数
func LazyInitCount() int64 {
	return atomic.LoadInt64(&lazyInitCount)
}

// poolNewCount 池中实际新建对象的次数（供测试与演示观察复用效果）
var poolNewCount int64

// Buffer 对象池演示用的可复用缓冲区对象
type Buffer struct {
	data []byte
}

// NewBuffer 创建缓冲区（预留 capacity 容量以便复用）
func NewBuffer(capacity int) *Buffer {
	return &Buffer{data: make([]byte, 0, capacity)}
}

// ObjectPool 基于 sync.Pool 的对象池
// 思路：Acquire 从池中取对象（池空时调用 New 创建），Release 将对象归还
// 注意：sync.Pool 中的对象可能被 GC 随时回收，因此只适合存放"可重建"的临时对象
// 作用：演示对象复用，减少频繁分配带来的 GC 压力
// 业务场景：网络缓冲区、JSON 编解码中间对象等频繁创建销毁的临时对象
type ObjectPool struct {
	pool sync.Pool
}

// NewObjectPool 创建对象池
func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: sync.Pool{
			New: func() interface{} {
				atomic.AddInt64(&poolNewCount, 1)
				return NewBuffer(1024) // 模拟一个可复用的缓冲区对象
			},
		},
	}
}

// Acquire 从池中获取一个对象；池为空时调用 New 创建
func (p *ObjectPool) Acquire() interface{} {
	return p.pool.Get()
}

// Release 将对象归还池中以便复用
func (p *ObjectPool) Release(obj interface{}) {
	p.pool.Put(obj)
}

// PoolNewCount 返回池中新建对象的累计次数
func PoolNewCount() int64 {
	return atomic.LoadInt64(&poolNewCount)
}
