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
	dummyHeadL1, dummyHeadL2 := &ListNode{}, &ListNode{}
	dummyHeadL1.Next = l1
	dummyHeadL2.Next = l2
	sub := 0
	startL1, startL2 := dummyHeadL1, dummyHeadL2
	for startL1.Next != nil && startL2.Next != nil {
		startL1 = startL1.Next
		startL2 = startL2.Next
		tmp := startL1.Val + startL2.Val + sub
		startL1.Val = tmp % 10
		sub = tmp / 10
	}
	if startL2.Next != nil {
		startL1.Next = startL2.Next
	}
	for startL1.Next != nil {
		startL1 = startL1.Next
		tmp := startL1.Val + sub
		startL1.Val = tmp % 10
		sub = tmp / 10
	}
	if sub != 0 {
		startL1.Next = &ListNode{Val: sub}
	}
	return dummyHeadL1.Next
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
