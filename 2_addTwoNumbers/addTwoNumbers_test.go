package main

import (
	"runtime"
	"testing"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/24 14:27
// USER:21005
func TestAddTwoNumbers(t *testing.T) {
	// defer是逆序执行，在前面的后执行，先显示时间再显示内存占用
	defer memoryUsage(t)
	// 时间太短
	//defer timeCost(time.Now(), t)
	// 构造测试数据
	l1 := intToListNode(342)
	l2 := intToListNode(465)
	expected := intToListNode(807)
	// 执行函数
	result := addTwoNumbers(l1, l2)
	// 检查结果是否符合预期
	if listNodeToInt(result) != listNodeToInt(expected) {
		t.Errorf("AddTwoNumbers(%v, %v) = %v; want %v", l1, l2, result, expected)
	}
}

// 耗时统计函数 时间太短使用基准测试计算时间
//func timeCost(start time.Time, t *testing.T) {
//	tc := time.Since(start)
//	t.Logf("time cost:%s\n", tc)
//}

// 压测函数：测试函数运行平均耗时
func BenchmarkAddTwoNumbers(b *testing.B) {
	l1 := intToListNode(342)
	l2 := intToListNode(465)
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}

// 统计内存占用
func memoryUsage(t *testing.T) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	memUsage := float64(mem.Alloc) / (1024 * 1024)
	t.Logf("Memory usage for lengthOfLongestSubstring: %.4f MB", memUsage)
}
