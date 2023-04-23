package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/19 0:40
// USER:21005

// 方法一：暴力枚举
// 思路及算法
// 最容易想到的方法是枚举数组中的每一个数 x，寻找数组中是否存在 target - x。
// 当我们使用遍历整个数组的方式寻找 target - x 时，需要注意到每一个位于 x 之前的元素都已经和 x 匹配过，因此不需要再进行匹配。而每一个元素不能被使用两次，所以我们只需要在 x 后面的元素中寻找 target - x。
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

// 方法二：哈希表
// 思路及算法
// 注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。
// 使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N)O(N)O(N) 降低到 O(1)O(1)O(1)。
// 这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。
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
