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
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	label1, label2 := dummy1, dummy2
	cur := head
	for cur != nil {
		if cur.Val < x {
			label1.Next = cur
			label1 = label1.Next
		} else {
			label2.Next = cur
			label2 = label2.Next
		}
		cur = cur.Next
	}
	label2.Next = nil
	label1.Next = dummy2.Next
	return dummy1.Next
}

// @lc code=end

func Test(t *testing.T) {
	n := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	partition(n, 3)
	for _, v := range []int{1, 2, 2, 4, 3, 5} {
		assert.Equal(t, v, n.Val)
		n = n.Next
	}
}
