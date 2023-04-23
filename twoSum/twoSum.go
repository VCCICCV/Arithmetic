package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Arithmetic
// DATE:2023/4/19 0:40
// USER:21005
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
	start := time.Now() // 获取当前时间
	result := twoSum2(nums, target)
	elapsed := time.Since(start).Milliseconds()
	if result == nil {
		fmt.Println("No two elements in the array add up to the target number")
	} else {
		fmt.Printf("The indices of two elements in the array that add up to %d are: %v\n", target, result)
	}
	fmt.Printf("The elapsed time:%d", elapsed)
}
func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
