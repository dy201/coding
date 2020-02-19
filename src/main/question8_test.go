package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func myAtoi(str string) int {
	var resArr []string
	s := strings.ReplaceAll(str, " ", "")

	if len(s) == 0 {
		return 0
	}

	if len(s) > 2 {
		if (s[0] < '0' || s[0] > '9') && (s[1] < '0' || s[1] > '9'){
			return 0
		}
	}

	if !(s[0] >= 48 && s[0] <=75 || s[0] == 43 || s[0] ==45) {
		return 0
	}

	for _, v := range s {
		if v >= 48 && v <= 57  {
			resArr = append(resArr, string(v))
		}else if v == 45 || v == 43{
			continue
		}else {
			break
		}

	}

	res := 0
	for _, v := range resArr {
		i, _ := strconv.Atoi(v)
		res = res * 10 + i
	}

	if strings.Contains(str,"-"){
		res = -res
	}

	if res > 1<<31 -1 {
		return 1 << 31
	}else if res < -(1 << 31) {
		return -(1 << 31)
	}else {
		return res
	}
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("   1"))
	fmt.Println(myAtoi("   +1"))
	fmt.Println(myAtoi("   +-1"))
	fmt.Println(myAtoi("   -1"))
	fmt.Println(myAtoi("   123.333"))
	fmt.Println(myAtoi("   "))
	fmt.Println(myAtoi("   1abc"))
	fmt.Println(myAtoi("   1123123abc"))
	fmt.Println(myAtoi("abcd0023aa"))
	fmt.Println(myAtoi("abcd-0023aa"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("91283472332"))
}

