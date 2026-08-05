package bloomfilter

import (
	"testing"
)

func TestBloomFilter_AddAndContains(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	bf.AddString("apple")
	bf.AddString("banana")
	bf.AddString("orange")

	if !bf.ContainsString("apple") {
		t.Error("expected Contains('apple') to be true")
	}
	if !bf.ContainsString("banana") {
		t.Error("expected Contains('banana') to be true")
	}
	if bf.ContainsString("grape") {
		t.Error("expected Contains('grape') to be false")
	}
}

func TestBloomFilter_FalsePositive(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	// 插入 1000 个元素
	for i := 0; i < 1000; i++ {
		bf.AddString(string(rune('A' + i%26)) + string(rune('a'+i%26)) + string(rune('0'+i%10)))
	}

	// 检查一些未插入的元素
	fpCount := 0
	total := 1000
	for i := 0; i < total; i++ {
		testStr := "not_inserted_" + string(rune('0'+i%10))
		if bf.ContainsString(testStr) {
			fpCount++
		}
	}

	actualFPR := float64(fpCount) / float64(total)
	estimatedFPR := bf.EstimatedFalsePositiveRate()

	// 实际误判率可能在估算值附近波动
	t.Logf("Actual FPR: %.4f, Estimated FPR: %.4f", actualFPR, estimatedFPR)
	if actualFPR > 0.1 {
		t.Errorf("false positive rate too high: %.4f", actualFPR)
	}
}

func TestBloomFilter_Empty(t *testing.T) {
	bf := NewBloomFilter(100, 0.01)

	if bf.ContainsString("anything") {
		t.Error("empty bloom filter should not contain anything")
	}
	if bf.Count() != 0 {
		t.Errorf("expected count 0, got %d", bf.Count())
	}
}

func TestBloomFilter_Clear(t *testing.T) {
	bf := NewBloomFilter(100, 0.01)
	bf.AddString("test")

	if !bf.ContainsString("test") {
		t.Error("expected Contains('test') to be true")
	}

	bf.Clear()
	if bf.ContainsString("test") {
		t.Error("expected Contains('test') to be false after Clear")
	}
	if bf.Count() != 0 {
		t.Errorf("expected count 0 after Clear, got %d", bf.Count())
	}
}

func TestBloomFilter_Merge(t *testing.T) {
	bf1 := NewBloomFilterWithSize(10000, 5)
	bf2 := NewBloomFilterWithSize(10000, 5)

	bf1.AddString("apple")
	bf2.AddString("banana")

	err := bf1.Merge(bf2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !bf1.ContainsString("apple") {
		t.Error("expected Contains('apple') to be true after merge")
	}
	if !bf1.ContainsString("banana") {
		t.Error("expected Contains('banana') to be true after merge")
	}
}

func TestBloomFilter_FillRatio(t *testing.T) {
	bf := NewBloomFilter(100, 0.01)

	for i := 0; i < 50; i++ {
		bf.AddString(string(rune('a' + i)))
	}

	ratio := bf.FillRatio()
	t.Logf("Fill ratio: %.4f", ratio)
	if ratio <= 0 || ratio > 1 {
		t.Errorf("invalid fill ratio: %.4f", ratio)
	}
}
