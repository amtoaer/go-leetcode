package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		sum /= 10
	}
	return dummyHead.Next
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		addTwoNumbers(
			&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6}}},
		),
		&ListNode{
			Val:  7,
			Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
	)
}
