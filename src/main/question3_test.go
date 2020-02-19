package main

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	// 字符串转字符数组
	str := []byte(s)
	// 保存每次循环的结果
	max := 0

	// 如果字符串长度为1，返回1。
	if len(str) == 1 {
		return 1
	}
	// 依次遍历以当前字母开头的不重复字串
	for key, value := range str {
		// 构建set
		m := map[byte]bool{}
		// 计数器（1表示当前字符）
		cnt := 1
		// 将当前字符插入set中
		m[value] = true

		// 寻找当前字母开头的不重复字串
		for j := key + 1; j< len(str);j++  {
			// 字符重复，返回当前最大字串长度
			if m[str[j]] == true  {
				if max < cnt {
					max = cnt
				}
				break
			}else if j == len(str) - 1 { // 遍历到最后一位，返回当前最大字串长度
				cnt ++
				if max < cnt {
					max = cnt
				}
			}else{ // 字串不重复，将字符插入set，并计数
				m[str[j]] = true
				cnt ++
			}
		}
	}
	return max
}

func TestLengthOfLongestSubstring(t *testing.T)  {
	fmt.Println(lengthOfLongestSubstring("")) // 0
	fmt.Println(lengthOfLongestSubstring(" ")) // 1
	fmt.Println(lengthOfLongestSubstring("abcaedbb")) // 5
	fmt.Println(lengthOfLongestSubstring("abcef")) // 5
	fmt.Println(lengthOfLongestSubstring("aaaa")) // 1
}


