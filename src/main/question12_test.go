package main

import (
	"testing"
)

func intToRoman(num int) string {
	// 定义结果集
	var result string
	//var result bytes.Buffer
	// 如果数字小于1，返回空字符串
	if num < 1 {
		return ""
	}
	// 将固定数字从大到小排序，罗马数字对于排列
	nums := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	numsRoman := []string{"M","CM","D","CD","C",
		"XC","L","XL","X","IX","V","IV","I"}

	// 首先把数字分为最优组合，
	for i := 0; i < len(nums) && num > 0; {
		if  num >= nums[i]{
			result += numsRoman[i]
			//result.WriteString(numsRoman[i])
			num = num - nums[i]
		}else {
			i++
		}
	}
	//fmt.Println(result.String())
	//return result.String()
	return result
}

func TestIntToRoman(t *testing.T) {
	intToRoman(3)
	intToRoman(4)
	intToRoman(5)
	intToRoman(1994)
}

