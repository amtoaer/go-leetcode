package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

// @lc code=end

func Test(t *testing.T) {
	var head *ListNode
	var left int
	var right int
	var hope *ListNode
	var ret *ListNode

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
	}
	left = 2
	right = 4
	hope = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}},
		},
	}
	ret = reverseBetween(head, left, right)
	for ret != nil {
		if ret.Val != hope.Val {
			t.Fatalf("hope %d but ret %d", hope.Val, ret.Val)
		}
		ret = ret.Next
		hope = hope.Next
	}

	head = &ListNode{Val: 5}
	left = 1
	right = 1
	hope = &ListNode{Val: 5}
	ret = reverseBetween(head, left, right)
	for ret != nil {
		if ret.Val != hope.Val {
			t.Fatalf("hope %d but ret %d", hope.Val, ret.Val)
		}
		ret = ret.Next
		hope = hope.Next
	}

	head = &ListNode{Val: 3, Next: &ListNode{Val: 5}}
	left = 1
	right = 2
	hope = &ListNode{Val: 5, Next: &ListNode{Val: 3}}
	ret = reverseBetween(head, left, right)
	for ret != nil {
		if ret.Val != hope.Val {
			t.Fatalf("hope %d but ret %d", hope.Val, ret.Val)
		}
		ret = ret.Next
		hope = hope.Next
	}
}
