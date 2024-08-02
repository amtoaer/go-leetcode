package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *ListNode
		output bool
	}{
		{
			input: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}},
			},
			output: false,
		},
		{input: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, output: false},
		{input: &ListNode{Val: 1}, output: false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, hasCycle(tt.input))
	}
}
