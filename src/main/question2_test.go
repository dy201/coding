package main

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 定义链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 思路：递归处理，从第一位开始相加，超过10进1.

	// 如果两个链表均为空，返回空，即递归中，结果链表的 Next = nil
	if l1 == nil && l2 == nil {
		return nil
	}

	// 如果l1结果为空，返回l2。即递归中，结果链表的 Next = l2
	if l1 == nil {
		return l2
	}

	// 如果l2结果为空，返回l2。即递归中，结果链表的 Next = l1
	if l2 == nil {
		return l1
	}

	// 链表当前节点的 Val
	sum := l1.Val + l2.Val
	// 链表当前节点的 Next
	nextNode := addTwoNumbers(l1.Next,l2.Next)

	// 判断两数相加结果是够超过10
	// 不超过，继续递归
	if sum < 10 {
		return &ListNode{Val:sum,Next:nextNode}
	}else { // 超过10，当前节点的 Val-10，下一节点 Val+1。
			// 这里使用新建一个 <值为1，Next为nil的节点> ，
			// 与下一节点进行两数相加来达到目的
		tmpNode := &ListNode{
			Val:  1,
			Next: nil,
		}
		return &ListNode{
			Val:sum-10,
			Next:addTwoNumbers(nextNode,tmpNode),
		}
	}
}

// 测试
func TestAddTwoNumbers(t *testing.T)  {
	l1 := makeListNode([]int{3,0,0,1})
	l2 := makeListNode([]int{7})
	numbers := addTwoNumbers(l1, l2)
	fmt.Println(numbers)
}


// 构建单链表
func makeListNode(nums []int) *ListNode{
	if len(nums) == 0 {
		return nil
	}
	// 第一个节点
	list := &ListNode{Val:nums[0]}
	// temp节点代替当前节点向下递归
	temp := list
	// 循环递归，插入数据
	for i := 1; i< len(nums); i++ {
		// 设置下一个节点的 Val
		temp.Next = &ListNode{Val:nums[i]}
		// 设置下一个节点的 Next
		temp = temp.Next
	}
	return list
}