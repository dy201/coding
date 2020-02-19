package main

import (
	"fmt"
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {
	nums:= strconv.Itoa(x)

	for i := 0; i<= len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1]{
			return false
		}
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(0))
}


