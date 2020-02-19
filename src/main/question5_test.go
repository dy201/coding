package main

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	// 字符串长度
	lens := len(s)
	// 返回值
	result := ""
	// 定义二维数组回文串(Palindrome)，i,j分别表示回文串的起始位置
	// 1000 是题目中给出的字符串最大长度
	var pd [1000][1000]bool

	// 遍历字符串s，查看从j到i是否是回文串
	for i := 0; i < lens; i++ {
		for j := 0; j <= i; j++ {
			// 如果s[j]和s[i]相等，且中间相差 0/1 位字符，则s[j:i]是回文串
			// 或者
			// 如果s[j]和s[i]相等，且向内一层是回文串，则s[j:i]是回文串
			if s[i] == s[j] && (i - j < 2 || pd[j+1][i-1]) {
				pd[j][i] = true
			}
			// 如果 s[j:i] 是回文串，且s[j:i]的长度大于result，更新result。
			// (即更新最大回文串)
			if pd[j][i] && ( i - j + 1) > len(result) {
				result = s[j:i+1]
			}
		}
	}
	return result
}


func TestLongestPalindrome(t *testing.T){
	fmt.Println(longestPalindrome("abedcab")) // a
	fmt.Println(longestPalindrome("abba")) // abba
	fmt.Println(longestPalindrome("babab")) // babab
	fmt.Println(longestPalindrome("alevelb")) // level
}