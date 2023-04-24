package main

import "testing"

// PROJECT_NAME:Arithmetic
// DATE:2023/4/24 14:27
// USER:21005

func TestAddTwoNumbers(t *testing.T) {
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
