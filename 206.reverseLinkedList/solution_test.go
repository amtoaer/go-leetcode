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
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
	}
	return pre
}

// @lc code=end

func Test(t *testing.T) {
	node4 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	node := reverseList(node1)
	assert.Equal(t, 4, node.Val)
	assert.Equal(t, 3, node.Next.Val)
	assert.Equal(t, 2, node.Next.Next.Val)
	assert.Equal(t, 1, node.Next.Next.Next.Val)
	assert.Nil(t, node.Next.Next.Next.Next)
}
