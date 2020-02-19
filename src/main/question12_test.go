package main

import (
	"fmt"
	"testing"
)

func intToRoman(num int) string {
	return "nil"
}

func TestIntToRoman(t *testing.T) {
	s := "hello,你好"
	for i := 0; i< len(s);i++{
		ch := s[i]
		fmt.Println(i,ch)
	}

	fmt.Println("----------")

	for k, v := range s {
		fmt.Println(k,v)
	}
}

