package advanced

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCounter(t *testing.T) {
	c := NewCounter()
	const goroutines = 50
	const increments = 100

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()

	expected := goroutines * increments
	if c.Value() != expected {
		t.Errorf("Expected %d, got %d", expected, c.Value())
	}
}

func TestRWCache(t *testing.T) {
	cache := NewRWCache()
	const count = 100

	var wg sync.WaitGroup
	// 并发写入 100 个键值对
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			k := fmt.Sprintf("key-%d", id)
			cache.Set(k, fmt.Sprintf("value-%d", id))
		}(i)
	}
	wg.Wait()

	if cache.Size() != count {
		t.Errorf("Expected size %d, got %d", count, cache.Size())
	}

	// 并发读取验证（读锁可并发）
	var miss int64
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			k := fmt.Sprintf("key-%d", id)
			v, ok := cache.Get(k)
			if !ok || v != fmt.Sprintf("value-%d", id) {
				atomic.AddInt64(&miss, 1)
			}
		}(i)
	}
	wg.Wait()

	if miss != 0 {
		t.Errorf("Expected no cache misses, got %d", miss)
	}

	// 不存在的 key
	if _, ok := cache.Get("not-exist"); ok {
		t.Error("Expected miss for non-existent key")
	}
}

func TestGetInstance(t *testing.T) {
	// 并发获取单例：所有 goroutine 必须拿到同一个实例
	const goroutines = 100
	instances := make([]*Singleton, goroutines)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetInstance()
		}(i)
	}
	wg.Wait()

	for i := 1; i < goroutines; i++ {
		if instances[i] != instances[0] {
			t.Errorf("Instance %d differs from instance 0", i)
		}
	}
	if instances[0].Name() != "singleton-demo" {
		t.Errorf("Expected name 'singleton-demo', got %s", instances[0].Name())
	}
}

func TestLazyInit(t *testing.T) {
	atomic.StoreInt64(&lazyInitCount, 0) // 复位计数，保证测试可重复执行

	l := NewLazyInit()
	const goroutines = 50
	values := make([]int, goroutines)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			values[idx] = l.Value()
		}(i)
	}
	wg.Wait()

	// 所有 goroutine 拿到同一个结果
	for i, v := range values {
		if v != 42 {
			t.Errorf("Value at %d: expected 42, got %d", i, v)
		}
	}
	// sync.Once 保证初始化只执行了一次
	if got := LazyInitCount(); got != 1 {
		t.Errorf("Expected init count 1, got %d", got)
	}
}

func TestObjectPool(t *testing.T) {
	atomic.StoreInt64(&poolNewCount, 0)

	pool := NewObjectPool()

	// 池为空时，首次 Acquire 必然通过 New 创建对象（sync.Pool 的确定性契约行为）
	obj1 := pool.Acquire()
	if obj1 == nil {
		t.Fatal("Expected non-nil object from pool")
	}
	if got := PoolNewCount(); got != 1 {
		t.Errorf("Expected exactly 1 new object after first acquire, got %d", got)
	}

	// 注意：sync.Pool 不保证复用——GC 可随时回收池中对象（对象可能进入 victim 缓存），
	// 因此不能断言"拿到同一个对象"。确定性验证方式是：
	// 大量 Acquire/Release 中新建对象数远小于获取次数，说明对象确实被复用了
	before := PoolNewCount()
	for i := 0; i < 100; i++ {
		obj := pool.Acquire()
		if obj == nil {
			t.Fatalf("Acquire returned nil at iteration %d", i)
		}
		pool.Release(obj)
	}
	after := PoolNewCount()

	if after < before {
		t.Errorf("New count should not decrease: before %d, after %d", before, after)
	}
	if after-before >= 100 {
		t.Errorf("Expected reuse, got %d new objects for 100 acquires", after-before)
	}
}

func TestObjectPoolAfterGC(t *testing.T) {
	pool := NewObjectPool()
	obj := pool.Acquire()
	pool.Release(obj)

	// GC 会清空池（对象进入 victim 缓存或被回收），但 Get 仍然返回非 nil
	runtime.GC()
	if got := pool.Acquire(); got == nil {
		t.Error("Expected non-nil object after GC")
	}
}
