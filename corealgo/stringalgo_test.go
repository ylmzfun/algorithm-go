package corealgo

import (
	"reflect"
	"testing"
)

func TestKMPSearch(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		want    []int
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB", []int{10}},
		{"aaaa", "aa", []int{0, 1, 2}}, // 重叠匹配
		{"abababa", "aba", []int{0, 2, 4}},
		{"aaa", "a", []int{0, 1, 2}},
		{"hello", "world", []int{}}, // 无匹配
		{"abc", "", []int{}},        // 空模式串
		{"ab", "abc", []int{}},      // 模式串长于主串
		{"a", "a", []int{0}},
		{"", "a", []int{}}, // 空主串
	}
	for _, tc := range cases {
		got := KMPSearch(tc.text, tc.pattern)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("KMPSearch(%q, %q): Expected %v, got %v", tc.text, tc.pattern, tc.want, got)
		}
	}
}

func TestRabinKarpSearch(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		want    []int
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB", []int{10}},
		{"aaaa", "aa", []int{0, 1, 2}}, // 重叠匹配
		{"abababa", "aba", []int{0, 2, 4}},
		{"aaa", "a", []int{0, 1, 2}},
		{"hello", "world", []int{}}, // 无匹配
		{"abc", "", []int{}},        // 空模式串
		{"ab", "abc", []int{}},      // 模式串长于主串
		{"a", "a", []int{0}},
		{"", "a", []int{}}, // 空主串
	}
	for _, tc := range cases {
		got := RabinKarpSearch(tc.text, tc.pattern)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("RabinKarpSearch(%q, %q): Expected %v, got %v", tc.text, tc.pattern, tc.want, got)
		}
	}
}
