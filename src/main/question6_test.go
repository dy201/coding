package main

import (
	"fmt"
	"testing"
)

func convert(s string, numRows int) string {
	// 不满足返回
	if len(s) <= 2 || numRows <= 1 {
		return s
	}

	// 创建一个字符串数组，分别存放每一行的字符串
	strArr := make([]string, numRows)

	// 当前位置（字符串数组下标）
	currentPosition := 0
	// 当前方向
	// 我们把初始位置当做转角，方便下面的转向判断
	director := -1

	for _, v := range s {

		// 遇到"z"的拐角，转向掉头
		// 当 位置累加=numRows，或者位置累减=0时，转向
		if currentPosition == 0 || currentPosition == numRows - 1 {
			director = - director
		}

		// 我们用当前位置currentPosition，来表示数组下标
		strArr[currentPosition] += string(v)
		// 控制字符串下标，根据当前方向，做累加（减）
		currentPosition += director
	}

	// 定义返回值字符串
	var result string

	for _, v := range strArr {
		result += v
	}

	return result
}

func TestConvert(t *testing.T) {
	fmt.Println(convert("LEETCODEISHIRING",3))
	fmt.Println(convert("",3))
	fmt.Println(convert(" ",3))
}