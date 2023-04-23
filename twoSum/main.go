package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/19 0:40
// USER:21005

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

func twoSum1(nums []int, target int) []int {
	// 遍历数组中的每一个元素
	for i, x := range nums {
		// 遍历数组中从 i+1 开始的每一个元素
		for j := i + 1; j < len(nums); j++ {
			// 判断两个元素的和是否等于目标值
			if x+nums[j] == target {
				// 符合条件，返回它们的下标
				return []int{i, j}
			}
		}
	}
	// 不符合条件，返回 nil
	return nil
}
func twoSum2(nums []int, target int) []int {
	// 创建一个哈希表，用于记录每个元素的值和下标
	hashTable := map[int]int{}
	// 遍历数组中的每一个元素
	for i, x := range nums {
		// 判断哈希表中是否存在 target-x 这个键
		if p, ok := hashTable[target-x]; ok {
			// 符合条件，返回它们的下标
			return []int{p, i}
		}
		// 将当前元素的值和下标存入哈希表中
		hashTable[x] = i
	}
	// 不符合条件，返回 nil
	return nil
}
func main() {
	var len int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scanln(&len)
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Printf("Enter element %d:", i)
		fmt.Scanln(&nums[i])
	}
	var target int
	fmt.Print("Enter the target number: ")
	fmt.Scanln(&target)
	// 获取当前时间
	start := time.Now()
	result := twoSum2(nums, target)
	// 计算经过多少时间
	elapsed := time.Since(start).Milliseconds()
	// 根据结果输出不同的信息
	if result == nil {
		fmt.Println("No two elements in the array add up to the target number")
	} else {
		fmt.Printf("The indices of two elements in the array that add up to %d are: %v\n", target, result)
	}
	fmt.Printf("The elapsed time:%d", elapsed)
}
