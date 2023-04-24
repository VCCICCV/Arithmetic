package lengthoflongestsubstring

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/24 14:59
// USER:21005
func TestLengthOfLongestSubstring(t *testing.T) {
	// 测试函数的内存占用
	defer memoryUsage(t)
	defer timeCost(time.Now(), t)
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

// 耗时统计函数
func timeCost(start time.Time, t *testing.T) {
	tc := time.Since(start)
	t.Logf("time cost:%s\n", tc)
}

// 统计内存占用
func memoryUsage(t *testing.T) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	memUsage := float64(mem.Alloc) / (1024 * 1024)
	t.Logf("Memory usage for lengthOfLongestSubstring: %.4f MB", memUsage)
}
