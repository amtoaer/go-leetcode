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
 * @lc app=leetcode.cn id=2181 lang=golang
 *
 * [2181] Merge Nodes in Between Zeros
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	ptr := dummy
	tmp := 0
	for cur := head; cur != nil; cur = cur.Next {
		if cur == head {
			continue
		}
		if cur.Next == nil || cur.Val == 0 {
			ptr.Next = &ListNode{
				Val: tmp,
			}
			tmp = 0
			ptr = ptr.Next
			continue
		}
		tmp += cur.Val
	}
	return dummy.Next
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *ListNode
		output *ListNode
	}{
		{nil, nil},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, mergeNodes(tt.input))
	}
}
