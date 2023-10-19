package lengthoflongestsubstring

import (
	"reflect"
	"runtime"
	"testing"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/24 14:59
// USER:21005
//
// 单元测试：测试函数是否正确计算最长不重复子串
func TestLengthOfLongestSubstring(t *testing.T) {
	// 测试函数的内存占用
	defer memoryUsage(t)
	testCases := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, test := range testCases {
		result := lengthOfLongestSubstring(test.s)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("lengthOfLongestSubstring(%v) = %v, want %v", test.s, result, test.expected)
		}
	}
}

// 压测函数：测试函数运行平均耗时
func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb" // 测试字符串
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s) // 测试函数
	}
}

// 统计内存占用
func memoryUsage(t *testing.T) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	memUsage := float64(mem.Alloc) / (1024 * 1024)
	t.Logf("Memory usage for lengthOfLongestSubstring: %.4f MB", memUsage)
}
