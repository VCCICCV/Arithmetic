package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/24 0:58
// USER:21005
// Definition for singly-linked list.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个虚拟头节点
	dummy := &ListNode{}
	// 指向虚拟头节点
	cur := dummy
	// 进位值
	carry := 0
	// 遍历链表直到为空
	for l1 != nil || l2 != nil {
		// 获取当前节点的值，如果当前节点为空，则值为默认值 0
		var x, y int
		if l1 != nil {
			x = l1.Val   // 获取 l1 的当前节点的值
			l1 = l1.Next // 指向 l1 的下一个节点
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		// 当前节点和进位的和
		sum := x + y + carry
		// 当前节点的进位
		carry = sum / 10
		// 当前节点的和取一位加入链表
		cur.Next = &ListNode{Val: sum % 10}
		// 指向下一个节点
		cur = cur.Next
	}
	// 最高位有进位但链表结束，将进位加入新的节点
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	// 返回虚拟链表的头结点
	return dummy.Next
}

// 将整数转换为链表
func intToListNode(num int) *ListNode {
	// 创建一个头节点
	head := &ListNode{}
	// cur 指向头节点
	cur := head
	// 将整数按位拆分
	for num > 0 {
		// 取模存最低位到链表节点
		cur.Next = &ListNode{Val: num % 10}
		cur = cur.Next
		// num = num / 10
		num /= 10
	}
	// 返回链表的头节点
	return head.Next
}

// 将链表转换为整数
func listNodeToInt(l *ListNode) int {
	num := 0
	base := 1
	// 按个十百位相加
	for l != nil {
		num += l.Val * base
		base *= 10
		l = l.Next
	}
	return num
}

// 输出链表
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
	fmt.Println("nil")
}

// 手动测试
func main() {
	// 测试用例
	l1 := intToListNode(342)
	l2 := intToListNode(465)
	expected := intToListNode(807)
	// 获取当前时间
	start := time.Now()
	result := addTwoNumbers(l1, l2)
	//time.Sleep(time.Second * 1)
	// 计算经过多少时间
	elapsed := time.Since(start)
	fmt.Printf("The elapsed time:%v\n", elapsed)
	//比较结果和期望值
	if listNodeToInt(result) == listNodeToInt(expected) {
		fmt.Println("测试通过")
	} else {
		fmt.Println("测试不通过")
		fmt.Println("期望值：")
		printList(expected)
		fmt.Println("实际值：")
		printList(result)
	}
}
