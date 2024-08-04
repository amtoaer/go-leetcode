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
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(left, right *ListNode) *ListNode {
	pre := left
	cur := pre.Next
	tmp := right.Next
	for cur.Next != tmp {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return cur
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy
	for {
		for i := 0; i < k && right != nil; i++ {
			right = right.Next
		}
		if right == nil {
			break
		}
		left = reverse(left, right)
		right = left
	}
	return dummy.Next
}

// @lc code=end

func Test(t *testing.T) {
	n := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
	}
	n = reverseKGroup(n, 3)
	for _, v := range []int{3, 2, 1, 4, 5} {
		assert.Equal(t, n.Val, v)
		n = n.Next
	}
}
