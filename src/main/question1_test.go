package main

import (
	"fmt"
	"testing"
)

// 暴力破解法
func twoSum(nums []int, target int) []int {
	result := make([]int,2)
	for i := 0; i< len(nums); i++ {
		result[0] = i
		last := target - nums[result[0]]
		for k, v := range nums {
			if v == last && k!=i {
				result[1] = k
				return result
			}
		}
	}
	return nil
}

// map方式
func twoSum2(nums []int, target int) []int {
	maps := make(map[int]int)
	result := make([]int,2)

	// 目标数组nums放入map中, k,v反置
	for k, v := range nums {
		maps[v] = k
	}

	// 循环遍历目标数组nums
	for k, v := range nums {
		result[0] = k
		last := target - v
		// map中是否包含last
		index,ok := maps[last]
		if ok && index!=k {
			result[1] = index
			break
		}
	}
	return result
}


func TestTwoSum(t *testing.T){
	nums := []int{2, 7, 11, 15}
	// 测试第一种方法
	sum := twoSum(nums,13)
	fmt.Println(sum)
	// 测试第二种方法
	sum2 := twoSum2(nums,13)
	fmt.Println(sum2)
}
