package main

import (
	"fmt"
	"testing"
)

func reverse(x int) int {

	// 题目要求32位有符号整数
	// 输入值越界直接返回0
	if x <= -1<<31 || x >= 1<<31-1 {
		return 0
	}
	// 返回值
	res := 0

	for {
		tmp := x % 10

		if x == 0 {
			// 输出值越界直接返回0
			if res <= -1<<31 || res >= 1<<31-1{
				return 0
			}else {
				return res
			}
		}

		x = x / 10
		res = res * 10 + tmp
	}
}

func TestReverse(t *testing.T){
	fmt.Println(reverse(12345))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(9646324351))
	fmt.Println(reverse(1534236469))
	//fmt.Println(1<<31-1)
	//fmt.Println(-1<<31)
}