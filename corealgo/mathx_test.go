package corealgo

import (
	"reflect"
	"testing"
)

func TestSieveOfEratosthenes(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
		{10, []int{2, 3, 5, 7}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
			53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for _, tc := range cases {
		got := SieveOfEratosthenes(tc.n)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("SieveOfEratosthenes(%d): Expected %v, got %v", tc.n, tc.want, got)
		}
	}
}

func TestGCD(t *testing.T) {
	cases := []struct {
		a, b int
		want int
	}{
		{48, 36, 12},
		{17, 5, 1}, // 互质
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		{-48, 36, 12}, // 负数取绝对值
		{48, -36, 12},
		{100, 25, 25},
	}
	for _, tc := range cases {
		got := GCD(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("GCD(%d, %d): Expected %d, got %d", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestPowMod(t *testing.T) {
	cases := []struct {
		base, exp, mod int64
		want           int64
	}{
		{2, 10, 1000, 24},            // 1024 % 1000
		{3, 5, 7, 5},                 // 243 % 7
		{2, 0, 7, 1},                 // 零次幂
		{2, 20, 1000000007, 1048576}, // 结果小于模数
		{3, 100, 101, 1},             // 费马小定理：3^100 ≡ 1 (mod 101)
		{-2, 3, 5, 2},                // (-8) mod 5 = 2
		{5, 3, 0, 0},                 // 模数为 0
	}
	for _, tc := range cases {
		got := PowMod(tc.base, tc.exp, tc.mod)
		if got != tc.want {
			t.Errorf("PowMod(%d, %d, %d): Expected %d, got %d",
				tc.base, tc.exp, tc.mod, tc.want, got)
		}
	}
}

func TestCombination(t *testing.T) {
	cases := []struct {
		n, k int
		want int
	}{
		{5, 2, 10},
		{10, 3, 120},
		{6, 0, 1},
		{6, 6, 1}, // C(n,n)=1
		{0, 0, 1},
		{5, 7, 0},                    // k > n
		{5, -1, 0},                   // k < 0
		{60, 30, 118264581564861424}, // 大数防溢出用例
		{60, 1, 60},
	}
	for _, tc := range cases {
		got := Combination(tc.n, tc.k)
		if got != tc.want {
			t.Errorf("Combination(%d, %d): Expected %d, got %d", tc.n, tc.k, tc.want, got)
		}
	}
}
