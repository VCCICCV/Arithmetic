package main

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/24 9:20
// USER:21005
//

// 自动化测试
func TestTwoSum1(t *testing.T) {
	// 测试函数的内存占用
	defer memoryUsage(t)
	defer timeCost(time.Now(), t)

	// 测试用例1：正常情况
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}
	if result1 := twoSum1(nums1, target1); !equal(result1, expected1) {
		t.Errorf("twoSum1(%v, %d) = %v; want %v", nums1, target1, result1, expected1)
	}

	// 测试用例2：空数组
	nums2 := []int{}
	target2 := 9
	expected2 := []int(nil)
	if result2 := twoSum1(nums2, target2); !equal(result2, expected2) {
		t.Errorf("twoSum1(%v, %d) = %v; want %v", nums2, target2, result2, expected2)
	}

	// 测试用例3：无解的数组
	nums3 := []int{2, 7, 11, 15}
	target3 := 10
	expected3 := []int(nil)
	if result3 := twoSum1(nums3, target3); !equal(result3, expected3) {
		t.Errorf("twoSum1(%v, %d) = %v; want %v", nums3, target3, result3, expected3)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum2(t *testing.T) {
	// 测试函数的内存占用
	defer memoryUsage(t)
	defer timeCost(time.Now(), t)

	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tc := range testCases {
		result := twoSum2(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("twoSum2(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
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
