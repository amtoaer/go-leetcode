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
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		others := cur.Next.Next.Next
		first := cur.Next
		second := cur.Next.Next
		cur.Next = second
		second.Next = first
		first.Next = others
		cur = first
	}
	return dummyHead.Next
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		&ListNode{
			Val:  2,
			Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		},
		swapPairs(
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
		),
	)
}
