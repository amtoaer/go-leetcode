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
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	left, right := head, head
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
		if left == right {
			break
		}
	}
	if right == nil || right.Next == nil {
		return nil
	}
	left = head
	for left != right {
		left = left.Next
		right = right.Next
	}
	return left
}

// @lc code=end

func Test(t *testing.T) {
	nodes := []*ListNode{
		{Val: 3},
		{Val: 2},
		{Val: 0},
		{Val: -4},
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nodes[1]
	assert.Equal(t, nodes[1], detectCycle(nodes[0]))
}
