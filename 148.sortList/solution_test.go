package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right, pre := head, head, head
	for right != nil {
		pre = left
		left = left.Next
		right = right.Next
		if right != nil {
			right = right.Next
		}
	}
	pre.Next = nil
	return merge(sortList(head), sortList(left))
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	} else if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *ListNode
		output *ListNode
	}{
		{
			input:  &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}},
			output: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}
	for _, tt := range tc {
		sortList(tt.input)
	}
}
