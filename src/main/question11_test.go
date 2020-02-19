package main

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	maxArea := 0
	area := 0
	for k, v := range height {
		for i := k; i< len(height); i++ {
			if v <= height[i] {
				area = v * ( i - k )
			}else{
				area = height[i] * ( i - k )
			}
			if maxArea <= area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}