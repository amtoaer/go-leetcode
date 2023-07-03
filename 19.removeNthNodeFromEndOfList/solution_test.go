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
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	left, right := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		if right.Next == nil {
			return dummyHead.Next
		}
		right = right.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	if left.Next != nil {
		left.Next = left.Next.Next
	}
	return dummyHead.Next
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		removeNthFromEnd(
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
			1,
		),
	)
	assert.Equal(
		t,
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		removeNthFromEnd(
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
			2,
		),
	)
}
