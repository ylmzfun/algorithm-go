package advanced

import (
	"testing"
)

func TestProduceConsumeAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	// 无缓冲 channel：生产者/消费者通过 channel 同步
	got := ConsumeAll(Produce(nums))
	if len(got) != len(nums) {
		t.Fatalf("Expected %d elements, got %d", len(nums), len(got))
	}
	for i, v := range nums {
		if got[i] != v {
			t.Errorf("At index %d, expected %d, got %d", i, v, got[i])
		}
	}
}

func TestProduceEmpty(t *testing.T) {
	// 空输入：生产者直接 close，消费者 range 立即结束
	got := ConsumeAll(Produce(nil))
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

func TestBufferedProduce(t *testing.T) {
	nums := []int{10, 20, 30}

	// 容量充足：所有写入不阻塞
	got := ConsumeAll(BufferedProduce(nums, 10))
	if len(got) != 3 || got[0] != 10 || got[2] != 30 {
		t.Errorf("Expected [10 20 30], got %v", got)
	}

	// 边界：容量刚好等于元素个数
	got = ConsumeAll(BufferedProduce(nums, 3))
	if len(got) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(got))
	}

	// 边界：容量不足时自动扩容为元素个数（保证不会写满阻塞）
	got = ConsumeAll(BufferedProduce(nums, 1))
	if len(got) != 3 || got[0] != 10 {
		t.Errorf("Expected [10 20 30], got %v", got)
	}
}

func TestUnbufferedHandshake(t *testing.T) {
	// 无缓冲 channel 的同步握手：双方按严格次序交换数据
	a, b := UnbufferedHandshake()
	if a != 2 {
		t.Errorf("Expected a = 2, got %d", a)
	}
	if b != 1 {
		t.Errorf("Expected b = 1, got %d", b)
	}
}
